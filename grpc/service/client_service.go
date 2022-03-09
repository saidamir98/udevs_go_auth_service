package service

import (
	"context"
	"errors"
	"upm/udevs_go_auth_service/config"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/grpc/client"
	"upm/udevs_go_auth_service/pkg/logger"
	"upm/udevs_go_auth_service/storage"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type clientService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedClientServiceServer
}

func NewClientService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *clientService {
	return &clientService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *clientService) CreateClientPlatform(ctx context.Context, req *pb.CreateClientPlatformRequest) (*pb.ClientPlatform, error) {
	s.log.Info("---CreateClientPlatform--->", logger.Any("req", req))

	pKey, err := s.strg.ClientPlatform().Create(ctx, req)

	if err != nil {
		s.log.Error("!!!CreateClientPlatform--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.ClientPlatform().GetByPK(ctx, pKey)
}

func (s *clientService) GetClientPlatformByID(ctx context.Context, req *pb.ClientPlatformPrimaryKey) (*pb.ClientPlatform, error) {
	s.log.Info("---GetClientPlatformByID--->", logger.Any("req", req))

	res, err := s.strg.ClientPlatform().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetClientPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *clientService) GetClientPlatformByIDDetailed(ctx context.Context, req *pb.ClientPlatformPrimaryKey) (*pb.ClientPlatformDetailedResponse, error) {
	s.log.Info("---GetClientPlatformByID--->", logger.Any("req", req))

	res, err := s.strg.ClientPlatform().GetByPKDetailed(ctx, req)

	if err != nil {
		s.log.Error("!!!GetClientPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *clientService) GetClientPlatformList(ctx context.Context, req *pb.GetClientPlatformListRequest) (*pb.GetClientPlatformListResponse, error) {
	s.log.Info("---GetClientPlatformList--->", logger.Any("req", req))

	res, err := s.strg.ClientPlatform().GetList(ctx, req)

	if err != nil {
		s.log.Error("!!!GetClientPlatformList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *clientService) UpdateClientPlatform(ctx context.Context, req *pb.UpdateClientPlatformRequest) (*pb.ClientPlatform, error) {
	s.log.Info("---UpdateClientPlatform--->", logger.Any("req", req))

	rowsAffected, err := s.strg.ClientPlatform().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdateClientPlatform--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.ClientPlatform().GetByPK(ctx, &pb.ClientPlatformPrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("!!!UpdateClientPlatform--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *clientService) DeleteClientPlatform(ctx context.Context, req *pb.ClientPlatformPrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("---DeleteClientPlatform--->", logger.Any("req", req))

	res := &emptypb.Empty{}

	rowsAffected, err := s.strg.ClientPlatform().Delete(ctx, req)

	if err != nil {
		s.log.Error("!!!DeleteClientPlatform--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *clientService) CreateClientType(ctx context.Context, req *pb.CreateClientTypeRequest) (*pb.ClientType, error) {
	s.log.Info("---CreateClientType--->", logger.Any("req", req))

	pKey, err := s.strg.ClientType().Create(ctx, req)

	if err != nil {
		s.log.Error("!!!CreateClientType--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.ClientType().GetByPK(ctx, pKey)
}

func (s *clientService) GetClientTypeByID(ctx context.Context, req *pb.ClientTypePrimaryKey) (*pb.CompleteClientType, error) {
	s.log.Info("---GetClientTypeByID--->", logger.Any("req", req))

	res, err := s.strg.ClientType().GetCompleteByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetClientTypeByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *clientService) GetClientTypeList(ctx context.Context, req *pb.GetClientTypeListRequest) (*pb.GetClientTypeListResponse, error) {
	s.log.Info("---GetClientTypeList--->", logger.Any("req", req))

	res, err := s.strg.ClientType().GetList(ctx, req)

	if err != nil {
		s.log.Error("!!!GetClientTypeList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *clientService) UpdateClientType(ctx context.Context, req *pb.UpdateClientTypeRequest) (*pb.ClientType, error) {
	s.log.Info("---UpdateClientType--->", logger.Any("req", req))

	rowsAffected, err := s.strg.ClientType().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdateClientType--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.ClientType().GetByPK(ctx, &pb.ClientTypePrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("!!!UpdateClientType--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *clientService) DeleteClientType(ctx context.Context, req *pb.ClientTypePrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("---DeleteClientType--->", logger.Any("req", req))

	res := &emptypb.Empty{}

	rowsAffected, err := s.strg.ClientType().Delete(ctx, req)

	if err != nil {
		s.log.Error("!!!DeleteClientType--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *clientService) AddClient(ctx context.Context, req *pb.AddClientRequest) (*pb.Client, error) {
	s.log.Info("---AddClient--->", logger.Any("req", req))

	clientPlatform, err := s.strg.ClientPlatform().GetByPK(ctx, &pb.ClientPlatformPrimaryKey{Id: req.ClientPlatformId})
	if err != nil {
		s.log.Error("!!!AddClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	clientType, err := s.strg.ClientType().GetByPK(ctx, &pb.ClientTypePrimaryKey{Id: req.ClientTypeId})
	if err != nil {
		s.log.Error("!!!AddClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if clientPlatform.ProjectId != clientType.ProjectId {
		err = errors.New("mismatch between client platform&client project_id")
		s.log.Error("!!!AddClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.strg.Client().Add(ctx, clientPlatform.ProjectId, req)
	if err != nil {
		s.log.Error("!!!AddClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.Client().GetByPK(ctx, &pb.ClientPrimaryKey{
		ClientPlatformId: req.ClientPlatformId,
		ClientTypeId:     req.ClientTypeId,
	})
}

func (s *clientService) GetClientList(ctx context.Context, req *pb.GetClientListRequest) (*pb.GetClientListResponse, error) {
	s.log.Info("---GetClientList--->", logger.Any("req", req))

	res, err := s.strg.Client().GetList(ctx, req)

	if err != nil {
		s.log.Error("!!!GetClientList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *clientService) GetClientMatrix(ctx context.Context, req *pb.GetClientMatrixRequest) (*pb.GetClientMatrixResponse, error) {
	s.log.Info("---GetClientMatrix--->", logger.Any("req", req))

	res, err := s.strg.Client().GetMatrix(ctx, req)

	if err != nil {
		s.log.Error("!!!GetClientMatrix--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *clientService) UpdateClient(ctx context.Context, req *pb.UpdateClientRequest) (*pb.Client, error) {
	s.log.Info("---UpdateClient--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Client().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdateClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.Client().GetByPK(ctx, &pb.ClientPrimaryKey{
		ClientPlatformId: req.ClientPlatformId,
		ClientTypeId:     req.ClientTypeId,
	})

	if err != nil {
		s.log.Error("!!!UpdateClient--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *clientService) RemoveClient(ctx context.Context, req *pb.ClientPrimaryKey) (*pb.Client, error) {
	s.log.Info("---RemoveClient--->", logger.Any("req", req))

	res, err := s.strg.Client().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetClientPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.Client().Remove(ctx, req)

	if err != nil {
		s.log.Error("!!!RemoveClient--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *clientService) AddRelation(ctx context.Context, req *pb.AddRelationRequest) (*pb.Relation, error) {
	s.log.Info("---AddRelation--->", logger.Any("req", req))

	pKey, err := s.strg.Relation().Add(ctx, req)
	if err != nil {
		s.log.Error("!!!AddRelation--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.Relation().GetByPK(ctx, pKey)
}

func (s *clientService) UpdateRelation(ctx context.Context, req *pb.UpdateRelationRequest) (*pb.Relation, error) {
	s.log.Info("---UpdateRelation--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Relation().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdateRelation--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.Relation().GetByPK(ctx, &pb.RelationPrimaryKey{
		Id: req.Id,
	})

	if err != nil {
		s.log.Error("!!!UpdateRelation--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *clientService) RemoveRelation(ctx context.Context, req *pb.RelationPrimaryKey) (*pb.Relation, error) {
	s.log.Info("---RemoveRelation--->", logger.Any("req", req))

	res, err := s.strg.Relation().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetRelationPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.Relation().Remove(ctx, req)

	if err != nil {
		s.log.Error("!!!RemoveRelation--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *clientService) AddUserInfoField(ctx context.Context, req *pb.AddUserInfoFieldRequest) (*pb.UserInfoField, error) {
	s.log.Info("---AddUserInfoField--->", logger.Any("req", req))

	pKey, err := s.strg.UserInfoField().Add(ctx, req)
	if err != nil {
		s.log.Error("!!!AddUserInfoField--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.UserInfoField().GetByPK(ctx, pKey)
}

func (s *clientService) UpdateUserInfoField(ctx context.Context, req *pb.UpdateUserInfoFieldRequest) (*pb.UserInfoField, error) {
	s.log.Info("---UpdateUserInfoField--->", logger.Any("req", req))

	rowsAffected, err := s.strg.UserInfoField().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdateUserInfoField--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.UserInfoField().GetByPK(ctx, &pb.UserInfoFieldPrimaryKey{
		Id: req.Id,
	})

	if err != nil {
		s.log.Error("!!!UpdateUserInfoField--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *clientService) RemoveUserInfoField(ctx context.Context, req *pb.UserInfoFieldPrimaryKey) (*pb.UserInfoField, error) {
	s.log.Info("---RemoveUserInfoField--->", logger.Any("req", req))

	res, err := s.strg.UserInfoField().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetUserInfoFieldPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.UserInfoField().Remove(ctx, req)

	if err != nil {
		s.log.Error("!!!RemoveUserInfoField--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
