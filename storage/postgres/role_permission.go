package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jmoiron/sqlx"
)

type rolePermissionRepo struct {
	db *sqlx.DB
}

func NewRolePermissionRepo(db *sqlx.DB) storage.RolePermissionRepoI {
	return &rolePermissionRepo{
		db: db,
	}
}

func (r *rolePermissionRepo) Add(ctx context.Context, entity *pb.AddRolePermissionRequest) (pKey *pb.RolePermissionPrimaryKey, err error) {
	query := `INSERT INTO "role_permission" (
		role_id,
		permission_id
	) VALUES (
		$1,
		$2
	)`

	_, err = r.db.ExecContext(ctx, query,
		entity.RoleId,
		entity.PermissionId,
	)

	pKey = &pb.RolePermissionPrimaryKey{
		RoleId:       entity.RoleId,
		PermissionId: entity.PermissionId,
	}

	return pKey, err
}

func (r *rolePermissionRepo) GetByPK(ctx context.Context, pKey *pb.RolePermissionPrimaryKey) (res *pb.RolePermission, err error) {
	res = &pb.RolePermission{}
	query := `SELECT
		role_id,
		permission_id
	FROM
		"role_permission"
	WHERE
		role_id = $1 AND permission_id = $2`

	row, err := r.db.QueryContext(ctx, query, pKey.RoleId, pKey.PermissionId)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&res.RoleId,
			&res.PermissionId,
		)

		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *rolePermissionRepo) Remove(ctx context.Context, pKey *pb.RolePermissionPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM
		"role_permission"
	WHERE
		role_id = $1 AND permission_id = $2`

	result, err := r.db.ExecContext(ctx, query, pKey.RoleId, pKey.PermissionId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
