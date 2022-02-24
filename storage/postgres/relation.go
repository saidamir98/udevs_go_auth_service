package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type relationRepo struct {
	db *pgxpool.Pool
}

func NewRelationRepo(db *pgxpool.Pool) storage.RelationRepoI {
	return &relationRepo{
		db: db,
	}
}

func (r *relationRepo) Add(ctx context.Context, entity *pb.AddRelationRequest) (pKey *pb.RelationPrimaryKey, err error) {
	query := `INSERT INTO "relation" (
		id,
		client_type_id,
		type,
		name,
		description
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
		entity.ClientTypeId,
		entity.Type.String(),
		entity.Name,
		entity.Description,
	)

	pKey = &pb.RelationPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *relationRepo) GetByPK(ctx context.Context, pKey *pb.RelationPrimaryKey) (res *pb.Relation, err error) {
	var relationType string
	res = &pb.Relation{}
	query := `SELECT
		id,
		client_type_id,
		type,
		name,
		description
	FROM
		"relation"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ClientTypeId,
		&relationType,
		&res.Name,
		&res.Description,
	)

	res.Type = pb.RelationTypes(pb.RelationTypes_value[relationType])
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *relationRepo) Update(ctx context.Context, entity *pb.UpdateRelationRequest) (rowsAffected int64, err error) {
	query := `UPDATE "relation" SET
		client_type_id = :client_type_id,
		type = :type,
		name = :name,
		description = :description,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":             entity.Id,
		"client_type_id": entity.ClientTypeId,
		"type":           entity.Type.String(),
		"name":           entity.Name,
		"description":    entity.Description,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := r.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *relationRepo) Remove(ctx context.Context, pKey *pb.RelationPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "relation" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
