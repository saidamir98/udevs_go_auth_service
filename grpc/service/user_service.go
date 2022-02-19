package service

import (
	"context"
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

type userService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *userService {
	return &userService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {
	s.log.Info("---CreateUser--->", logger.Any("req", req))

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	req.Password = hashedPassword

	pKey, err := s.strg.User().Create(ctx, req)

	if err != nil {
		s.log.Error("!!!CreateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.User().GetByPK(ctx, pKey)
}

func (s *userService) GetUserByID(ctx context.Context, req *pb.UserPrimaryKey) (*pb.User, error) {
	s.log.Info("---GetUserByID--->", logger.Any("req", req))

	res, err := s.strg.User().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetUserByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *userService) GetUserListByIDs(ctx context.Context, req *pb.UserPrimaryKeyList) (*pb.GetUserListResponse, error) {
	s.log.Info("---GetUserListByIDs--->", logger.Any("req", req))

	res, err := s.strg.User().GetListByPKs(ctx, req)
	if err != nil {
		s.log.Error("!!!GetUserListByIDs--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *userService) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	s.log.Info("---GetUserList--->", logger.Any("req", req))

	res, err := s.strg.User().GetList(ctx, req)

	if err != nil {
		s.log.Error("!!!GetUserList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *userService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.User, error) {
	s.log.Info("---UpdateUser--->", logger.Any("req", req))

	rowsAffected, err := s.strg.User().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.User().GetByPK(ctx, &pb.UserPrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("!!!UpdateUser--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *userService) DeleteUser(ctx context.Context, req *pb.UserPrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("---DeleteUser--->", logger.Any("req", req))

	res := &emptypb.Empty{}

	rowsAffected, err := s.strg.User().Delete(ctx, req)

	if err != nil {
		s.log.Error("!!!DeleteUser--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *userService) AddUserRelation(ctx context.Context, req *pb.AddUserRelationRequest) (*pb.UserRelation, error) {
	s.log.Info("---AddUserRelation--->", logger.Any("req", req))

	pKey, err := s.strg.UserRelation().Add(ctx, req)
	if err != nil {
		s.log.Error("!!!AddUserRelation--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.UserRelation().GetByPK(ctx, pKey)
}

func (s *userService) RemoveUserRelation(ctx context.Context, req *pb.UserRelationPrimaryKey) (*pb.UserRelation, error) {
	s.log.Info("---RemoveUserRelation--->", logger.Any("req", req))

	res, err := s.strg.UserRelation().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetUserRelationPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.UserRelation().Remove(ctx, req)

	if err != nil {
		s.log.Error("!!!RemoveUserRelation--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *userService) UpsertUserInfo(ctx context.Context, req *pb.UpsertUserInfoRequest) (*pb.UserInfo, error) {
	s.log.Info("---UpsertUserInfo--->", logger.Any("req", req))

	pKey, err := s.strg.UserInfo().Upsert(ctx, req)

	if err != nil {
		s.log.Error("!!!UpsertUserInfo--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.UserInfo().GetByPK(ctx, pKey)
}
