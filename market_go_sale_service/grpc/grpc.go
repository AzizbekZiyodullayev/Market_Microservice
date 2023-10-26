package grpc

import (
	"genproto/sale_service"

	"gitlab.com/market_mc/market_go_sale_service/config"

	"gitlab.com/market_mc/market_go_sale_service/grpc/client"
	"gitlab.com/market_mc/market_go_sale_service/grpc/service"
	"gitlab.com/market_mc/market_go_sale_service/pkg/logger"
	"gitlab.com/market_mc/market_go_sale_service/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvc client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	sale_service.RegisterSaleServiceServer(grpcServer, service.NewSaleService(cfg, log, strg, srvc))
	sale_service.RegisterSaleProductServiceServer(grpcServer, service.NewSaleProductService(cfg, log, strg, srvc))
	sale_service.RegisterStaffTransactionServiceServer(grpcServer, service.NewStaffTransactionService(cfg, log, strg, srvc))

	reflection.Register(grpcServer)

	return
}
