package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type scopeRepo struct {
	db *pgxpool.Pool
}

func NewScopeRepo(db *pgxpool.Pool) storage.ScopeRepoI {
	return &scopeRepo{
		db: db,
	}
}

func (r *scopeRepo) Upsert(ctx context.Context, entity *pb.UpsertScopeRequest) (pKey *pb.ScopePrimaryKey, err error) {
	query := `INSERT INTO "scope" (
		client_platform_id,
		path,
		method,
		requests
	) values (
		$1,
		$2,
		$3,
		$4
	) ON CONFLICT (
		client_platform_id,
		path,
		method
	) DO UPDATE SET requests = "scope".requests + $4, updated_at = NOW()`

	_, err = r.db.Exec(ctx, query,
		entity.ClientPlatformId,
		entity.Path,
		entity.Method,
		1,
	)

	pKey = &pb.ScopePrimaryKey{
		ClientPlatformId: entity.ClientPlatformId,
		Path:             entity.Path,
		Method:           entity.Method,
	}

	return pKey, err
}

func (r *scopeRepo) GetByPK(ctx context.Context, pKey *pb.ScopePrimaryKey) (res *pb.Scope, err error) {
	res = &pb.Scope{}
	query := `SELECT
		client_platform_id,
		path,
		method,
		requests
	FROM
		"scope"
	WHERE
		client_platform_id = $1 AND path = $2 AND method = $3`

	err = r.db.QueryRow(ctx, query, pKey.ClientPlatformId, pKey.Path, pKey.Method).Scan(
		&res.ClientPlatformId,
		&res.Path,
		&res.Method,
		&res.Requests,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}
