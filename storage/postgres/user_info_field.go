package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userInfoFieldRepo struct {
	db *pgxpool.Pool
}

func NewUserInfoFieldRepo(db *pgxpool.Pool) storage.UserInfoFieldRepoI {
	return &userInfoFieldRepo{
		db: db,
	}
}

func (r *userInfoFieldRepo) Add(ctx context.Context, entity *pb.AddUserInfoFieldRequest) (pKey *pb.UserInfoFieldPrimaryKey, err error) {
	query := `INSERT INTO "user_info_field" (
		id,
		client_type_id,
		field_name,
		field_type,
		data_type
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
		entity.FieldName,
		entity.FieldType,
		entity.DataType,
	)

	pKey = &pb.UserInfoFieldPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *userInfoFieldRepo) GetByPK(ctx context.Context, pKey *pb.UserInfoFieldPrimaryKey) (res *pb.UserInfoField, err error) {
	res = &pb.UserInfoField{}
	query := `SELECT
		id,
		client_type_id,
		field_name,
		field_type,
		data_type
	FROM
		"user_info_field"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ClientTypeId,
		&res.FieldName,
		&res.FieldType,
		&res.DataType,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *userInfoFieldRepo) Update(ctx context.Context, entity *pb.UpdateUserInfoFieldRequest) (rowsAffected int64, err error) {
	query := `UPDATE "user_info_field" SET
		client_type_id = :client_type_id,
		field_name = :field_name,
		field_type = :field_type,
		data_type = :data_type,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":             entity.Id,
		"client_type_id": entity.ClientTypeId,
		"field_name":     entity.FieldName,
		"field_type":     entity.FieldType,
		"data_type":      entity.DataType,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := r.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *userInfoFieldRepo) Remove(ctx context.Context, pKey *pb.UserInfoFieldPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "user_info_field" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
