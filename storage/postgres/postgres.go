package postgres

import (
	"upm/udevs_go_auth_service/storage"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sqlx.DB

	clientPlatform storage.ClientPlatformRepoI
	clientType     storage.ClientTypeRepoI
	client         storage.ClientRepoI
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
