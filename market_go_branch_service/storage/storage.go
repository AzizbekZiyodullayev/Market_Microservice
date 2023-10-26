package storage

import (
	"context"
	"genproto/branch_service"
)

type StorageI interface {
	CloseDB()
	Branch() BranchRepoI
	BranchProduct() BranchProductRepoI
}

type BranchRepoI interface {
	Create(ctx context.Context, req *branch_service.BranchCreateReq) (*branch_service.BranchCreateResp, error)
	GetList(ctx context.Context, req *branch_service.BranchGetListReq) (*branch_service.BranchGetListResp, error)
	GetById(ctx context.Context, req *branch_service.BranchIdReq) (*branch_service.Branch, error)
	Update(ctx context.Context, req *branch_service.BranchUpdateReq) (*branch_service.BranchUpdateResp, error)
	Delete(ctx context.Context, req *branch_service.BranchIdReq) (*branch_service.BranchDeleteResp, error)
}

type BranchProductRepoI interface {
}
