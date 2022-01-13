package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type roleRepo struct {
	db *sqlx.DB
}

func NewRoleRepo(db *sqlx.DB) storage.RoleRepoI {
	return &roleRepo{
		db: db,
	}
}

func (r *roleRepo) Add(entity *pb.AddRoleRequest) (pKey *pb.RolePrimaryKey, err error) {
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

	_, err = r.db.Exec(query,
		uuid,
		entity.ClientTypeId,
		entity.Name,
	)

	pKey = &pb.RolePrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *roleRepo) GetByPK(pKey *pb.RolePrimaryKey) (res *pb.Role, err error) {
	res = &pb.Role{}
	query := `SELECT
		id,
		client_type_id,
		name
	FROM
		"role"
	WHERE
		id = $1`

	row, err := r.db.Query(query, pKey.Id)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {

		err = row.Scan(
			&res.Id,
			&res.ClientTypeId,
			&res.Name,
		)

		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *roleRepo) Update(entity *pb.UpdateRoleRequest) (rowsAffected int64, err error) {
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

func (r *roleRepo) Remove(pKey *pb.RolePrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "role" WHERE id = $1`

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
