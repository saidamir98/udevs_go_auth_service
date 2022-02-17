package postgres

import (
	"context"
	"fmt"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type permissionScopeRepo struct {
	db *pgxpool.Pool
}

func NewPermissionScopeRepo(db *pgxpool.Pool) storage.PermissionScopeRepoI {
	return &permissionScopeRepo{
		db: db,
	}
}

func (r *permissionScopeRepo) Add(ctx context.Context, entity *pb.AddPermissionScopeRequest) (pKey *pb.PermissionScopePrimaryKey, err error) {
	query := `INSERT INTO "permission_scope" (
		permission_id,
		client_platform_id,
		path,
		method
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	_, err = r.db.Exec(ctx, query,
		entity.PermissionId,
		entity.ClientPlatformId,
		entity.Path,
		entity.Method,
	)

	pKey = &pb.PermissionScopePrimaryKey{
		PermissionId:     entity.PermissionId,
		ClientPlatformId: entity.ClientPlatformId,
		Path:             entity.Path,
		Method:           entity.Method,
	}

	return pKey, err
}

func (r *permissionScopeRepo) GetByPK(ctx context.Context, pKey *pb.PermissionScopePrimaryKey) (res *pb.PermissionScope, err error) {
	res = &pb.PermissionScope{}
	query := `SELECT
		permission_id,
		client_platform_id,
		path,
		method
	FROM
		"permission_scope"
	WHERE
		permission_id = $1 AND client_platform_id = $2 AND path = $3 AND method = $4`

	err = r.db.QueryRow(ctx, query, pKey.PermissionId, pKey.ClientPlatformId, pKey.Path, pKey.Method).Scan(
		&res.PermissionId,
		&res.ClientPlatformId,
		&res.Path,
		&res.Method,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *permissionScopeRepo) Remove(ctx context.Context, pKey *pb.PermissionScopePrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM
		"permission_scope"
	WHERE
		permission_id = $1 AND client_platform_id = $2 AND path = $3 AND method = $4`

	result, err := r.db.Exec(ctx, query, pKey.PermissionId, pKey.ClientPlatformId, pKey.Path, pKey.Method)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *permissionScopeRepo) HasAccess(ctx context.Context, roleID, clientPlatformID, path, method string) (hasAccess bool, err error) {
	var count int32

	query := `SELECT COUNT(*) FROM
	(SELECT * FROM "role_permission" 
	WHERE role_id = $1
	) AS rp
	INNER JOIN
	(SELECT * FROM "permission_scope"
	WHERE client_platform_id = $2 AND path = $3 AND method = $4) AS ps
	ON rp.permission_id = ps.permission_id`

	err = r.db.QueryRow(ctx, query, roleID, clientPlatformID, path, method).Scan(
		&count,
	)
	if err != nil {
		return hasAccess, err
	}

	fmt.Println(count)

	if count > 0 {
		return true, nil
	}

	return false, nil
}
