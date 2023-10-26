package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/market_mc/market_go_product_service/config"
	"gitlab.com/market_mc/market_go_product_service/storage"
)

type Store struct {
	db         *pgxpool.Pool
	products   *ProductRepo
	categories *CategoryRepo
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

func (s *Store) Product() storage.ProductRepoI {
	if s.products == nil {
		s.products = NewProductRepo(s.db)
	}
	return s.products
}

func (s *Store) Category() storage.CategoryRepoI {
	if s.categories == nil {
		s.categories = NewCategoryRepo(s.db)
	}
	return s.categories
}
