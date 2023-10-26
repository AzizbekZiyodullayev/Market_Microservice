package client

import "gitlab.com/market_mc/market_go_sale_service/config"

type ServiceManagerI interface {
}

type grpcClients struct {
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	return &grpcClients{}, nil
}
