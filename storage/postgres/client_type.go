package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/pkg/util"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type clientTypeRepo struct {
	db *pgxpool.Pool
}

func NewClientTypeRepo(db *pgxpool.Pool) storage.ClientTypeRepoI {
	return &clientTypeRepo{
		db: db,
	}
}

func (r *clientTypeRepo) Create(ctx context.Context, entity *pb.CreateClientTypeRequest) (pKey *pb.ClientTypePrimaryKey, err error) {
	query := `INSERT INTO "client_type" (
		id,
		project_id,
		name,
		confirm_by,
		self_register,
		self_recover
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6
	)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return pKey, err
	}

	_, err = r.db.Exec(ctx, query,
		uuid.String(),
		entity.ProjectId,
		entity.Name,
		entity.ConfirmBy.String(),
		entity.SelfRegister,
		entity.SelfRecover,
	)

	pKey = &pb.ClientTypePrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *clientTypeRepo) GetByPK(ctx context.Context, pKey *pb.ClientTypePrimaryKey) (res *pb.ClientType, err error) {
	res = &pb.ClientType{}
	var confirmBy string
	query := `SELECT
		id,
		project_id,
		name,
		confirm_by,
		self_register,
		self_recover
	FROM
		"client_type"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ProjectId,
		&res.Name,
		&confirmBy,
		&res.SelfRegister,
		&res.SelfRecover,
	)

	res.ConfirmBy = pb.ConfirmStrategies(pb.ConfirmStrategies_value[confirmBy])

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *clientTypeRepo) GetList(ctx context.Context, queryParam *pb.GetClientTypeListRequest) (res *pb.GetClientTypeListResponse, err error) {
	res = &pb.GetClientTypeListResponse{}
	params := make(map[string]interface{})
	query := `SELECT
		id,
		project_id,
		name,
		confirm_by,
		self_register,
		self_recover
	FROM
		"client_type"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if util.IsValidUUID(queryParam.ProjectId) {
		params["project_id"] = queryParam.ProjectId
		filter += " AND (project_id = :project_id)"
	}

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (name ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "client_type"` + filter
	var arr []interface{}

	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err = r.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)

	if err != nil {
		return res, err
	}

	q := query + filter + order + arrangement + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := r.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.ClientType{}
		var confirmBy string
		err = rows.Scan(
			&obj.Id,
			&obj.ProjectId,
			&obj.Name,
			&confirmBy,
			&obj.SelfRegister,
			&obj.SelfRecover,
		)
		if err != nil {
			return res, err
		}
		obj.ConfirmBy = pb.ConfirmStrategies(pb.ConfirmStrategies_value[confirmBy])
		res.ClientTypes = append(res.ClientTypes, obj)
	}

	return res, nil
}

func (r *clientTypeRepo) Update(ctx context.Context, entity *pb.UpdateClientTypeRequest) (rowsAffected int64, err error) {
	query := `UPDATE "client_type" SET
		name = :name,
		confirm_by = :confirm_by,
		self_register = :self_register,
		self_recover = :self_recover,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":            entity.Id,
		"name":          entity.Name,
		"confirm_by":    entity.ConfirmBy.String(),
		"self_register": entity.SelfRegister,
		"self_recover":  entity.SelfRecover,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := r.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *clientTypeRepo) Delete(ctx context.Context, pKey *pb.ClientTypePrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "client_type" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *clientTypeRepo) GetCompleteByPK(ctx context.Context, pKey *pb.ClientTypePrimaryKey) (res *pb.CompleteClientType, err error) {
	res = &pb.CompleteClientType{
		ClientType:     &pb.ClientType{},
		Relations:      make([]*pb.Relation, 0),
		UserInfoFields: make([]*pb.UserInfoField, 0),
		Roles:          make([]*pb.Role, 0),
	}

	query := `SELECT
		id,
		project_id,
		name,
		confirm_by,
		self_register,
		self_recover
	FROM
		"client_type"
	WHERE
		id = $1`

	var confirmBy string

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.ClientType.Id,
		&res.ClientType.ProjectId,
		&res.ClientType.Name,
		&confirmBy,
		&res.ClientType.SelfRegister,
		&res.ClientType.SelfRecover,
	)

	res.ClientType.ConfirmBy = pb.ConfirmStrategies(pb.ConfirmStrategies_value[confirmBy])

	if err != nil {
		return res, err
	}

	query1 := `SELECT
		id,
		client_type_id,
		type,
		name,
		description
	FROM
		"relation"
	WHERE
		client_type_id = $1`

	rows1, err := r.db.Query(ctx, query1, res.ClientType.Id)
	if err != nil {
		return res, err
	}
	defer rows1.Close()

	for rows1.Next() {
		obj := &pb.Relation{}
		var relationType string

		err = rows1.Scan(
			&obj.Id,
			&obj.ClientTypeId,
			&relationType,
			&obj.Name,
			&obj.Description,
		)

		if err != nil {
			return res, err
		}
		obj.Type = pb.RelationTypes(pb.RelationTypes_value[relationType])

		res.Relations = append(res.Relations, obj)
	}

	query2 := `SELECT
		id,
		client_type_id,
		field_name,
		field_type,
		data_type
	FROM
		"user_info_field"
	WHERE
		client_type_id = $1`

	rows2, err := r.db.Query(ctx, query2, res.ClientType.Id)
	if err != nil {
		return res, err
	}
	defer rows2.Close()

	for rows2.Next() {
		obj := &pb.UserInfoField{}

		err = rows2.Scan(
			&obj.Id,
			&obj.ClientTypeId,
			&obj.FieldName,
			&obj.FieldType,
			&obj.DataType,
		)

		if err != nil {
			return res, err
		}

		res.UserInfoFields = append(res.UserInfoFields, obj)
	}

	query3 := `SELECT
		id,
		client_type_id,
		name
		FROM
			"role"
		WHERE
			client_type_id = $1`

	rows3, err := r.db.Query(ctx, query3, res.ClientType.Id)
	if err != nil {
		return res, err
	}
	defer rows3.Close()

	for rows3.Next() {
		obj := &pb.Role{}
		err = rows3.Scan(
			&obj.Id,
			&obj.ClientTypeId,
			&obj.Name,
		)
		if err != nil {
			return res, err
		}
		res.Roles = append(res.Roles, obj)
	}

	return res, nil
}
