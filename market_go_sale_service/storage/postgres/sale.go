package postgres

import (
	"context"
	"fmt"
	"genproto/sale_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SaleRepo struct {
	db *pgxpool.Pool
}

func NewSaleRepo(db *pgxpool.Pool) *SaleRepo {
	return &SaleRepo{
		db: db,
	}
}

func (r *SaleRepo) Create(ctx context.Context, req *sale_service.SaleCreateReq) (*sale_service.SaleCreateResp, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO sales(
		id,
		branch_id,
		shop_assistant_id,
		cashier_id,
		price,
		payment_type,
		status,
		client_name
	)
	VALUES($1,$2,$3,$4,$5,$6,$7,$8);
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.BranchId,
		req.ShopAssistantId,
		req.CashierId,
		req.Price,
		req.PaymentType,
		req.Status,
		req.ClientName,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}

	return &sale_service.SaleCreateResp{
		Msg: "sale created with id: " + id,
	}, nil
}

func (r *SaleRepo) GetList(ctx context.Context, req *sale_service.SaleGetListReq) (*sale_service.SaleGetListResp, error) {
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
		branch_id,
		shop_assistant_id,
		cashier_id,
		price,
		payment_type,
		status,
		client_name,
		created_at::TEXT,
		updated_at::TEXT 
	FROM sales `

	if req.ClientName != "" {
		filter += ` AND name ILIKE ` + "'%" + req.ClientName + "%' OR address ILIKE '%" + req.ClientName + "%' "
	}
	if req.BranchId != "" {
		filter += ` AND branch_id='` + req.BranchId + `' `
	}
	if req.CashierId != "" {
		filter += ` AND cashier_id='` + req.CashierId + `' `
	}
	if req.ShopAssistantId != "" {
		filter += ` AND shop_assistant_id='` + req.ShopAssistantId + `' `
	}
	if req.PaymentType != "" {
		filter += ` AND payment_type='` + req.PaymentType + `' `
	}
	if req.Status != "" {
		filter += ` AND status='` + req.Status + `' `
	}

	filter += ` AND price BETWEEN ` + fmt.Sprintf("%f", req.PriceFrom) + ` AND ` + fmt.Sprintf("%f", req.PriceTo)

	filter += ` AND created_at BETWEEN '` + req.CreatedAtFrom + `' AND '` + req.CreatedAtTo + `' `

	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ
	// fmt.Println(query)

	countS := `SELECT COUNT(*) FROM sales` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp := &sale_service.SaleGetListResp{}
	for rows.Next() {
		var sale = sale_service.Sale{}
		err := rows.Scan(
			&sale.Id,
			&sale.BranchId,
			&sale.ShopAssistantId,
			&sale.CashierId,
			&sale.Price,
			&sale.PaymentType,
			&sale.Status,
			&sale.ClientName,
			&sale.CreatedAt,
			&sale.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp.Sales = append(resp.Sales, &sale)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *SaleRepo) GetById(ctx context.Context, req *sale_service.SaleIdReq) (*sale_service.Sale, error) {
	query := `
	SELECT 
		id,
		branch_id,
		shop_assistant_id,
		cashier_id,
		price,
		payment_type,
		status,
		client_name,
		created_at::TEXT,
		updated_at::TEXT
	FROM sales
	WHERE id=$1 AND deleted_at IS NULL;
	`

	var sale = sale_service.Sale{}
	if err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&sale.Id,
		&sale.BranchId,
		&sale.ShopAssistantId,
		&sale.CashierId,
		&sale.Price,
		&sale.PaymentType,
		&sale.Status,
		&sale.ClientName,
		&sale.CreatedAt,
		&sale.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &sale, nil
}

func (r *SaleRepo) Update(ctx context.Context, req *sale_service.SaleUpdateReq) (*sale_service.SaleUpdateResp, error) {
	query := `
	UPDATE sales 
	SET 
		branch_id=$2,
		shop_assistant_id=$3,
		cashier_id=$4,
		price=$5,
		payment_type=$6,
		status=$7,
		client_name=$8,
		updated_at=NOW()
	WHERE id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
		req.BranchId,
		req.ShopAssistantId,
		req.CashierId,
		req.Price,
		req.PaymentType,
		req.Status,
		req.ClientName,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}

	return &sale_service.SaleUpdateResp{Msg: "OK"}, nil
}

func (r *SaleRepo) Delete(ctx context.Context, req *sale_service.SaleIdReq) (*sale_service.SaleDeleteResp, error) {
	query := `
    UPDATE sales 
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

	return &sale_service.SaleDeleteResp{Msg: "OK"}, nil
}
