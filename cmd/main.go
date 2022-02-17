package main

import (
	"fmt"
	"net"
	"upm/udevs_go_auth_service/api"
	"upm/udevs_go_auth_service/api/handlers"
	"upm/udevs_go_auth_service/config"
	"upm/udevs_go_auth_service/grpc"
	"upm/udevs_go_auth_service/grpc/client"
	"upm/udevs_go_auth_service/pkg/logger"
	"upm/udevs_go_auth_service/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	pgStore, err := postgres.NewPostgres(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	), cfg)
	if err != nil {
		panic(err)
	}

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.SetUpServer(cfg, log, pgStore, svcs)
	go func() {
		lis, err := net.Listen("tcp", cfg.AuthGRPCPort)
		if err != nil {
			panic(err)
		}

		log.Info("GRPC: Server being started...", logger.String("port", cfg.AuthGRPCPort))

		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	h := handlers.NewHandler(cfg, log, svcs)

	r := api.SetUpRouter(h, cfg)

	r.Run(cfg.HTTPPort)
}
