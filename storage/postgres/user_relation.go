package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jmoiron/sqlx"
)

type userRelationRepo struct {
	db *sqlx.DB
}

func NewUserRelationRepo(db *sqlx.DB) storage.UserRelationRepoI {
	return &userRelationRepo{
		db: db,
	}
}

func (r *userRelationRepo) Add(entity *pb.AddUserRelationRequest) (pKey *pb.UserRelationPrimaryKey, err error) {
	query := `INSERT INTO "user_relation" (
		user_id,
		relation_id
	) VALUES (
		$1,
		$2
	)`

	_, err = r.db.Exec(query,
		entity.UserId,
		entity.RelationId,
	)

	pKey = &pb.UserRelationPrimaryKey{
		UserId:     entity.UserId,
		RelationId: entity.RelationId,
	}

	return pKey, err
}

func (r *userRelationRepo) GetByPK(pKey *pb.UserRelationPrimaryKey) (res *pb.UserRelation, err error) {
	res = &pb.UserRelation{}
	query := `SELECT
		user_id,
		relation_id
	FROM
		"user_relation"
	WHERE
		user_id = $1 AND relation_id = $2`

	row, err := r.db.Query(query, pKey.UserId, pKey.RelationId)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&res.UserId,
			&res.RelationId,
		)

		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *userRelationRepo) Remove(pKey *pb.UserRelationPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM
		"user_relation"
	WHERE
		user_id = $1 AND relation_id = $2`

	result, err := r.db.Exec(query, pKey.UserId, pKey.RelationId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
