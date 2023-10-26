package service

import (
	"context"
	"genproto/branch_service"

	"gitlab.com/market_mc/market_go_branch_service/config"
	"gitlab.com/market_mc/market_go_branch_service/grpc/client"
	"gitlab.com/market_mc/market_go_branch_service/pkg/logger"
	"gitlab.com/market_mc/market_go_branch_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BranchService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*branch_service.UnimplementedBranchServiceServer
}

func NewBranchService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *BranchService {
	return &BranchService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *BranchService) Create(ctx context.Context, req *branch_service.BranchCreateReq) (*branch_service.BranchCreateResp, error) {
	u.log.Info("====== Branch Create ======", logger.Any("req", req))

	resp, err := u.strg.Branch().Create(ctx, req)
	if err != nil {
		u.log.Error("error while creating branch", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *BranchService) GetList(ctx context.Context, req *branch_service.BranchGetListReq) (*branch_service.BranchGetListResp, error) {
	u.log.Info("====== Branch Get List ======", logger.Any("req", req))

	resp, err := u.strg.Branch().GetList(ctx, req)
	if err != nil {
		u.log.Error("error while getting branch list", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *BranchService) GetById(ctx context.Context, req *branch_service.BranchIdReq) (*branch_service.Branch, error) {
	u.log.Info("====== Branch Get by Id ======", logger.Any("req", req))

	resp, err := u.strg.Branch().GetById(ctx, req)
	if err != nil {
		u.log.Error("error while getting branch by id", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *BranchService) Update(ctx context.Context, req *branch_service.BranchUpdateReq) (*branch_service.BranchUpdateResp, error) {
	u.log.Info("====== Branch Update ======", logger.Any("req", req))

	resp, err := u.strg.Branch().Update(ctx, req)
	if err != nil {
		u.log.Error("error while updating branch", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *BranchService) Delete(ctx context.Context, req *branch_service.BranchIdReq) (*branch_service.BranchDeleteResp, error) {
	u.log.Info("====== Branch Delete ======", logger.Any("req", req))

	resp, err := u.strg.Branch().Delete(ctx, req)
	if err != nil {
		u.log.Error("error while deleting branch", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}
