package postgres

import (
	"context"
	"fmt"
	"genproto/staff_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffTariffRepo struct {
	db *pgxpool.Pool
}

func NewStaffTariffRepo(db *pgxpool.Pool) *StaffTariffRepo {
	return &StaffTariffRepo{
		db: db,
	}
}

func (r *StaffTariffRepo) Create(ctx context.Context, req *staff_service.TariffCreateReq) (*staff_service.TariffCreateResp, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO staff_tariffs(
		id,
		name,
		type,
		amount_for_cash,
		amount_for_card
	) VALUES($1, $2, $3, $4, $5);
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.Name,
		req.Typ,
		req.AmountForCash,
		req.AmountForCard,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}

	return &staff_service.TariffCreateResp{
		Msg: "staff tariff created with id: " + id,
	}, nil
}

func (r *StaffTariffRepo) GetList(ctx context.Context, req *staff_service.TariffGetListReq) (*staff_service.TariffGetListResp, error) {
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
		type,
		amount_for_cash,
		amount_for_card,
		created_at::TEXT,
		updated_at::TEXT 
	FROM staff_tariffs `

	if req.Typ != "" {
		filter += ` AND type = ` + "'" + req.Typ + "' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM staff_tariffs` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp := &staff_service.TariffGetListResp{}
	for rows.Next() {
		var tariff = staff_service.StaffTariff{}
		err := rows.Scan(
			&tariff.Id,
			&tariff.Name,
			&tariff.Typ,
			&tariff.AmountForCash,
			&tariff.AmountForCard,
			&tariff.CreatedAt,
			&tariff.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Tariffs = append(resp.Tariffs, &tariff)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *StaffTariffRepo) GetById(ctx context.Context, req *staff_service.TariffIdReq) (*staff_service.StaffTariff, error) {
	var tariff = staff_service.StaffTariff{}

	query := `
    SELECT 
        id,
        name,
        type,
        amount_for_cash,
        amount_for_card,
        created_at::TEXT,
        updated_at::TEXT 
    FROM staff_tariffs 
    WHERE id = $1
    `

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&tariff.Id,
		&tariff.Name,
		&tariff.Typ,
		&tariff.AmountForCash,
		&tariff.AmountForCard,
		&tariff.CreatedAt,
		&tariff.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &tariff, nil
}

func (r *StaffTariffRepo) Update(ctx context.Context, req *staff_service.TariffUpdateReq) (*staff_service.TariffUpdateResp, error) {
	query := `
	UPDATE staff_tariffs 
	SET 
		name=$2,
		type=$3,
		amount_for_cash=$4,
		amount_for_card=$5,
		updated_at=NOW() 
	WHERE id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
		req.Name,
		req.Typ,
		req.AmountForCash,
		req.AmountForCard,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	return &staff_service.TariffUpdateResp{Msg: "OK"}, nil
}

func (r *StaffTariffRepo) Delete(ctx context.Context, req *staff_service.TariffIdReq) (*staff_service.TariffDeleteResp, error) {
	query := `
    UPDATE staff_tariffs 
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

	return &staff_service.TariffDeleteResp{Msg: "OK"}, nil
}
