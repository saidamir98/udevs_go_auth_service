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

	pKey, err := s.strg.ClientPlatform().Create(req)

	if err != nil {
		s.log.Error("!!!CreateClientPlatform--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.ClientPlatform().GetByPK(pKey)
}

func (s *clientService) GetClientPlatformByID(ctx context.Context, req *pb.ClientPlatformPrimaryKey) (*pb.ClientPlatform, error) {
	s.log.Info("---GetClientPlatformByID--->", logger.Any("req", req))

	res, err := s.strg.ClientPlatform().GetByPK(req)

	if err != nil {
		s.log.Error("!!!GetClientPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *clientService) GetClientPlatformList(ctx context.Context, req *pb.GetClientPlatformListRequest) (*pb.GetClientPlatformListResponse, error) {
	s.log.Info("---GetClientPlatformList--->", logger.Any("req", req))

	res, err := s.strg.ClientPlatform().GetList(req)

	if err != nil {
		s.log.Error("!!!GetClientPlatformList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *clientService) UpdateClientPlatform(ctx context.Context, req *pb.UpdateClientPlatformRequest) (*pb.ClientPlatform, error) {
	s.log.Info("---UpdateClientPlatform--->", logger.Any("req", req))

	rowsAffected, err := s.strg.ClientPlatform().Update(req)

	if err != nil {
		s.log.Error("!!!UpdateClientPlatform--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.ClientPlatform().GetByPK(&pb.ClientPlatformPrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("!!!UpdateClientPlatform--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *clientService) DeleteClientPlatform(ctx context.Context, req *pb.ClientPlatformPrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("---DeleteClientPlatform--->", logger.Any("req", req))

	res := &emptypb.Empty{}

	rowsAffected, err := s.strg.ClientPlatform().Delete(req)

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

	pKey, err := s.strg.ClientType().Create(req)

	if err != nil {
		s.log.Error("!!!CreateClientType--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.ClientType().GetByPK(pKey)
}

func (s *clientService) GetClientTypeByID(ctx context.Context, req *pb.ClientTypePrimaryKey) (*pb.CompleteClientType, error) {
	s.log.Info("---GetClientTypeByID--->", logger.Any("req", req))

	res, err := s.strg.ClientType().GetCompleteByPK(req)

	if err != nil {
		s.log.Error("!!!GetClientTypeByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *clientService) GetClientTypeList(ctx context.Context, req *pb.GetClientTypeListRequest) (*pb.GetClientTypeListResponse, error) {
	s.log.Info("---GetClientTypeList--->", logger.Any("req", req))

	res, err := s.strg.ClientType().GetList(req)

	if err != nil {
		s.log.Error("!!!GetClientTypeList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *clientService) UpdateClientType(ctx context.Context, req *pb.UpdateClientTypeRequest) (*pb.ClientType, error) {
	s.log.Info("---UpdateClientType--->", logger.Any("req", req))

	rowsAffected, err := s.strg.ClientType().Update(req)

	if err != nil {
		s.log.Error("!!!UpdateClientType--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.ClientType().GetByPK(&pb.ClientTypePrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("!!!UpdateClientType--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *clientService) DeleteClientType(ctx context.Context, req *pb.ClientTypePrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("---DeleteClientType--->", logger.Any("req", req))

	res := &emptypb.Empty{}

	rowsAffected, err := s.strg.ClientType().Delete(req)

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

	err := s.strg.Client().Add(req)
	if err != nil {
		s.log.Error("!!!AddClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.Client().GetByPK(&pb.ClientPrimaryKey{
		ClientPlatformId: req.ClientPlatformId,
		ClientTypeId:     req.ClientTypeId,
	})
}

func (s *clientService) UpdateClient(ctx context.Context, req *pb.UpdateClientRequest) (*pb.Client, error) {
	s.log.Info("---UpdateClient--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Client().Update(req)

	if err != nil {
		s.log.Error("!!!UpdateClient--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.Client().GetByPK(&pb.ClientPrimaryKey{
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

	res, err := s.strg.Client().GetByPK(req)

	if err != nil {
		s.log.Error("!!!GetClientPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.Client().Remove(req)

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

	pKey, err := s.strg.Relation().Add(req)
	if err != nil {
		s.log.Error("!!!AddRelation--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.Relation().GetByPK(pKey)
}

func (s *clientService) UpdateRelation(ctx context.Context, req *pb.UpdateRelationRequest) (*pb.Relation, error) {
	s.log.Info("---UpdateRelation--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Relation().Update(req)

	if err != nil {
		s.log.Error("!!!UpdateRelation--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.Relation().GetByPK(&pb.RelationPrimaryKey{
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

	res, err := s.strg.Relation().GetByPK(req)

	if err != nil {
		s.log.Error("!!!GetRelationPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.Relation().Remove(req)

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

	pKey, err := s.strg.UserInfoField().Add(req)
	if err != nil {
		s.log.Error("!!!AddUserInfoField--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.UserInfoField().GetByPK(pKey)
}

func (s *clientService) UpdateUserInfoField(ctx context.Context, req *pb.UpdateUserInfoFieldRequest) (*pb.UserInfoField, error) {
	s.log.Info("---UpdateUserInfoField--->", logger.Any("req", req))

	rowsAffected, err := s.strg.UserInfoField().Update(req)

	if err != nil {
		s.log.Error("!!!UpdateUserInfoField--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.UserInfoField().GetByPK(&pb.UserInfoFieldPrimaryKey{
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

	res, err := s.strg.UserInfoField().GetByPK(req)

	if err != nil {
		s.log.Error("!!!GetUserInfoFieldPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.UserInfoField().Remove(req)

	if err != nil {
		s.log.Error("!!!RemoveUserInfoField--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
