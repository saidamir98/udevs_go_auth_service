package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jmoiron/sqlx"
)

type permissionScopeRepo struct {
	db *sqlx.DB
}

func NewPermissionScopeRepo(db *sqlx.DB) storage.PermissionScopeRepoI {
	return &permissionScopeRepo{
		db: db,
	}
}

func (r *permissionScopeRepo) Add(entity *pb.AddPermissionScopeRequest) (pKey *pb.PermissionScopePrimaryKey, err error) {
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

	_, err = r.db.Exec(query,
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

func (r *permissionScopeRepo) GetByPK(pKey *pb.PermissionScopePrimaryKey) (res *pb.PermissionScope, err error) {
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

	row, err := r.db.Query(query, pKey.PermissionId, pKey.ClientPlatformId, pKey.Path, pKey.Method)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&res.PermissionId,
			&res.ClientPlatformId,
			&res.Path,
			&res.Method,
		)

		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *permissionScopeRepo) Remove(pKey *pb.PermissionScopePrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM
		"permission_scope"
	WHERE
		permission_id = $1 AND client_platform_id = $2 AND path = $3 AND method = $4`

	result, err := r.db.Exec(query, pKey.PermissionId, pKey.ClientPlatformId, pKey.Path, pKey.Method)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
