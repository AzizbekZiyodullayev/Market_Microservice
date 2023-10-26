package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/market_mc/market_go_sale_service/config"
	"gitlab.com/market_mc/market_go_sale_service/storage"
)

type Store struct {
	db                        *pgxpool.Pool
	sales                     *SaleRepo
	saleProducts              *SaleProductRepo
	staffTransactions         *StaffTransactionRepo
	branchProductTransactions *BranchProductTransactionRepo
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

func (s *Store) Sale() storage.SaleRepoI {
	if s.sales == nil {
		s.sales = NewSaleRepo(s.db)
	}

	return s.sales
}

func (s *Store) SaleProduct() storage.SaleProductRepoI {
	if s.saleProducts == nil {
		s.saleProducts = NewSaleProductRepo(s.db)
	}

	return s.saleProducts
}

func (s *Store) StaffTransaction() storage.StaffTransactionRepoI {
	if s.staffTransactions == nil {
		s.staffTransactions = NewStaffTransactionRepo(s.db)
	}

	return s.staffTransactions
}

func (s *Store) BranchProductTransaction() storage.BranchProductTransactionRepoI {
	if s.branchProductTransactions == nil {
		s.branchProductTransactions = NewBranchProductTransactionRepo(s.db)
	}

	return s.branchProductTransactions
}
