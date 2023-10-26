package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type BranchProductRepo struct {
	db *pgxpool.Pool
}

func NewBranchProductRepo(db *pgxpool.Pool) *BranchProductRepo {
	return &BranchProductRepo{
		db: db,
	}
}
