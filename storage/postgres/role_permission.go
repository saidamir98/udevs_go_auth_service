package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type rolePermissionRepo struct {
	db *pgxpool.Pool
}

func NewRolePermissionRepo(db *pgxpool.Pool) storage.RolePermissionRepoI {
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

	_, err = r.db.Exec(ctx, query,
		entity.RoleId,
		entity.PermissionId,
	)

	pKey = &pb.RolePermissionPrimaryKey{
		RoleId:       entity.RoleId,
		PermissionId: entity.PermissionId,
	}

	return pKey, err
}

func (r *rolePermissionRepo) AddMultiple(ctx context.Context, entity *pb.AddRolePermissionsRequest) (rowsAffected int64, err error) {
	query := `INSERT INTO "role_permission" (
		role_id,
		permission_id
	) VALUES (
		$1,
		$2
	)`

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return rowsAffected, err
	}

	// call function to commit or rollback transaction at the end
	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	for _, value := range entity.GetPermissions() {
		result, err := tx.Exec(ctx, query,
			value.RoleId,
			value.PermissionId,
		)
		if err != nil {
			return rowsAffected, err
		}

		rowsAffected += result.RowsAffected()
	}

	return rowsAffected, err
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

	err = r.db.QueryRow(ctx, query, pKey.RoleId, pKey.PermissionId).Scan(
		&res.RoleId,
		&res.PermissionId,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *rolePermissionRepo) Remove(ctx context.Context, pKey *pb.RolePermissionPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM
		"role_permission"
	WHERE
		role_id = $1 AND permission_id = $2`

	result, err := r.db.Exec(ctx, query, pKey.RoleId, pKey.PermissionId)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
