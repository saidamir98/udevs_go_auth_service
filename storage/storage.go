package storage

import (
	"errors"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
)

var ErrorNotFound = errors.New("record not found")

type StorageI interface {
	ClientPlatform() ClientPlatformRepoI
	ClientType() ClientTypeRepoI
	Client() ClientRepoI
	Relation() RelationRepoI
	UserInfoField() UserInfoFieldRepoI
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
