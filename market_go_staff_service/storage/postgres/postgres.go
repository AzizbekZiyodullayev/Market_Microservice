package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/market_mc/market_go_staff_service/config"
	"gitlab.com/market_mc/market_go_staff_service/storage"
)

type Store struct {
	db           *pgxpool.Pool
	staffTariffs *StaffTariffRepo
	staffs       *StaffRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: pool,
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) StaffTariff() storage.StaffTariffRepoI {
	if s.staffTariffs == nil {
		s.staffTariffs = NewStaffTariffRepo(s.db)
	}
	return s.staffTariffs
}

func (s *Store) Staff() storage.StaffRepoI {
	if s.staffs == nil {
		s.staffs = NewStaffRepo(s.db)
	}
	return s.staffs
}
