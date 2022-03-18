package client

import (
	"upm/udevs_go_auth_service/config"
	"upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/genproto/settings_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	SphereService() settings_service.SphereServiceClient
	PositionService() settings_service.PositionServiceClient
	IntegrationService() auth_service.IntegrationServiceClient
	ClientService() auth_service.ClientServiceClient
	PermissionService() auth_service.PermissionServiceClient
	UserService() auth_service.UserServiceClient
	SessionService() auth_service.SessionServiceClient
}

type grpcClients struct {
	sphereService      settings_service.SphereServiceClient
	positionService    settings_service.PositionServiceClient
	integrationService auth_service.IntegrationServiceClient
	clientService      auth_service.ClientServiceClient
	permissionService  auth_service.PermissionServiceClient
	userService        auth_service.UserServiceClient
	sessionService     auth_service.SessionServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connSettingsService, err := grpc.Dial(
		cfg.SettingsServiceHost+cfg.SettingsGRPCPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	connAuthService, err := grpc.Dial(
		cfg.AuthServiceHost+cfg.AuthGRPCPort,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	return &grpcClients{
		sphereService:      settings_service.NewSphereServiceClient(connSettingsService),
		positionService:    settings_service.NewPositionServiceClient(connSettingsService),
		clientService:      auth_service.NewClientServiceClient(connAuthService),
		permissionService:  auth_service.NewPermissionServiceClient(connAuthService),
		userService:        auth_service.NewUserServiceClient(connAuthService),
		sessionService:     auth_service.NewSessionServiceClient(connAuthService),
		integrationService: auth_service.NewIntegrationServiceClient(connAuthService),
	}, nil
}

func (g *grpcClients) SphereService() settings_service.SphereServiceClient {
	return g.sphereService
}

func (g *grpcClients) PositionService() settings_service.PositionServiceClient {
	return g.positionService
}

func (g *grpcClients) ClientService() auth_service.ClientServiceClient {
	return g.clientService
}

func (g *grpcClients) PermissionService() auth_service.PermissionServiceClient {
	return g.permissionService
}

func (g *grpcClients) UserService() auth_service.UserServiceClient {
	return g.userService
}

func (g *grpcClients) SessionService() auth_service.SessionServiceClient {
	return g.sessionService
}

func (g *grpcClients) IntegrationService() auth_service.IntegrationServiceClient {
	return g.integrationService
}
