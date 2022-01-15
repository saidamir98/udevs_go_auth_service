package grpc

import (
	"upm/udevs_go_auth_service/config"
	"upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/genproto/ping_service"
	"upm/udevs_go_auth_service/grpc/client"
	"upm/udevs_go_auth_service/grpc/service"
	"upm/udevs_go_auth_service/pkg/logger"
	"upm/udevs_go_auth_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	ping_service.RegisterPingServiceServer(grpcServer, service.NewPingService(cfg, log, strg, svcs))
	auth_service.RegisterClientServiceServer(grpcServer, service.NewClientService(cfg, log, strg, svcs))
	auth_service.RegisterPermissionServiceServer(grpcServer, service.NewPermissionService(cfg, log, strg, svcs))
	auth_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}
