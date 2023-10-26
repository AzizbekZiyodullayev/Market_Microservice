package storage

import (
	"context"
	"genproto/staff_service"
)

type StorageI interface {
	CloseDB()
	StaffTariff() StaffTariffRepoI
	Staff() StaffRepoI
}

type StaffTariffRepoI interface {
	Create(ctx context.Context, req *staff_service.TariffCreateReq) (*staff_service.TariffCreateResp, error)
	GetList(ctx context.Context, req *staff_service.TariffGetListReq) (*staff_service.TariffGetListResp, error)
	GetById(ctx context.Context, req *staff_service.TariffIdReq) (*staff_service.StaffTariff, error)
	Update(ctx context.Context, req *staff_service.TariffUpdateReq) (*staff_service.TariffUpdateResp, error)
	Delete(ctx context.Context, req *staff_service.TariffIdReq) (*staff_service.TariffDeleteResp, error)
}

type StaffRepoI interface {
	Create(ctx context.Context, req *staff_service.StaffCreateReq) (*staff_service.StaffCreateResp, error)
	GetList(ctx context.Context, req *staff_service.StaffGetListReq) (*staff_service.StaffGetListResp, error)
	GetById(ctx context.Context, req *staff_service.StaffIdReq) (*staff_service.Staff, error)
	Update(ctx context.Context, req *staff_service.StaffUpdateReq) (*staff_service.StaffUpdateResp, error)
	Delete(ctx context.Context, req *staff_service.StaffIdReq) (*staff_service.StaffDeleteResp, error)
}
