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
		name,
		client_platform_id,
		project_id
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return pKey, err
	}

	_, err = r.db.Exec(ctx, query,
		uuid.String(),
		entity.GetClientTypeId(),
		entity.GetName(),
		entity.GetClientPlatformId(),
		entity.GetProjectId(),
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
		name,
		client_platform_id,
		project_id
	FROM
		"role"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ClientTypeId,
		&res.Name,
		&res.ClientPlatformId,
		&res.ProjectId,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *roleRepo) GetRoleByIdDetailed(ctx context.Context, entity *pb.RolePrimaryKey) (res *pb.GetRoleByIdResponse, err error) {
	res = &pb.GetRoleByIdResponse{}
	var confirmBy string
	query := `SELECT 
		rl.id,
		rl.client_type_id,
		rl.name,
		-- rl.client_platform_id,
		-- rl.project_id,
		ct.id AS client_type_id,
		ct.confirm_by,
		ct.project_id,
		ct.self_register,
		ct.self_recover,
		ct.name AS client_type_name
	FROM 
		"role" 
	AS
		rl
	
	INNER JOIN
		"client_type" AS ct
	
	ON
		ct.id = rl.client_type_id
	WHERE
		rl.id = $1`

	res.ClientType = new(pb.ClientType)
	err = r.db.QueryRow(ctx, query, entity.Id).Scan(
		&res.Id,
		&res.ClientTypeId,
		&res.Name,
		&res.ClientType.Id,
		&confirmBy,
		&res.ClientType.ProjectId,
		&res.ClientType.SelfRegister,
		&res.ClientType.SelfRecover,
		&res.ClientType.Name,
	)
	res.ClientType.ConfirmBy = pb.ConfirmStrategies(pb.ConfirmStrategies_value[confirmBy])

	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *roleRepo) Update(ctx context.Context, entity *pb.UpdateRoleRequest) (rowsAffected int64, err error) {
	query := `UPDATE "role" SET
		client_type_id = :client_type_id,
		name = :name,
		client_platform_id = :client_platform_id,
		project_id = :project_id,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":                 entity.Id,
		"client_type_id":     entity.ClientTypeId,
		"name":               entity.Name,
		"client_platform_id": entity.ClientPlatformId,
		"project_id":         entity.ProjectId,
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
