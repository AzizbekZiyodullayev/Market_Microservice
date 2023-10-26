package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/market_mc/market_go_branch_service/config"
	"gitlab.com/market_mc/market_go_branch_service/storage"
)

type Store struct {
	db              *pgxpool.Pool
	branches        *BranchRepo
	branch_products *BranchProductRepo
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

func (s *Store) Branch() storage.BranchRepoI {
	if s.branches == nil {
		s.branches = NewBranchRepo(s.db)
	}
	return s.branches
}

func (s *Store) BranchProduct() storage.BranchProductRepoI {
	if s.branch_products == nil {
		s.branch_products = NewBranchProductRepo(s.db)
	}
	return s.branch_products
}
