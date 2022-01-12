package client

import (
	"upm/udevs_go_auth_service/config"
	"upm/udevs_go_auth_service/genproto/settings_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	SphereService() settings_service.SphereServiceClient
	PositionService() settings_service.PositionServiceClient
}

type grpcClients struct {
	sphereService   settings_service.SphereServiceClient
	positionService settings_service.PositionServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connSettingsService, err := grpc.Dial(
		cfg.SettingsServiceHost+cfg.SettingsGRPCPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		sphereService:   settings_service.NewSphereServiceClient(connSettingsService),
		positionService: settings_service.NewPositionServiceClient(connSettingsService),
	}, nil
}

func (g *grpcClients) SphereService() settings_service.SphereServiceClient {
	return g.sphereService
}

func (g *grpcClients) PositionService() settings_service.PositionServiceClient {
	return g.positionService
}
