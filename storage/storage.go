package storage

import (
	"context"
	"errors"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
)

var ErrorTheSameId = errors.New("cannot use the same uuid for 'id' and 'parent_id' fields")
var ErrorProjectId = errors.New("not valid 'project_id'")

type StorageI interface {
	CloseDB()
	ClientPlatform() ClientPlatformRepoI
	ClientType() ClientTypeRepoI
	Client() ClientRepoI
	Relation() RelationRepoI
	UserInfoField() UserInfoFieldRepoI
	Role() RoleRepoI
	Permission() PermissionRepoI
	Scope() ScopeRepoI
	PermissionScope() PermissionScopeRepoI
	RolePermission() RolePermissionRepoI
	User() UserRepoI
	Integration() IntegrationRepoI
	UserRelation() UserRelationRepoI
	UserInfo() UserInfoRepoI
	Session() SessionRepoI
}

type ClientPlatformRepoI interface {
	Create(ctx context.Context, entity *pb.CreateClientPlatformRequest) (pKey *pb.ClientPlatformPrimaryKey, err error)
	GetList(ctx context.Context, queryParam *pb.GetClientPlatformListRequest) (res *pb.GetClientPlatformListResponse, err error)
	GetByPK(ctx context.Context, pKey *pb.ClientPlatformPrimaryKey) (res *pb.ClientPlatform, err error)
	GetByPKDetailed(ctx context.Context, pKey *pb.ClientPlatformPrimaryKey) (res *pb.ClientPlatformDetailedResponse, err error)
	Update(ctx context.Context, entity *pb.UpdateClientPlatformRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, pKey *pb.ClientPlatformPrimaryKey) (rowsAffected int64, err error)
}

type ClientTypeRepoI interface {
	Create(ctx context.Context, entity *pb.CreateClientTypeRequest) (pKey *pb.ClientTypePrimaryKey, err error)
	GetList(ctx context.Context, queryParam *pb.GetClientTypeListRequest) (res *pb.GetClientTypeListResponse, err error)
	GetByPK(ctx context.Context, pKey *pb.ClientTypePrimaryKey) (res *pb.ClientType, err error)
	Update(ctx context.Context, entity *pb.UpdateClientTypeRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, pKey *pb.ClientTypePrimaryKey) (rowsAffected int64, err error)
	GetCompleteByPK(ctx context.Context, pKey *pb.ClientTypePrimaryKey) (res *pb.CompleteClientType, err error)
}

type ClientRepoI interface {
	Add(ctx context.Context, projectID string, entity *pb.AddClientRequest) (err error)
	GetByPK(ctx context.Context, entity *pb.ClientPrimaryKey) (res *pb.Client, err error)
	Update(ctx context.Context, entity *pb.UpdateClientRequest) (rowsAffected int64, err error)
	Remove(ctx context.Context, entity *pb.ClientPrimaryKey) (rowsAffected int64, err error)
	GetList(ctx context.Context, queryParam *pb.GetClientListRequest) (res *pb.GetClientListResponse, err error)
	GetMatrix(ctx context.Context, req *pb.GetClientMatrixRequest) (res *pb.GetClientMatrixResponse, err error)
}

type RelationRepoI interface {
	Add(ctx context.Context, entity *pb.AddRelationRequest) (pKey *pb.RelationPrimaryKey, err error)
	GetByPK(ctx context.Context, entity *pb.RelationPrimaryKey) (res *pb.Relation, err error)
	Update(ctx context.Context, entity *pb.UpdateRelationRequest) (rowsAffected int64, err error)
	Remove(ctx context.Context, entity *pb.RelationPrimaryKey) (rowsAffected int64, err error)
}

type UserInfoFieldRepoI interface {
	Add(ctx context.Context, entity *pb.AddUserInfoFieldRequest) (pKey *pb.UserInfoFieldPrimaryKey, err error)
	GetByPK(ctx context.Context, entity *pb.UserInfoFieldPrimaryKey) (res *pb.UserInfoField, err error)
	Update(ctx context.Context, entity *pb.UpdateUserInfoFieldRequest) (rowsAffected int64, err error)
	Remove(ctx context.Context, entity *pb.UserInfoFieldPrimaryKey) (rowsAffected int64, err error)
}

type RoleRepoI interface {
	Add(ctx context.Context, entity *pb.AddRoleRequest) (pKey *pb.RolePrimaryKey, err error)
	GetByPK(ctx context.Context, entity *pb.RolePrimaryKey) (res *pb.Role, err error)
	GetList(ctx context.Context, entity *pb.GetRolesListRequest) (res *pb.GetRolesResponse, err error)
	GetRoleByIdDetailed(ctx context.Context, entity *pb.RolePrimaryKey) (res *pb.GetRoleByIdResponse, err error)
	Update(ctx context.Context, entity *pb.UpdateRoleRequest) (rowsAffected int64, err error)
	Remove(ctx context.Context, entity *pb.RolePrimaryKey) (rowsAffected int64, err error)
}

type PermissionRepoI interface {
	Create(ctx context.Context, entity *pb.CreatePermissionRequest) (pKey *pb.PermissionPrimaryKey, err error)
	GetList(ctx context.Context, queryParam *pb.GetPermissionListRequest) (res *pb.GetPermissionListResponse, err error)
	GetByPK(ctx context.Context, pKey *pb.PermissionPrimaryKey) (res *pb.GetPermissionByIDResponse, err error)
	Update(ctx context.Context, entity *pb.UpdatePermissionRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, pKey *pb.PermissionPrimaryKey) (rowsAffected int64, err error)
	GetListByClientPlatformId(ctx context.Context, clientPlatformID string) (res []*pb.Permission, err error)
}

