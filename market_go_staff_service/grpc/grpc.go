package grpc

import (
	"genproto/staff_service"

	"gitlab.com/market_mc/market_go_staff_service/config"
	"gitlab.com/market_mc/market_go_staff_service/grpc/client"
	"gitlab.com/market_mc/market_go_staff_service/grpc/service"
	"gitlab.com/market_mc/market_go_staff_service/pkg/logger"
	"gitlab.com/market_mc/market_go_staff_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	staff_service.RegisterStaffTariffServiceServer(grpcServer, service.NewStaffTariffService(cfg, log, strg, srvc))
	staff_service.RegisterStaffServiceServer(grpcServer, service.NewStaffService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)

	return
}
