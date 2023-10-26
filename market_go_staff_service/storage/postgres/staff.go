package postgres

import (
	"context"
	"fmt"
	"genproto/staff_service"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffRepo struct {
	db *pgxpool.Pool
}

func NewStaffRepo(db *pgxpool.Pool) *StaffRepo {
	return &StaffRepo{
		db: db,
	}
}

func (r *StaffRepo) Create(ctx context.Context, req *staff_service.StaffCreateReq) (*staff_service.StaffCreateResp, error) {
	id := uuid.NewString()

	query := `
	INSERT INTO staffs (
		id,
		branch_id,
		tariff_id,
		name,
		type,
		balance,
		birth_date
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7);
	`

	_, err := r.db.Exec(ctx, query,
		id,
		req.BranchId,
		req.TariffId,
		req.Name,
		req.Typ,
		req.Balance,
		req.BirthDate,
	)

	if err != nil {
		fmt.Println("error:", err.Error())
		return nil, err
	}

	return &staff_service.StaffCreateResp{Msg: "staff created with id: " + id}, nil
}

func (r *StaffRepo) GetList(ctx context.Context, req *staff_service.StaffGetListReq) (*staff_service.StaffGetListResp, error) {
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
		tariff_id,
		name,
		type,
        balance,
        birth_date::TEXT,
		created_at::TEXT,
		updated_at::TEXT 
	FROM staffs `

	if req.Name != "" {
		filter += ` AND name ILIKE ` + "'%" + req.Name + "%' "
	}
	if req.BranchId != "" {
		filter += ` AND branch_id = ` + "'" + req.BranchId + "' "
	}
	if req.TarifId != "" {
		filter += ` AND tariff_id = ` + "'" + req.TarifId + "' "
	}
	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}
	if offset > 0 {
		offsetQ = fmt.Sprintf("OFFSET %d", offset)
	}

	query := s + filter + limit + offsetQ

	countS := `SELECT COUNT(*) FROM staffs` + filter

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	err = r.db.QueryRow(ctx, countS).Scan(&count)
	if err != nil {
		return nil, err
	}

	resp := &staff_service.StaffGetListResp{}
	for rows.Next() {
		var staff = staff_service.Staff{}
		err := rows.Scan(
			&staff.Id,
			&staff.BranchId,
			&staff.TariffId,
			&staff.Name,
			&staff.Typ,
			&staff.Balance,
			&staff.BirthDate,
			&staff.CreatedAt,
			&staff.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		resp.Staffs = append(resp.Staffs, &staff)
		resp.Count = int64(count)
	}

	return resp, nil
}

func (r *StaffRepo) GetById(ctx context.Context, req *staff_service.StaffIdReq) (*staff_service.Staff, error) {
	query := `
	SELECT 
		id,
		branch_id,
		tariff_id,
		name,
		type,
		balance,
		birth_date::TEXT,
		created_at::TEXT,
		updated_at::TEXT 
	FROM staffs 
	WHERE id=$1 AND deleted_at IS NULL;`

	var staff = staff_service.Staff{}
	if err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&staff.Id,
		&staff.BranchId,
		&staff.TariffId,
		&staff.Name,
		&staff.Typ,
		&staff.Balance,
		&staff.BirthDate,
		&staff.CreatedAt,
		&staff.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &staff, nil
}

func (r *StaffRepo) Update(ctx context.Context, req *staff_service.StaffUpdateReq) (*staff_service.StaffUpdateResp, error) {
	query := `
	UPDATE staffs 
	SET 
		branch_id=$2,
		tariff_id=$3,
		name=$4,
		balance=$5,
		birth_date=$6,
		updated_at=NOW() 
	WHERE id=$1;`

	res, err := r.db.Exec(ctx, query,
		req.Id,
		req.BranchId,
		req.TariffId,
		req.Name,
		req.Balance,
		req.BirthDate,
	)

	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, pgx.ErrNoRows
	}
	return &staff_service.StaffUpdateResp{Msg: "OK"}, nil
}

func (r *StaffRepo) Delete(ctx context.Context, req *staff_service.StaffIdReq) (*staff_service.StaffDeleteResp, error) {
	query := `
    UPDATE staffs 
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

	return &staff_service.StaffDeleteResp{Msg: "OK"}, nil
}
