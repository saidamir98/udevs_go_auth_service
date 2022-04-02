package service

import (
	"context"
	"upm/udevs_go_auth_service/config"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/grpc/client"
	"upm/udevs_go_auth_service/pkg/logger"
	"upm/udevs_go_auth_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type integrationService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedIntegrationServiceServer
}

func NewIntegrationService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *integrationService {
	return &integrationService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *integrationService) CreateIntegration(ctx context.Context, req *pb.CreateIntegrationRequest) (*pb.Integration, error) {
	s.log.Info("---CreateUser--->", logger.Any("req", req))

	pKey, err := s.strg.Integration().Create(ctx, req)

	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.Integration().GetByPK(ctx, pKey)
}

func (s *integrationService) GetIntegrationByID(ctx context.Context, req *pb.IntegrationPrimaryKey) (*pb.Integration, error) {
	s.log.Info("---GetUserByID--->", logger.Any("req", req))

	res, err := s.strg.Integration().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetUserByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *integrationService) GetIntegrationSessions(ctx context.Context, req *pb.IntegrationPrimaryKey) (*pb.GetIntegrationSessionsResponse, error) {
	s.log.Info("---GetIntegrationByUserId--->", logger.Any("req", req))

	res, err := s.strg.Integration().GetIntegrationSessions(ctx, req)
	if err != nil {
		s.log.Error("---GetIntegrationByUserId--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *integrationService) GetIntegrationListByIDs(ctx context.Context, req *pb.IntegrationPrimaryKeyList) (*pb.GetIntegrationListResponse, error) {
	s.log.Info("---GetUserListByIDs--->", logger.Any("req", req))

	res, err := s.strg.Integration().GetListByPKs(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUserListByIDs--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *integrationService) GetIntegrationList(ctx context.Context, req *pb.GetIntegrationListRequest) (*pb.GetIntegrationListResponse, error) {
	s.log.Info("---GetUserList--->", logger.Any("req", req))

	res, err := s.strg.Integration().GetList(ctx, req)

	if err != nil {
		s.log.Error("!!!GetUserList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *integrationService) UpdateIntegration(ctx context.Context, req *pb.UpdateIntegrationRequest) (*pb.Integration, error) {
	s.log.Info("---UpdateUser--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Integration().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdateIntegration--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.Integration().GetByPK(ctx, &pb.IntegrationPrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("!!!UpdateIntegration--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *integrationService) DeleteIntegration(ctx context.Context, req *pb.IntegrationPrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("---DeleteUser--->", logger.Any("req", req))

	res := &emptypb.Empty{}

	rowsAffected, err := s.strg.Integration().Delete(ctx, req)

	if err != nil {
		s.log.Error("!!!DeleteIntegration--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
