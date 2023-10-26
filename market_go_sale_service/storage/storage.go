package storage

import (
	"context"
	"genproto/sale_service"
)

type StorageI interface {
	CloseDB()
	Sale() SaleRepoI
	SaleProduct() SaleProductRepoI
	StaffTransaction() StaffTransactionRepoI
	BranchProductTransaction() BranchProductTransactionRepoI
}

type SaleRepoI interface {
	Create(ctx context.Context, req *sale_service.SaleCreateReq) (*sale_service.SaleCreateResp, error)
	GetList(ctx context.Context, req *sale_service.SaleGetListReq) (*sale_service.SaleGetListResp, error)
	GetById(ctx context.Context, req *sale_service.SaleIdReq) (*sale_service.Sale, error)
	Update(ctx context.Context, req *sale_service.SaleUpdateReq) (*sale_service.SaleUpdateResp, error)
	Delete(ctx context.Context, req *sale_service.SaleIdReq) (*sale_service.SaleDeleteResp, error)
}

type SaleProductRepoI interface {
	Create(ctx context.Context, req *sale_service.SaleProductCreateReq) (*sale_service.SaleProductCreateResp, error)
	GetList(ctx context.Context, req *sale_service.SaleProductGetListReq) (*sale_service.SaleProductGetListResp, error)
	GetById(ctx context.Context, req *sale_service.SaleProductIdReq) (*sale_service.SaleProduct, error)
	Update(ctx context.Context, req *sale_service.SaleProductUpdateReq) (*sale_service.SaleProductUpdateResp, error)
	Delete(ctx context.Context, req *sale_service.SaleProductIdReq) (*sale_service.SaleProductDeleteResp, error)
}

type StaffTransactionRepoI interface {
	Create(ctx context.Context, req *sale_service.StaffTrCreateReq) (*sale_service.StaffTrCreateResp, error)
	GetList(ctx context.Context, req *sale_service.StaffTrGetListReq) (*sale_service.StaffTrGetListResp, error)
	GetById(ctx context.Context, req *sale_service.StaffTrIdReq) (*sale_service.StaffTransaction, error)
	Update(ctx context.Context, req *sale_service.StaffTrUpdateReq) (*sale_service.StaffTrUpdateResp, error)
	Delete(ctx context.Context, req *sale_service.StaffTrIdReq) (*sale_service.StaffTrDeleteResp, error)
}

type BranchProductTransactionRepoI interface {
	Create(ctx context.Context, req *sale_service.BrPrTrCreateReq) (*sale_service.BrPrTrCreateResp, error)
	GetList(ctx context.Context, req *sale_service.BrPrTrGetListReq) (*sale_service.BrPrTrGetListResp, error)
	GetById(ctx context.Context, req *sale_service.BrPrTrIdReq) (*sale_service.BrPrTransaction, error)
	Update(ctx context.Context, req *sale_service.BrPrTrUpdateReq) (*sale_service.BrPrTrUpdateResp, error)
	Delete(ctx context.Context, req *sale_service.BrPrTrIdReq) (*sale_service.BrPrTrDeleteResp, error)
}
