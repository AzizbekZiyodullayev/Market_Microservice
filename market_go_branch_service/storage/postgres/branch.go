package postgres

import (
	"context"
	"fmt"
	"genproto/branch_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BranchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *BranchRepo {
	return &BranchRepo{
		db: db,
	}
}

func (r *BranchRepo) Create(ctx context.Context, req *branch_service.BranchCreateReq) (*branch_service.BranchCreateResp, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO branches (
		id,
		name,
		address,
		founded_at
	)
	VALUES ($1,$2,$3,$4);
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Address,
		req.FoundedAt,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}

	return &branch_service.BranchCreateResp{
		Msg: "branch created with id: " + id,
	}, nil
}

func (r *BranchRepo) GetList(ctx context.Context, req *branch_service.BranchGetListReq) (*branch_service.BranchGetListResp, error) {
	var (
		filter  = " WHERE deleted_at IS NULL "
		offsetQ = " OFFSET 0;"
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   int
	)

	s := `
	SELECT 
		id,
		name,
		address,
		founded_at::TEXT,
		created_at::TEXT,
		updated_at::TEXT 
	FROM branches `

	if req.Search != "" {
		filter += ` AND name ILIKE ` + "'%" + req.Search + "%' OR address ILIKE '%" + req.Search + "%' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM branches` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp := &branch_service.BranchGetListResp{}
	for rows.Next() {
		var branch = branch_service.Branch{}
		err := rows.Scan(
			&branch.Id,
			&branch.Name,
			&branch.Address,
			&branch.FoundedAt,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp.Branches = append(resp.Branches, &branch)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *BranchRepo) GetById(ctx context.Context, req *branch_service.BranchIdReq) (*branch_service.Branch, error) {
	query := `
    SELECT 
        id,
        name,
        address,
        founded_at::TEXT,
        created_at::TEXT,
        updated_at::TEXT 
    FROM branches 
    WHERE id = $1
    `

	var branch = branch_service.Branch{}

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&branch.Id,
		&branch.Name,
		&branch.Address,
		&branch.FoundedAt,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &branch, nil
}

func (r *BranchRepo) Update(ctx context.Context, req *branch_service.BranchUpdateReq) (*branch_service.BranchUpdateResp, error) {
	query := `
    UPDATE branches 
    SET 
        name = $1,
        address = $2,
        founded_at = $3,
		updated_at = NOW()
    WHERE id = $4
    `

	_, err := r.db.Exec(ctx, query,
		req.Name,
		req.Address,
		req.FoundedAt,
		req.Id,
	)

	if err != nil {
		return nil, err
	}

	return &branch_service.BranchUpdateResp{
		Msg: "branch updated",
	}, nil
}

func (r *BranchRepo) Delete(ctx context.Context, req *branch_service.BranchIdReq) (*branch_service.BranchDeleteResp, error) {
	query := `
    UPDATE branches 
    SET 
        deleted_at = NOW()
    WHERE id = $1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
	)

	if err != nil {
		return nil, err
	}

	if res.RowsAffected() == 0 {
		return nil, fmt.Errorf("branch not found")
	}

	return &branch_service.BranchDeleteResp{
		Msg: "branch deleted",
	}, nil
}
