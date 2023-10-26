package service

import (
	"context"
	"genproto/staff_service"

	"gitlab.com/market_mc/market_go_staff_service/config"
	"gitlab.com/market_mc/market_go_staff_service/grpc/client"
	"gitlab.com/market_mc/market_go_staff_service/pkg/logger"
	"gitlab.com/market_mc/market_go_staff_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StaffTariffService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	*staff_service.UnimplementedStaffTariffServiceServer
}

func NewStaffTariffService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) *StaffTariffService {
	return &StaffTariffService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvc,
	}
}

func (u *StaffTariffService) Create(ctx context.Context, req *staff_service.TariffCreateReq) (*staff_service.TariffCreateResp, error) {
	u.log.Info("====== StaffTariff Create ======", logger.Any("req", req))

	resp, err := u.strg.StaffTariff().Create(ctx, req)
	if err != nil {
		u.log.Error("error while creating staff tariff", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffTariffService) GetList(ctx context.Context, req *staff_service.TariffGetListReq) (*staff_service.TariffGetListResp, error) {
	u.log.Info("====== StaffTariff GetList ======", logger.Any("req", req))

	resp, err := u.strg.StaffTariff().GetList(ctx, req)
	if err != nil {
		u.log.Error("error while getting staff tariffs", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffTariffService) GetById(ctx context.Context, req *staff_service.TariffIdReq) (*staff_service.StaffTariff, error) {
	u.log.Info("====== StaffTariff GetById ======", logger.Any("req", req))

	resp, err := u.strg.StaffTariff().GetById(ctx, req)
	if err != nil {
		u.log.Error("error while getting staff tariff", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffTariffService) Update(ctx context.Context, req *staff_service.TariffUpdateReq) (*staff_service.TariffUpdateResp, error) {
	u.log.Info("====== StaffTariff Update ======", logger.Any("req", req))

	resp, err := u.strg.StaffTariff().Update(ctx, req)
	if err != nil {
		u.log.Error("error while updating staff tariff", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (u *StaffTariffService) Delete(ctx context.Context, req *staff_service.TariffIdReq) (*staff_service.TariffDeleteResp, error) {
	u.log.Info("====== StaffTariff Delete ======", logger.Any("req", req))

	resp, err := u.strg.StaffTariff().Delete(ctx, req)
	if err != nil {
		u.log.Error("error while deleting staff tariff", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}
