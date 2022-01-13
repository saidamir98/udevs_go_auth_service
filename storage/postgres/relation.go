package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type relationRepo struct {
	db *sqlx.DB
}

func NewRelationRepo(db *sqlx.DB) storage.RelationRepoI {
	return &relationRepo{
		db: db,
	}
}

func (r *relationRepo) Add(entity *pb.AddRelationRequest) (pKey *pb.RelationPrimaryKey, err error) {
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

	_, err = r.db.Exec(query,
		uuid,
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

func (r *relationRepo) GetByPK(pKey *pb.RelationPrimaryKey) (res *pb.Relation, err error) {
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

	row, err := r.db.Query(query, pKey.Id)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		var relationType string

		err = row.Scan(
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
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *relationRepo) Update(entity *pb.UpdateRelationRequest) (rowsAffected int64, err error) {
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

	result, err := r.db.NamedExec(query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *relationRepo) Remove(pKey *pb.RelationPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "relation" WHERE id = $1`

	result, err := r.db.Exec(query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
