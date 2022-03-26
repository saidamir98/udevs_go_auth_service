package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"upm/udevs_go_auth_service/config"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type IntegrationRepo struct {
	db *pgxpool.Pool
}

func NewIntegrationRepo(db *pgxpool.Pool) storage.IntegrationRepoI {
	return &IntegrationRepo{
		db: db,
	}
}

func (r *IntegrationRepo) Create(ctx context.Context, entity *pb.CreateIntegrationRequest) (pKey *pb.IntegrationPrimaryKey, err error) {
	query := `INSERT INTO "integration" (
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		title,
		secret_key,
		ip_whitelist,
		active,
		expires_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9,
		$10
	)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return pKey, err
	}

	jsonStruct, err := json.Marshal(entity.IpWhitelist)
	if err != nil {
		return nil, err
	}

	_, err = r.db.Exec(ctx, query,
		uuid.String(),
		entity.ProjectId,
		entity.ClientPlatformId,
		entity.ClientTypeId,
		entity.RoleId,
		entity.Title,
		entity.SecretKey,
		jsonStruct,
		entity.Active,
		entity.ExpiresAt,
	)

	pKey = &pb.IntegrationPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *IntegrationRepo) GetByPK(ctx context.Context, pKey *pb.IntegrationPrimaryKey) (res *pb.Integration, err error) {
	res = &pb.Integration{}
	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		title,
		secret_key,
		ip_whitelist,
		active,
		TO_CHAR(expires_at, ` + config.DatabaseQueryTimeLayout + `) AS expires_at,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"integration"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ProjectId,
		&res.ClientPlatformId,
		&res.ClientTypeId,
		&res.RoleId,
		&res.Title,
		&res.SecretKey,
		&res.IpWhitelist,
		&res.Active,
		&res.ExpiresAt,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *IntegrationRepo) GetIntegrationSessions(ctx context.Context, pKey *pb.IntegrationPrimaryKey) (res *pb.GetIntegrationSessionsResponse, err error) {
	res = &pb.GetIntegrationSessionsResponse{}
	query := `
		SELECT 
			id,
			integration_id,
			project_id,
			client_type_id,
			client_platform_id,
			role_id,
			ip,
			data,
			TO_CHAR(expires_at, ` + config.DatabaseQueryTimeLayout + `) AS expires_at,
			TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
			TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
		FROM
			"session"
		WHERE integration_id = $1
	`

	rows, err := r.db.Query(ctx, query, pKey.GetId())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		temp := new(pb.Session)
		err = rows.Scan(
			&temp.Id,
			&temp.IntegrationId,
			&temp.ProjectId,
			&temp.ClientTypeId,
			&temp.ClientPlatformId,
			&temp.RoleId,
			&temp.Ip,
			&temp.Data,
			&temp.ExpiresAt,
			&temp.CreatedAt,
			&temp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		res.Sessions = append(res.Sessions, temp)
	}

	return res, nil
}

func (r *IntegrationRepo) GetListByPKs(ctx context.Context, pKeys *pb.IntegrationPrimaryKeyList) (res *pb.GetIntegrationListResponse, err error) {
	res = &pb.GetIntegrationListResponse{}
	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		title,
		secret_key,
		ip_whitelist,
		active,
		expires_at,
		created_at,
		updated_at
	FROM
		"integration"
	WHERE
		id = ANY($1)`

	rows, err := r.db.Query(ctx, query, pq.Array(pKeys.Ids))
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			active    sql.NullInt32
			expiresAt sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		integration := &pb.Integration{}
		err = rows.Scan(
			&integration.Id,
			&integration.ProjectId,
			&integration.ClientPlatformId,
			&integration.ClientTypeId,
			&integration.RoleId,
			&integration.Title,
			&integration.SecretKey,
			&integration.IpWhitelist,
			&active,
			&expiresAt,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return res, err
		}

		if active.Valid {
			integration.Active = active.Int32
		}

		if expiresAt.Valid {
			integration.ExpiresAt = expiresAt.String
		}

		if createdAt.Valid {
			integration.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			integration.UpdatedAt = updatedAt.String
		}

		res.Integrations = append(res.Integrations, integration)
	}

	return res, nil
}

func (r *IntegrationRepo) GetList(ctx context.Context, queryParam *pb.GetIntegrationListRequest) (res *pb.GetIntegrationListResponse, err error) {
	res = &pb.GetIntegrationListResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		title,
		secret_key,
		ip_whitelist::varchar,
		active,
		expires_at,
		created_at,
		updated_at
	FROM
		"integration"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND ((title) ILIKE ('%' || :search || '%'))"
	}

	if len(queryParam.ClientPlatformId) > 0 {
		params["client_platform_id"] = queryParam.ClientPlatformId
		filter += " AND client_platform_id = :client_platform_id"
	}
	if len(queryParam.ProjectId) > 0 {
		params["project_id"] = queryParam.ProjectId
		filter += " AND project_id = :project_id"
	}

	if len(queryParam.ClientTypeId) > 0 {
		params["client_type_id"] = queryParam.ClientTypeId
		filter += " AND client_type_id = :client_type_id"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "integration"` + filter
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
		obj := &pb.Integration{}
		var (
			active    sql.NullInt32
			expiresAt sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err = rows.Scan(
			&obj.Id,
			&obj.ProjectId,
			&obj.ClientPlatformId,
			&obj.ClientTypeId,
			&obj.RoleId,
			&obj.Title,
			&obj.SecretKey,
			&obj.IpWhitelist,
			&active,
			&expiresAt,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return res, err
		}

		if active.Valid {
			obj.Active = active.Int32
		}

		if expiresAt.Valid {
			obj.ExpiresAt = expiresAt.String
		}

		if createdAt.Valid {
			obj.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			obj.UpdatedAt = updatedAt.String
		}

		res.Integrations = append(res.Integrations, obj)
	}

	return res, nil
}

func (r *IntegrationRepo) Update(ctx context.Context, entity *pb.UpdateIntegrationRequest) (rowsAffected int64, err error) {
	query := `UPDATE "integration" SET
		project_id = :project_id,
		client_platform_id = :client_platform_id,
		client_type_id = :client_type_id,
		role_id = :role_id,
		title = :title,
		ip_whitelist = :ip_whitelist,
		active = :active,
		expires_at = :expires_at,
		updated_at = now()
	WHERE
		id = :id`

	jsonStruct, err := json.Marshal(entity.IpWhitelist)
	if err != nil {
		return 0, err
	}
	params := map[string]interface{}{
		"id":                 entity.Id,
		"project_id":         entity.ProjectId,
		"client_platform_id": entity.ClientPlatformId,
		"client_type_id":     entity.ClientTypeId,
		"role_id":            entity.RoleId,
		"title":              entity.Name,
		"ip_whitelist":       jsonStruct,
		"active":             entity.Active,
		"expires_at":         entity.ExpiresAt,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := r.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *IntegrationRepo) Delete(ctx context.Context, pKey *pb.IntegrationPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "integration" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
