package postgres

import (
	"context"
	"encoding/json"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userInfoRepo struct {
	db *pgxpool.Pool
}

func NewUserInfoRepo(db *pgxpool.Pool) storage.UserInfoRepoI {
	return &userInfoRepo{
		db: db,
	}
}

func (r *userInfoRepo) Upsert(ctx context.Context, entity *pb.UpsertUserInfoRequest) (pKey *pb.UserInfoPrimaryKey, err error) {
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

	_, err = r.db.Exec(ctx, query,
		entity.UserId,
		data,
	)

	pKey = &pb.UserInfoPrimaryKey{
		UserId: entity.UserId,
	}

	return pKey, err
}

func (r *userInfoRepo) GetByPK(ctx context.Context, pKey *pb.UserInfoPrimaryKey) (res *pb.UserInfo, err error) {
	res = &pb.UserInfo{}
	var data []byte

	query := `SELECT
		user_id,
		data
	FROM
		"user_info"
	WHERE
		user_id = $1`

	err = r.db.QueryRow(ctx, query, pKey.UserId).Scan(
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

	return res, nil
}
