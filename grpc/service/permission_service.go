package service

import (
	"context"
	"upm/udevs_go_auth_service/config"
	"upm/udevs_go_auth_service/grpc/client"
	"upm/udevs_go_auth_service/pkg/logger"
	"upm/udevs_go_auth_service/storage"

	pb "upm/udevs_go_auth_service/genproto/auth_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type permissionService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	pb.UnimplementedPermissionServiceServer
}

func NewPermissionService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *permissionService {
	return &permissionService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (s *permissionService) AddRole(ctx context.Context, req *pb.AddRoleRequest) (*pb.Role, error) {
	s.log.Info("---AddRole--->", logger.Any("req", req))

	pKey, err := s.strg.Role().Add(ctx, req)
	if err != nil {
		s.log.Error("!!!AddRole--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.Role().GetByPK(ctx, pKey)
}

func (s *permissionService) GetRoleById(ctx context.Context, req *pb.RolePrimaryKey) (*pb.GetRoleByIdResponse, error) {
	s.log.Info("---GetRoleById--->", logger.Any("req", req))

	res, err := s.strg.Role().GetRoleByIdDetailed(ctx, req)

	if err != nil {
		return res, err
	}

	return s.strg.Role().GetRoleByIdDetailed(ctx, req)
}

func (s *permissionService) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest) (*pb.Role, error) {
	s.log.Info("---UpdateRole--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Role().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdateRole--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.Role().GetByPK(ctx, &pb.RolePrimaryKey{
		Id: req.Id,
	})

	if err != nil {
		s.log.Error("!!!UpdateRole--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *permissionService) RemoveRole(ctx context.Context, req *pb.RolePrimaryKey) (*pb.Role, error) {
	s.log.Info("---RemoveRole--->", logger.Any("req", req))

	res, err := s.strg.Role().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetRolePlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.Role().Remove(ctx, req)

	if err != nil {
		s.log.Error("!!!RemoveRole--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *permissionService) CreatePermission(ctx context.Context, req *pb.CreatePermissionRequest) (*pb.GetPermissionByIDResponse, error) {
	s.log.Info("---CreatePermission--->", logger.Any("req", req))

	pKey, err := s.strg.Permission().Create(ctx, req)

	if err != nil {
		s.log.Error("!!!CreatePermission--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.Permission().GetByPK(ctx, pKey)
}

func (s *permissionService) GetPermissionByID(ctx context.Context, req *pb.PermissionPrimaryKey) (*pb.GetPermissionByIDResponse, error) {
	s.log.Info("---GetPermissionByID--->", logger.Any("req", req))

	res, err := s.strg.Permission().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetPermissionByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, nil
}

func (s *permissionService) GetPermissionList(ctx context.Context, req *pb.GetPermissionListRequest) (*pb.GetPermissionListResponse, error) {
	s.log.Info("---GetPermissionList--->", logger.Any("req", req))

	res, err := s.strg.Permission().GetList(ctx, req)

	if err != nil {
		s.log.Error("!!!GetPermissionList--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return res, err
}

func (s *permissionService) UpdatePermission(ctx context.Context, req *pb.UpdatePermissionRequest) (*pb.GetPermissionByIDResponse, error) {
	s.log.Info("---UpdatePermission--->", logger.Any("req", req))

	rowsAffected, err := s.strg.Permission().Update(ctx, req)

	if err != nil {
		s.log.Error("!!!UpdatePermission--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	res, err := s.strg.Permission().GetByPK(ctx, &pb.PermissionPrimaryKey{Id: req.Id})
	if err != nil {
		s.log.Error("!!!UpdatePermission--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return res, err
}

func (s *permissionService) DeletePermission(ctx context.Context, req *pb.PermissionPrimaryKey) (*emptypb.Empty, error) {
	s.log.Info("---DeletePermission--->", logger.Any("req", req))

	res := &emptypb.Empty{}

	rowsAffected, err := s.strg.Permission().Delete(ctx, req)

	if err != nil {
		s.log.Error("!!!DeletePermission--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *permissionService) UpsertScope(ctx context.Context, req *pb.UpsertScopeRequest) (*pb.Scope, error) {
	s.log.Info("---UpsertScope--->", logger.Any("req", req))

	pKey, err := s.strg.Scope().Upsert(ctx, req)

	if err != nil {
		s.log.Error("!!!UpsertScope--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.Scope().GetByPK(ctx, pKey)
}

func (s *permissionService) AddPermissionScope(ctx context.Context, req *pb.AddPermissionScopeRequest) (*pb.PermissionScope, error) {
	s.log.Info("---AddPermissionScope--->", logger.Any("req", req))

	pKey, err := s.strg.PermissionScope().Add(ctx, req)
	if err != nil {
		s.log.Error("!!!AddPermissionScope--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.PermissionScope().GetByPK(ctx, pKey)
}

func (s *permissionService) RemovePermissionScope(ctx context.Context, req *pb.PermissionScopePrimaryKey) (*pb.PermissionScope, error) {
	s.log.Info("---RemovePermissionScope--->", logger.Any("req", req))

	res, err := s.strg.PermissionScope().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetPermissionScopePlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.PermissionScope().Remove(ctx, req)

	if err != nil {
		s.log.Error("!!!RemovePermissionScope--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}

func (s *permissionService) AddRolePermission(ctx context.Context, req *pb.AddRolePermissionRequest) (*pb.RolePermission, error) {
	s.log.Info("---AddRolePermission--->", logger.Any("req", req))

	pKey, err := s.strg.RolePermission().Add(ctx, req)
	if err != nil {
		s.log.Error("!!!AddRolePermission--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return s.strg.RolePermission().GetByPK(ctx, pKey)
}

func (s *permissionService) AddRolePermissions(ctx context.Context, req *pb.AddRolePermissionsRequest) (*pb.AddRolePermissionsResponse, error) {
	s.log.Info("---AddRolePermissions--->", logger.Any("req", req))

	rowsAffected, err := s.strg.RolePermission().AddMultiple(ctx, req)
	if err != nil {
		s.log.Error("!!!AddRolePermissions--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.AddRolePermissionsResponse{
		AddedRoles: rowsAffected,
	}, err
}

func (s *permissionService) RemoveRolePermission(ctx context.Context, req *pb.RolePermissionPrimaryKey) (*pb.RolePermission, error) {
	s.log.Info("---RemoveRolePermission--->", logger.Any("req", req))

	res, err := s.strg.RolePermission().GetByPK(ctx, req)

	if err != nil {
		s.log.Error("!!!GetRolePermissionPlatformByID--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	rowsAffected, err := s.strg.RolePermission().Remove(ctx, req)

	if err != nil {
		s.log.Error("!!!RemoveRolePermission--->", logger.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	return res, nil
}
