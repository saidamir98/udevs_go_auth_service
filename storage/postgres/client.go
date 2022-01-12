package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jmoiron/sqlx"
)

type clientRepo struct {
	db *sqlx.DB
}

func NewClientRepo(db *sqlx.DB) storage.ClientRepoI {
	return &clientRepo{
		db: db,
	}
}

func (r *clientRepo) Add(entity *pb.AddClientRequest) (err error) {
	query := `INSERT INTO "client" (
		client_platform_id,
		client_type_id,
		login_strategy
	) VALUES (
		$1,
		$2,
		$3
	)`

	_, err = r.db.Exec(query,
		entity.ClientPlatformId,
		entity.ClientTypeId,
		entity.LoginStrategy.String(),
	)

	return err
}

func (r *clientRepo) GetByPK(pKey *pb.ClientPrimaryKey) (res *pb.Client, err error) {
	res = &pb.Client{}
	query := `SELECT
		client_platform_id,
		client_type_id,
		login_strategy
	FROM
		"client"
	WHERE
		client_platform_id = $1 AND client_type_id = $2`

	row, err := r.db.Query(query, pKey.ClientPlatformId, pKey.ClientTypeId)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		var loginStrategy string

		err = row.Scan(
			&res.ClientPlatformId,
			&res.ClientTypeId,
			&loginStrategy,
		)

		res.LoginStrategy = pb.LoginStrategies(pb.LoginStrategies_value[loginStrategy])

		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *clientRepo) Update(entity *pb.UpdateClientRequest) (rowsAffected int64, err error) {
	query := `UPDATE "client" SET
		login_strategy = :login_strategy,
		updated_at = now()
	WHERE
		client_platform_id = :client_platform_id AND client_type_id = :client_type_id`

	params := map[string]interface{}{
		"client_platform_id": entity.ClientPlatformId,
		"client_type_id":     entity.ClientTypeId,
		"login_strategy":     entity.LoginStrategy.String(),
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

func (r *clientRepo) Remove(pKey *pb.ClientPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "client" WHERE client_platform_id = $1 AND client_type_id = $2`

	result, err := r.db.Exec(query, pKey.ClientPlatformId, pKey.ClientTypeId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
