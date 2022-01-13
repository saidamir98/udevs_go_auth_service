package postgres

import (
	"upm/udevs_go_auth_service/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db              *sqlx.DB
	clientPlatform  storage.ClientPlatformRepoI
	clientType      storage.ClientTypeRepoI
	client          storage.ClientRepoI
	relation        storage.RelationRepoI
	userInfoField   storage.UserInfoFieldRepoI
	role            storage.RoleRepoI
	permission      storage.PermissionRepoI
	scope           storage.ScopeRepoI
	permissionScope storage.PermissionScopeRepoI
	rolePermission  storage.RolePermissionRepoI
}

func NewPostgres(psqlConnString string) (storage.StorageI, error) {
	db, err := sqlx.Connect("postgres", psqlConnString)

	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}

func (s *Store) ClientPlatform() storage.ClientPlatformRepoI {
	if s.clientPlatform == nil {
		s.clientPlatform = NewClientPlatformRepo(s.db)
	}

	return s.clientPlatform
}

func (s *Store) ClientType() storage.ClientTypeRepoI {
	if s.clientType == nil {
		s.clientType = NewClientTypeRepo(s.db)
	}

	return s.clientType
}

func (s *Store) Client() storage.ClientRepoI {
	if s.client == nil {
		s.client = NewClientRepo(s.db)
	}

	return s.client
}

func (s *Store) Relation() storage.RelationRepoI {
	if s.relation == nil {
		s.relation = NewRelationRepo(s.db)
	}

	return s.relation
}

func (s *Store) UserInfoField() storage.UserInfoFieldRepoI {
	if s.userInfoField == nil {
		s.userInfoField = NewUserInfoFieldRepo(s.db)
	}

	return s.userInfoField
}

func (s *Store) Role() storage.RoleRepoI {
	if s.role == nil {
		s.role = NewRoleRepo(s.db)
	}

	return s.role
}

func (s *Store) Permission() storage.PermissionRepoI {
	if s.permission == nil {
		s.permission = NewPermissionRepo(s.db)
	}

	return s.permission
}

func (s *Store) Scope() storage.ScopeRepoI {
	if s.scope == nil {
		s.scope = NewScopeRepo(s.db)
	}

	return s.scope
}

func (s *Store) PermissionScope() storage.PermissionScopeRepoI {
	if s.permissionScope == nil {
		s.permissionScope = NewPermissionScopeRepo(s.db)
	}

	return s.permissionScope
}

func (s *Store) RolePermission() storage.RolePermissionRepoI {
	if s.rolePermission == nil {
		s.rolePermission = NewRolePermissionRepo(s.db)
	}

	return s.rolePermission
}
