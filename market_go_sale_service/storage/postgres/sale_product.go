package postgres

import (
	"context"
	"fmt"
	"genproto/sale_service"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SaleProductRepo struct {
	db *pgxpool.Pool
}

func NewSaleProductRepo(db *pgxpool.Pool) *SaleProductRepo {
	return &SaleProductRepo{
		db: db,
	}
}

func (r *SaleProductRepo) Create(ctx context.Context, req *sale_service.SaleProductCreateReq) (*sale_service.SaleProductCreateResp, error) {
	query := `
	INSERT INTO sale_products(
        sale_id,
        product_id,
		quantity,
		price
	)
	VALUES($1,$2,$3,$4);
	`

	_, err := r.db.Exec(ctx, query,
		req.SaleId,
		req.ProductId,
		req.Quantity,
		req.Price,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}

	return &sale_service.SaleProductCreateResp{
		Msg: "sale product created with id: " + req.SaleId,
	}, nil
}

func (r *SaleProductRepo) GetList(ctx context.Context, req *sale_service.SaleProductGetListReq) (*sale_service.SaleProductGetListResp, error) {
	var (
		filter  = " WHERE TRUE "
		offsetQ = " OFFSET 0;"
		limit   = " LIMIT 10 "
		offset  = (req.Page - 1) * req.Limit
		count   int
	)

	s := `
	SELECT 
		sale_id,
		product_id,
		quantity,
		price
	FROM sale_products `

	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM sale_products` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err = r.db.QueryRow(ctx, countS).Scan(&count); err != nil {
		return nil, err
	}

	resp := &sale_service.SaleProductGetListResp{}
	for rows.Next() {
		var saleProduct = sale_service.SaleProduct{}
		err := rows.Scan(
			&saleProduct.SaleId,
			&saleProduct.ProductId,
			&saleProduct.Quantity,
			&saleProduct.Price,
		)

		if err != nil {
			return nil, err
		}
		resp.SaleProducts = append(resp.SaleProducts, &saleProduct)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *SaleProductRepo) GetById(ctx context.Context, req *sale_service.SaleProductIdReq) (*sale_service.SaleProduct, error) {
	query := `
    SELECT 
        sale_id,
        product_id,
        quantity,
        price
    FROM sale_products
    WHERE sale_id=$1;
    `

	var saleProduct = sale_service.SaleProduct{}
	if err := r.db.QueryRow(ctx, query, req.SaleId).Scan(
		&saleProduct.SaleId,
		&saleProduct.ProductId,
		&saleProduct.Quantity,
		&saleProduct.Price,
	); err != nil {
		return nil, err
	}

	return &saleProduct, nil
}

func (r *SaleProductRepo) Update(ctx context.Context, req *sale_service.SaleProductUpdateReq) (*sale_service.SaleProductUpdateResp, error) {
	query := `
    UPDATE sale_products 
    SET 
        product_id=$2,
        quantity=$3,
        price=$4
    WHERE sale_id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.SaleId,
		req.ProductId,
		req.Quantity,
		req.Price,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}

	return &sale_service.SaleProductUpdateResp{Msg: "OK"}, nil
}

func (r *SaleProductRepo) Delete(ctx context.Context, req *sale_service.SaleProductIdReq) (*sale_service.SaleProductDeleteResp, error) {
	query := `
	DELETE FROM 
		sale_products
	WHERE sale_id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.SaleId,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}

	return &sale_service.SaleProductDeleteResp{Msg: "OK"}, nil
}
