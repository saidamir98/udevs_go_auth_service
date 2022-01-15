package postgres

import (
	"encoding/json"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jmoiron/sqlx"
)

type userInfoRepo struct {
	db *sqlx.DB
}

func NewUserInfoRepo(db *sqlx.DB) storage.UserInfoRepoI {
	return &userInfoRepo{
		db: db,
	}
}

func (r *userInfoRepo) Upsert(entity *pb.UpsertUserInfoRequest) (pKey *pb.UserInfoPrimaryKey, err error) {
	data, err := json.Marshal(entity.Data)
	if err != nil {
		return pKey, err
	}

	query := `INSERT INTO "user_info" (
		user_id,
		data
	) VALUES (
		$1,
		$2
	) ON CONFLICT (
		user_id
	) DO UPDATE SET data = $2, updated_at = NOW()`

	_, err = r.db.Exec(query,
		entity.UserId,
		data,
	)

	pKey = &pb.UserInfoPrimaryKey{
		UserId: entity.UserId,
	}

	return pKey, err
}

func (r *userInfoRepo) GetByPK(pKey *pb.UserInfoPrimaryKey) (res *pb.UserInfo, err error) {
	res = &pb.UserInfo{}
	query := `SELECT
		user_id,
		data
	FROM
		"user_info"
	WHERE
		user_id = $1`

	row, err := r.db.Query(query, pKey.UserId)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		var data []byte
		err = row.Scan(
			&res.UserId,
			&data,
		)
		if err != nil {
			return res, err
		}

		err = json.Unmarshal(data, &res.Data)
		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}
