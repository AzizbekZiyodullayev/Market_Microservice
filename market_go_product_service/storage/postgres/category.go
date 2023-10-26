package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"genproto/product_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/market_mc/market_go_product_service/pkg/helper"
)

type CategoryRepo struct {
	db *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{
		db: db,
	}
}

func (r *CategoryRepo) Create(ctx context.Context, req *product_service.CategoryCreateReq) (*product_service.CategoryCreateResp, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO categories(
		id,
		name,
		parent_id
	) VALUES($1, $2, $3);
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		helper.NewNullString(req.ParentId),
	)
	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}

	return &product_service.CategoryCreateResp{Msg: "category created with id: " + id}, nil
}

func (r *CategoryRepo) GetList(ctx context.Context, req *product_service.CategoryGetListReq) (*product_service.CategoryGetListResp, error) {
	var (
		filter   = " WHERE deleted_at IS NULL "
		offsetQ  = " OFFSET 0;"
		limit    = " LIMIT 10 "
		offset   = (req.Page - 1) * req.Limit
		count    int
		parentId sql.NullString
	)

	s := `
	SELECT 
		id,
		name,
		parent_id,
		created_at::TEXT,
		updated_at::TEXT 
	FROM categories `

	if req.Name != "" {
		filter += ` AND name ILIKE ` + "'%" + req.Name + "%' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM categories` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp := &product_service.CategoryGetListResp{}
	for rows.Next() {
		var category = product_service.Category{}
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&parentId,
			&category.CreatedAt,
			&category.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp.Categories = append(resp.Categories, &product_service.Category{
			Id:        category.Id,
			Name:      category.Name,
			ParentId:  parentId.String,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		})
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *CategoryRepo) GetById(ctx context.Context, req *product_service.CategoryIdReq) (*product_service.Category, error) {
	query := `
    SELECT 
        id,
        name,
        parent_id,
        created_at::TEXT,
        updated_at::TEXT 
    FROM categories 
    WHERE id=$1 AND deleted_at IS NULL;`

	var (
		category = product_service.Category{}
		parentId sql.NullString
	)
	if err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&category.Id,
		&category.Name,
		&parentId,
		&category.CreatedAt,
		&category.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &product_service.Category{
		Id:        category.Id,
		Name:      category.Name,
		ParentId:  parentId.String,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}, nil
}

func (r *CategoryRepo) Update(ctx context.Context, req *product_service.CategoryUpdateReq) (*product_service.CategoryUpdateResp, error) {
	query := `
    UPDATE categories 
    SET 
        name=$2,
        parent_id=$3,
        updated_at=NOW() 
    WHERE id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
		req.Name,
		helper.NewNullString(req.ParentId),
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	return &product_service.CategoryUpdateResp{Msg: "OK"}, nil
}

func (r *CategoryRepo) Delete(ctx context.Context, req *product_service.CategoryIdReq) (*product_service.CategoryDeleteResp, error) {
	query := `
    UPDATE categories 
    SET 
        deleted_at=NOW() 
    WHERE id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}

	return &product_service.CategoryDeleteResp{Msg: "OK"}, nil
}
