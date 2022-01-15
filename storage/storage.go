package storage

import (
	"errors"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
)

var ErrorNotFound = errors.New("record not found")
var ErrorTheSameId = errors.New("cannot use the same uuid for 'id' and 'parent_id' fields")

type StorageI interface {
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
	UserRelation() UserRelationRepoI
	UserInfo() UserInfoRepoI
}

type ClientPlatformRepoI interface {
	Create(entity *pb.CreateClientPlatformRequest) (pKey *pb.ClientPlatformPrimaryKey, err error)
	GetList(queryParam *pb.GetClientPlatformListRequest) (res *pb.GetClientPlatformListResponse, err error)
	GetByPK(pKey *pb.ClientPlatformPrimaryKey) (res *pb.ClientPlatform, err error)
	Update(entity *pb.UpdateClientPlatformRequest) (rowsAffected int64, err error)
	Delete(pKey *pb.ClientPlatformPrimaryKey) (rowsAffected int64, err error)
}

type ClientTypeRepoI interface {
	Create(entity *pb.CreateClientTypeRequest) (pKey *pb.ClientTypePrimaryKey, err error)
	GetList(queryParam *pb.GetClientTypeListRequest) (res *pb.GetClientTypeListResponse, err error)
	GetByPK(pKey *pb.ClientTypePrimaryKey) (res *pb.ClientType, err error)
	Update(entity *pb.UpdateClientTypeRequest) (rowsAffected int64, err error)
	Delete(pKey *pb.ClientTypePrimaryKey) (rowsAffected int64, err error)
	GetCompleteByPK(pKey *pb.ClientTypePrimaryKey) (res *pb.CompleteClientType, err error)
}

type ClientRepoI interface {
	Add(entity *pb.AddClientRequest) (err error)
	GetByPK(entity *pb.ClientPrimaryKey) (res *pb.Client, err error)
	Update(entity *pb.UpdateClientRequest) (rowsAffected int64, err error)
	Remove(entity *pb.ClientPrimaryKey) (rowsAffected int64, err error)
}

type RelationRepoI interface {
	Add(entity *pb.AddRelationRequest) (pKey *pb.RelationPrimaryKey, err error)
	GetByPK(entity *pb.RelationPrimaryKey) (res *pb.Relation, err error)
	Update(entity *pb.UpdateRelationRequest) (rowsAffected int64, err error)
	Remove(entity *pb.RelationPrimaryKey) (rowsAffected int64, err error)
}

type UserInfoFieldRepoI interface {
	Add(entity *pb.AddUserInfoFieldRequest) (pKey *pb.UserInfoFieldPrimaryKey, err error)
	GetByPK(entity *pb.UserInfoFieldPrimaryKey) (res *pb.UserInfoField, err error)
	Update(entity *pb.UpdateUserInfoFieldRequest) (rowsAffected int64, err error)
	Remove(entity *pb.UserInfoFieldPrimaryKey) (rowsAffected int64, err error)
}

type RoleRepoI interface {
	Add(entity *pb.AddRoleRequest) (pKey *pb.RolePrimaryKey, err error)
	GetByPK(entity *pb.RolePrimaryKey) (res *pb.Role, err error)
	Update(entity *pb.UpdateRoleRequest) (rowsAffected int64, err error)
	Remove(entity *pb.RolePrimaryKey) (rowsAffected int64, err error)
}

type PermissionRepoI interface {
	Create(entity *pb.CreatePermissionRequest) (pKey *pb.PermissionPrimaryKey, err error)
	GetList(queryParam *pb.GetPermissionListRequest) (res *pb.GetPermissionListResponse, err error)
	GetByPK(pKey *pb.PermissionPrimaryKey) (res *pb.Permission, err error)
	Update(entity *pb.UpdatePermissionRequest) (rowsAffected int64, err error)
	Delete(pKey *pb.PermissionPrimaryKey) (rowsAffected int64, err error)
}

type ScopeRepoI interface {
	Upsert(entity *pb.UpsertScopeRequest) (res *pb.ScopePrimaryKey, err error)
	GetByPK(pKey *pb.ScopePrimaryKey) (res *pb.Scope, err error)
}

type PermissionScopeRepoI interface {
	Add(entity *pb.AddPermissionScopeRequest) (res *pb.PermissionScopePrimaryKey, err error)
	Remove(entity *pb.PermissionScopePrimaryKey) (rowsAffected int64, err error)
	GetByPK(pKey *pb.PermissionScopePrimaryKey) (res *pb.PermissionScope, err error)
}

type RolePermissionRepoI interface {
	Add(entity *pb.AddRolePermissionRequest) (res *pb.RolePermissionPrimaryKey, err error)
	Remove(entity *pb.RolePermissionPrimaryKey) (rowsAffected int64, err error)
	GetByPK(pKey *pb.RolePermissionPrimaryKey) (res *pb.RolePermission, err error)
}

type UserRepoI interface {
	Create(entity *pb.CreateUserRequest) (pKey *pb.UserPrimaryKey, err error)
	GetList(queryParam *pb.GetUserListRequest) (res *pb.GetUserListResponse, err error)
	GetByPK(pKey *pb.UserPrimaryKey) (res *pb.User, err error)
	Update(entity *pb.UpdateUserRequest) (rowsAffected int64, err error)
	Delete(pKey *pb.UserPrimaryKey) (rowsAffected int64, err error)
}

type UserRelationRepoI interface {
	Add(entity *pb.AddUserRelationRequest) (res *pb.UserRelationPrimaryKey, err error)
	Remove(entity *pb.UserRelationPrimaryKey) (rowsAffected int64, err error)
	GetByPK(pKey *pb.UserRelationPrimaryKey) (res *pb.UserRelation, err error)
}

type UserInfoRepoI interface {
	Upsert(entity *pb.UpsertUserInfoRequest) (res *pb.UserInfoPrimaryKey, err error)
	GetByPK(pKey *pb.UserInfoPrimaryKey) (res *pb.UserInfo, err error)
}