type ScopeRepoI interface {
	Upsert(ctx context.Context, entity *pb.UpsertScopeRequest) (res *pb.ScopePrimaryKey, err error)
	GetByPK(ctx context.Context, pKey *pb.ScopePrimaryKey) (res *pb.Scope, err error)
	GetList(ctx context.Context, queryParam *pb.GetScopeListRequest) (res *pb.GetScopesResponse, err error)
}

type PermissionScopeRepoI interface {
	Add(ctx context.Context, entity *pb.AddPermissionScopeRequest) (res *pb.PermissionScopePrimaryKey, err error)
	Remove(ctx context.Context, entity *pb.PermissionScopePrimaryKey) (rowsAffected int64, err error)
	GetByPK(ctx context.Context, pKey *pb.PermissionScopePrimaryKey) (res *pb.PermissionScope, err error)
	HasAccess(ctx context.Context, roleID, clientPlatformID, path, method string) (hasAccess bool, err error)
}

type RolePermissionRepoI interface {
	Add(ctx context.Context, entity *pb.AddRolePermissionRequest) (res *pb.RolePermissionPrimaryKey, err error)
	AddMultiple(ctx context.Context, entity *pb.AddRolePermissionsRequest) (rowsAffected int64, err error)
	Remove(ctx context.Context, entity *pb.RolePermissionPrimaryKey) (rowsAffected int64, err error)
	GetByPK(ctx context.Context, pKey *pb.RolePermissionPrimaryKey) (res *pb.RolePermission, err error)
}

type UserRepoI interface {
	GetListByPKs(ctx context.Context, pKeys *pb.UserPrimaryKeyList) (res *pb.GetUserListResponse, err error)
	Create(ctx context.Context, entity *pb.CreateUserRequest) (pKey *pb.UserPrimaryKey, err error)
	GetList(ctx context.Context, queryParam *pb.GetUserListRequest) (res *pb.GetUserListResponse, err error)
	GetByPK(ctx context.Context, pKey *pb.UserPrimaryKey) (res *pb.User, err error)
	Update(ctx context.Context, entity *pb.UpdateUserRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, pKey *pb.UserPrimaryKey) (rowsAffected int64, err error)
	GetByUsername(ctx context.Context, username string) (res *pb.User, err error)
	ResetPassword(ctx context.Context, user *pb.ResetPasswordRequest) (rowsAffected int64, err error)
}

type IntegrationRepoI interface {
	GetListByPKs(ctx context.Context, pKeys *pb.IntegrationPrimaryKeyList) (res *pb.GetIntegrationListResponse, err error)
	Create(ctx context.Context, entity *pb.CreateIntegrationRequest) (pKey *pb.IntegrationPrimaryKey, err error)
	GetList(ctx context.Context, queryParam *pb.GetIntegrationListRequest) (res *pb.GetIntegrationListResponse, err error)
	CreateSession(ctx context.Context, entity *pb.CreateSessionRequest) (pKey *pb.SessionPrimaryKey, err error)
	GetByPK(ctx context.Context, pKey *pb.IntegrationPrimaryKey) (res *pb.Integration, err error)
	Update(ctx context.Context, entity *pb.UpdateIntegrationRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, pKey *pb.IntegrationPrimaryKey) (rowsAffected int64, err error)
	GetIntegrationSessions(ctx context.Context, pKey *pb.IntegrationPrimaryKey) (res *pb.GetIntegrationSessionsResponse, err error)
	DeleteSession(ctx context.Context, pKey *pb.GetIntegrationTokenRequest) (rowsAffected int64, err error)
	GetIntegrationSession(ctx context.Context, req *pb.GetIntegrationTokenRequest) (res *pb.Session, err error)
}

type UserRelationRepoI interface {
	Add(ctx context.Context, entity *pb.AddUserRelationRequest) (res *pb.UserRelationPrimaryKey, err error)
	Remove(ctx context.Context, entity *pb.UserRelationPrimaryKey) (rowsAffected int64, err error)
	GetByPK(ctx context.Context, pKey *pb.UserRelationPrimaryKey) (res *pb.UserRelation, err error)
}

type UserInfoRepoI interface {
	Upsert(ctx context.Context, entity *pb.UpsertUserInfoRequest) (res *pb.UserInfoPrimaryKey, err error)
	GetByPK(ctx context.Context, pKey *pb.UserInfoPrimaryKey) (res *pb.UserInfo, err error)
}

type SessionRepoI interface {
	Create(ctx context.Context, entity *pb.CreateSessionRequest) (pKey *pb.SessionPrimaryKey, err error)
	GetList(ctx context.Context, queryParam *pb.GetSessionListRequest) (res *pb.GetSessionListResponse, err error)
	GetByPK(ctx context.Context, pKey *pb.SessionPrimaryKey) (res *pb.Session, err error)
	Update(ctx context.Context, entity *pb.UpdateSessionRequest) (rowsAffected int64, err error)
	Delete(ctx context.Context, pKey *pb.SessionPrimaryKey) (rowsAffected int64, err error)
	DeleteExpiredUserSessions(ctx context.Context, userID string) (rowsAffected int64, err error)
	DeleteExpiredIntegrationSessions(ctx context.Context, userID string) (rowsAffected int64, err error)
	GetSessionListByUserID(ctx context.Context, userID string) (res *pb.GetSessionListResponse, err error)
	GetSessionListByIntegrationID(ctx context.Context, userID string) (res *pb.GetSessionListResponse, err error)
}
