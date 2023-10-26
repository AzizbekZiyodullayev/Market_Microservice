package postgres

import (
	"context"
	"fmt"
	"genproto/sale_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffTransactionRepo struct {
	db *pgxpool.Pool
}

func NewStaffTransactionRepo(db *pgxpool.Pool) *StaffTransactionRepo {
	return &StaffTransactionRepo{
		db: db,
	}
}

func (r *StaffTransactionRepo) Create(ctx context.Context, req *sale_service.StaffTrCreateReq) (*sale_service.StaffTrCreateResp, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO staff_transactions(
		id,
		sale_id,
		staff_id,
		type,
		source_type,
		amount,
		about_text
	) VALUES($1,$2,$3,$4,$5,$6,$7);
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.SaleId,
		req.StaffId,
		req.Typ,
		req.SourceTyp,
		req.Amount,
		req.AboutText,
	)

	if err != nil {
		return nil, err
	}

	return &sale_service.StaffTrCreateResp{
		Msg: "staff_transaction created with id: " + id,
	}, nil
}

func (r *StaffTransactionRepo) GetList(ctx context.Context, req *sale_service.StaffTrGetListReq) (*sale_service.StaffTrGetListResp, error) {
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
		sale_id,
		staff_id,
		type,
		source_type,
		amount,
		about_text,
		created_at::TEXT,
		updated_at::TEXT 
	FROM staff_transactions `

	if req.SaleId != "" {
		filter += ` AND sale_id='` + req.SaleId + `' `
	}
	if req.Typ != "" {
		filter += ` AND type='` + req.Typ + `' `
	}
	if req.StaffId != "" {
		filter += ` AND staff_id='` + req.StaffId + `' `
	}

	filter += ` AND amount BETWEEN '` + fmt.Sprintf("%f", req.AmountFrom) + `' AND '` + fmt.Sprintf("%f", req.AmountTo) + `' `

	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ
	// fmt.Println(query)

	countS := `SELECT COUNT(*) FROM staff_transactions` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp := &sale_service.StaffTrGetListResp{}
	for rows.Next() {
		var staffTr = sale_service.StaffTransaction{}
		err := rows.Scan(
			&staffTr.Id,
			&staffTr.SaleId,
			&staffTr.StaffId,
			&staffTr.Typ,
			&staffTr.SourceTyp,
			&staffTr.Amount,
			&staffTr.AboutText,
			&staffTr.CreatedAt,
			&staffTr.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp.StaffTransactions = append(resp.StaffTransactions, &staffTr)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *StaffTransactionRepo) GetById(ctx context.Context, req *sale_service.StaffTrIdReq) (*sale_service.StaffTransaction, error) {
	query := `
	SELECT 
		id,
		sale_id,
		staff_id,
		type,
		source_type,
		amount,
		about_text,
		created_at::TEXT,
		updated_at::TEXT
	FROM staff_transactions
	WHERE id=$1 AND deleted_at IS NULL;
	`

	var staffTr = sale_service.StaffTransaction{}
	if err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&staffTr.Id,
		&staffTr.SaleId,
		&staffTr.StaffId,
		&staffTr.Typ,
		&staffTr.SourceTyp,
		&staffTr.Amount,
		&staffTr.AboutText,
		&staffTr.CreatedAt,
		&staffTr.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &staffTr, nil
}

func (r *StaffTransactionRepo) Update(ctx context.Context, req *sale_service.StaffTrUpdateReq) (*sale_service.StaffTrUpdateResp, error) {
	query := `
	UPDATE staff_transactions 
	SET 
		sale_id=$2,
		staff_id=$3,
		type=$4,
		source_type=$5,
		amount=$6,
		about_text=$7,
		updated_at=NOW()
	WHERE id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
		req.SaleId,
		req.StaffId,
		req.Typ,
		req.SourceTyp,
		req.Amount,
		req.AboutText,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}

	return &sale_service.StaffTrUpdateResp{Msg: "OK"}, nil
}

func (r *StaffTransactionRepo) Delete(ctx context.Context, req *sale_service.StaffTrIdReq) (*sale_service.StaffTrDeleteResp, error) {
	query := `
    UPDATE staff_transactions 
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

	return &sale_service.StaffTrDeleteResp{Msg: "OK"}, nil
}
