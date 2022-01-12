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
}

type ClientRepoI interface {
	Add(entity *pb.AddClientRequest) (err error)
	GetByPK(entity *pb.ClientPrimaryKey) (res *pb.Client, err error)
	Update(entity *pb.UpdateClientRequest) (rowsAffected int64, err error)
	Remove(entity *pb.ClientPrimaryKey) (rowsAffected int64, err error)
}
