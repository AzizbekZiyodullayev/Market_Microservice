package postgres

import (
	"context"
	"fmt"
	"genproto/sale_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BranchProductTransactionRepo struct {
	db *pgxpool.Pool
}

func NewBranchProductTransactionRepo(db *pgxpool.Pool) *BranchProductTransactionRepo {
	return &BranchProductTransactionRepo{
		db: db,
	}
}

func (r *BranchProductTransactionRepo) Create(ctx context.Context, req *sale_service.BrPrTrCreateReq) (*sale_service.BrPrTrCreateResp, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO branch_product_transactions(
		id,
		branch_id,
		staff_id,
		product_id,
		price,
		type,
		quantity
	)
	VALUES($1,$2,$3,$4,$5,$6,$7);
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.BranchId,
		req.StaffId,
		req.ProductId,
		req.Price,
		req.Typ,
		req.Quantity,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}

	return &sale_service.BrPrTrCreateResp{
		Msg: "branch product transaction created with id: " + id,
	}, nil
}

func (r *BranchProductTransactionRepo) GetList(ctx context.Context, req *sale_service.BrPrTrGetListReq) (*sale_service.BrPrTrGetListResp, error) {
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
		staff_id,
		product_id,
		price,
		type,
		quantity,
		created_at::TEXT,
		updated_at::TEXT 
	FROM branch_product_transactions `

	if req.BranchId != "" {
		filter += ` AND branch_id='` + req.BranchId + `' `
	}
	if req.StaffId != "" {
		filter += ` AND staff_id='` + req.StaffId + `' `
	}
	if req.ProductId != "" {
		filter += ` AND product_id='` + req.ProductId + `' `
	}
	if req.Typ != "" {
		filter += ` AND type='` + req.Typ + `' `
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ
	// fmt.Println(query)

	countS := `SELECT COUNT(*) FROM branch_product_transactions` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp := &sale_service.BrPrTrGetListResp{}
	for rows.Next() {
		var brPrTr = sale_service.BrPrTransaction{}
		err := rows.Scan(
			&brPrTr.Id,
			&brPrTr.BranchId,
			&brPrTr.StaffId,
			&brPrTr.ProductId,
			&brPrTr.Price,
			&brPrTr.Typ,
			&brPrTr.Quantity,
			&brPrTr.CreatedAt,
			&brPrTr.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp.BrPrTransactions = append(resp.BrPrTransactions, &brPrTr)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *BranchProductTransactionRepo) GetById(ctx context.Context, req *sale_service.BrPrTrIdReq) (*sale_service.BrPrTransaction, error) {
	query := `
	SELECT 
		id,
		branch_id,
		staff_id,
		product_id,
		price,
		type,
		quantity,
		created_at::TEXT,
		updated_at::TEXT 
	FROM branch_product_transactions
	WHERE id=$1 AND deleted_at IS NULL;
	`

	var brPrTr = sale_service.BrPrTransaction{}
	if err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&brPrTr.Id,
		&brPrTr.BranchId,
		&brPrTr.StaffId,
		&brPrTr.ProductId,
		&brPrTr.Price,
		&brPrTr.Typ,
		&brPrTr.Quantity,
		&brPrTr.CreatedAt,
		&brPrTr.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &brPrTr, nil
}

func (r *BranchProductTransactionRepo) Update(ctx context.Context, req *sale_service.BrPrTrUpdateReq) (*sale_service.BrPrTrUpdateResp, error) {
	query := `
	UPDATE branch_product_transactions 
	SET 
		branch_id=$2,
		staff_id=$3,
		product_id=$4,
		price=$5,
		type=$6,
		quantity=$7,
		updated_at=NOW()
	WHERE id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
		req.BranchId,
		req.StaffId,
		req.ProductId,
		req.Price,
		req.Typ,
		req.Quantity,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}

	return &sale_service.BrPrTrUpdateResp{Msg: "OK"}, nil
}

func (r *BranchProductTransactionRepo) Delete(ctx context.Context, req *sale_service.BrPrTrIdReq) (*sale_service.BrPrTrDeleteResp, error) {
	query := `
    UPDATE branch_product_transactions 
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

	return &sale_service.BrPrTrDeleteResp{Msg: "OK"}, nil
}
