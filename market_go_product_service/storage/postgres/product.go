package postgres

import (
	"context"
	"fmt"
	"genproto/product_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r *ProductRepo) Create(ctx context.Context, req *product_service.ProductCreateReq) (*product_service.ProductCreateResp, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO products (
		id,
		name,
		price,
		barcode,
		category_id
	)
	VALUES ($1,$2,$3,$4,$5);
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}

	return &product_service.ProductCreateResp{Msg: "product created with id: " + id}, nil
}

func (r *ProductRepo) GetList(ctx context.Context, req *product_service.ProductGetListReq) (resp *product_service.ProductGetListResp, err error) {
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
		price,
		barcode,
		category_id,
		created_at::TEXT,
		updated_at::TEXT 
	FROM products `

	if req.Name != "" {
		filter += ` AND name ILIKE ` + "'%" + req.Name + "%' "
	}
	if req.Barcode != "" {
		filter += ` AND barcode = ` + "'" + req.Barcode + "' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM products` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp = &product_service.ProductGetListResp{}
	for rows.Next() {
		var product = product_service.Product{}
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Barcode,
			&product.CategoryId,
			&product.CreatedAt,
			&product.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp.Products = append(resp.Products, &product)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *ProductRepo) GetById(ctx context.Context, req *product_service.ProductIdReq) (res *product_service.Product, err error) {
	query := `
	SELECT 
		id,
		name,
		price,
		barcode,
		category_id,
		created_at::TEXT,
		updated_at::TEXT 
	FROM products 
	WHERE id=$1 AND deleted_at IS NULL;`

	var product = product_service.Product{}
	if err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Barcode,
		&product.CategoryId,
		&product.CreatedAt,
		&product.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepo) Update(ctx context.Context, req *product_service.ProductUpdateReq) (resp *product_service.ProductUpdateResp, err error) {
	query := `
	UPDATE products 
	SET 
		name=$2,
		price=$3,
		barcode=$4,
		category_id=$5,
		updated_at=NOW() 
	WHERE id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.Price,
		req.Barcode,
		req.CategoryId,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	return &product_service.ProductUpdateResp{Msg: "OK"}, nil
}

func (r *ProductRepo) Delete(ctx context.Context, req *product_service.ProductIdReq) (*product_service.ProductDeleteResp, error) {
	query := `
    UPDATE products 
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

	return &product_service.ProductDeleteResp{Msg: "OK"}, nil
}
