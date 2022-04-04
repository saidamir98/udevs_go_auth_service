package service

import (
	"context"
	"errors"
	"time"
	"upm/udevs_go_auth_service/config"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/grpc/client"
	"upm/udevs_go_auth_service/pkg/logger"
	"upm/udevs_go_auth_service/pkg/security"
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

func (s *sessionService) AddSessionToIntegration(ctx context.Context, req *pb.AddSessionToIntegrationRequest) (*pb.AddSessionToIntegrationResponse, error) {
	res := &pb.AddSessionToIntegrationResponse{}

	if len(req.SecretKey) < 6 {
		err := errors.New("invalid key")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	integration, err := s.strg.Integration().GetByPK(ctx, &pb.IntegrationPrimaryKey{
		Id: req.IntegrationId,
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	match := integration.SecretKey == req.SecretKey
	if !match {
		err := errors.New("password is wrong")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if integration.Active < 0 {
		err := errors.New("integration is not active")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if integration.Active == 0 {
		err := errors.New("integration hasn't been activated yet")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	expiresAt, err := time.Parse(config.DatabaseTimeLayout, integration.ExpiresAt)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if expiresAt.Unix() < time.Now().Unix() {
		err := errors.New("integration has been expired")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.IntegrationFound = true
	res.Integration = integration

	clientType, err := s.strg.ClientType().GetByPK(ctx, &pb.ClientTypePrimaryKey{
		Id: integration.ClientTypeId,
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	res.ClientType = clientType

	clientPlatform, err := s.strg.ClientPlatform().GetByPK(ctx, &pb.ClientPlatformPrimaryKey{
		Id: integration.ClientPlatformId,
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.ClientPlatform = clientPlatform

	client, err := s.strg.Client().GetByPK(ctx, &pb.ClientPrimaryKey{
		ClientPlatformId: integration.ClientPlatformId,
		ClientTypeId:     integration.ClientTypeId,
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if client.LoginStrategy != pb.LoginStrategies_STANDARD {
		err := errors.New("incorrect login strategy")
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res.Integration.RoleId = integration.RoleId

	// TODO - Delete all old sessions & refresh token has this function too
	rowsAffected, err := s.strg.Session().DeleteExpiredIntegrationSessions(ctx, integration.Id)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.log.Info("Login--->DeleteExpiredIntegrationSessions", logger.Any("rowsAffected", rowsAffected))

	sessionPKey, err := s.strg.Session().Create(ctx, &pb.CreateSessionRequest{
		ProjectId:        integration.ProjectId,
		ClientPlatformId: integration.ClientPlatformId,
		ClientTypeId:     integration.ClientTypeId,
		IntegrationId:    integration.Id,
		RoleId:           integration.RoleId,
		Ip:               "0.0.0.0",
		Data:             integration.Data,
		ExpiresAt:        time.Now().Add(config.RefreshTokenExpiresInTime).Format(config.DatabaseTimeLayout),
	})
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	session, err := s.strg.Session().GetByPK(ctx, sessionPKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	res.Session = session

	// TODO - wrap in a function
	m := map[string]interface{}{
		"id":                 session.Id,
		"project_id":         session.ProjectId,
		"client_platform_id": session.ClientPlatformId,
		"client_type_id":     session.ClientTypeId,
		"integration_id":     session.IntegrationId,
		"role_id":            session.RoleId,
		"ip":                 session.Ip,
		"data":               session.Data,
	}

	accessToken, err := security.GenerateJWT(m, config.AccessTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	refreshToken, err := security.GenerateJWT(m, config.RefreshTokenExpiresInTime, s.cfg.SecretKey)
	if err != nil {
		s.log.Error("!!!Login--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	res.Token = &pb.Token{
		AccessToken:      accessToken,
		RefreshToken:     refreshToken,
		CreatedAt:        session.CreatedAt,
		UpdatedAt:        session.UpdatedAt,
		ExpiresAt:        session.ExpiresAt,
		RefreshInSeconds: int32(config.AccessTokenExpiresInTime.Seconds()),
	}

	return res, nil
}

func (s *integrationService) GetIntegrationByID(ctx context.Context, req *pb.IntegrationPrimaryKey) (*pb.Integration, error) {
	s.log.Info("---GetIntegrationByID--->", logger.Any("req", req))

	res, err := s.strg.Integration().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetIntegrationByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *integrationService) GetIntegrationToken(ctx context.Context, req *pb.GetIntegrationTokenRequest) (*pb.Token, error) {
	return nil, nil
}

func (s *integrationService) GetIntegrationSessions(ctx context.Context, req *pb.IntegrationPrimaryKey) (*pb.GetIntegrationSessionsResponse, error) {
	s.log.Info("---GetIntegrationSessions--->", logger.Any("req", req))

	res, err := s.strg.Integration().GetIntegrationSessions(ctx, req)
	if err != nil {
		s.log.Error("---GetIntegrationSessions--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, nil
}

func (s *integrationService) GetIntegrationListByIDs(ctx context.Context, req *pb.IntegrationPrimaryKeyList) (*pb.GetIntegrationListResponse, error) {
	s.log.Info("---GetIntegrationListByIDs--->", logger.Any("req", req))

	res, err := s.strg.Integration().GetListByPKs(ctx, req)
	if err != nil {
		s.log.Error("!!!GetIntegrationListByIDs--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *integrationService) GetIntegrationList(ctx context.Context, req *pb.GetIntegrationListRequest) (*pb.GetIntegrationListResponse, error) {
	s.log.Info("---GetIntegrationList--->", logger.Any("req", req))

	res, err := s.strg.Integration().GetList(ctx, req)

	if err != nil {
		s.log.Error("!!!GetIntegrationList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *integrationService) UpdateIntegration(ctx context.Context, req *pb.UpdateIntegrationRequest) (*pb.Integration, error) {
	s.log.Info("---UpdateIntegration--->", logger.Any("req", req))

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
	s.log.Info("---DeleteIntegration--->", logger.Any("req", req))

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

func (s *integrationService) DeleteSessionFromIntegration(ctx context.Context, req *pb.GetIntegrationTokenRequest) (*emptypb.Empty, error) {
	s.log.Info("---DeleteSessionFromIntegration--->", logger.Any("req", req))

	res := &emptypb.Empty{}

	rowsAffected, err := s.strg.Integration().DeleteSession(ctx, req)

	if err != nil {
		s.log.Error("!!!DeleteSessionFromIntegration--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
