package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type roleRepo struct {
	db *pgxpool.Pool
}

func NewRoleRepo(db *pgxpool.Pool) storage.RoleRepoI {
	return &roleRepo{
		db: db,
	}
}

func (r *roleRepo) Add(ctx context.Context, entity *pb.AddRoleRequest) (pKey *pb.RolePrimaryKey, err error) {
	query := `INSERT INTO "role" (
		id,
		client_type_id,
		name
	) VALUES (
		$1,
		$2,
		$3
	)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return pKey, err
	}

	_, err = r.db.Exec(ctx, query,
		uuid,
		entity.ClientTypeId,
		entity.Name,
	)

	pKey = &pb.RolePrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *roleRepo) GetByPK(ctx context.Context, pKey *pb.RolePrimaryKey) (res *pb.Role, err error) {
	res = &pb.Role{}
	query := `SELECT
		id,
		client_type_id,
		name
	FROM
		"role"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ClientTypeId,
		&res.Name,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *roleRepo) Update(ctx context.Context, entity *pb.UpdateRoleRequest) (rowsAffected int64, err error) {
	query := `UPDATE "role" SET
		client_type_id = :client_type_id,
		name = :name,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":             entity.Id,
		"client_type_id": entity.ClientTypeId,
		"name":           entity.Name,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := r.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *roleRepo) Remove(ctx context.Context, pKey *pb.RolePrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "role" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
