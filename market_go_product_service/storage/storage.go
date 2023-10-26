package storage

import (
	"context"
	"genproto/product_service"
)

type StorageI interface {
	CloseDB()
	Product() ProductRepoI
	Category() CategoryRepoI
}

type ProductRepoI interface {
	Create(ctx context.Context, req *product_service.ProductCreateReq) (*product_service.ProductCreateResp, error)
	GetList(ctx context.Context, req *product_service.ProductGetListReq) (resp *product_service.ProductGetListResp, err error)
	GetById(ctx context.Context, req *product_service.ProductIdReq) (res *product_service.Product, err error)
	Update(ctx context.Context, req *product_service.ProductUpdateReq) (resp *product_service.ProductUpdateResp, err error)
	Delete(ctx context.Context, req *product_service.ProductIdReq) (*product_service.ProductDeleteResp, error)
}

type CategoryRepoI interface {
	Create(ctx context.Context, req *product_service.CategoryCreateReq) (*product_service.CategoryCreateResp, error)
	GetList(ctx context.Context, req *product_service.CategoryGetListReq) (*product_service.CategoryGetListResp, error)
	GetById(ctx context.Context, req *product_service.CategoryIdReq) (*product_service.Category, error)
	Update(ctx context.Context, req *product_service.CategoryUpdateReq) (*product_service.CategoryUpdateResp, error)
	Delete(ctx context.Context, req *product_service.CategoryIdReq) (*product_service.CategoryDeleteResp, error)
}
