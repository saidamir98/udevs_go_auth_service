package postgres

import (
	"context"
	"database/sql"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/storage"

	"github.com/saidamir98/udevs_pkg/util"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type clientPlatformRepo struct {
	db *pgxpool.Pool
}

func NewClientPlatformRepo(db *pgxpool.Pool) storage.ClientPlatformRepoI {
	return &clientPlatformRepo{
		db: db,
	}
}

func (r *clientPlatformRepo) Create(ctx context.Context, entity *pb.CreateClientPlatformRequest) (pKey *pb.ClientPlatformPrimaryKey, err error) {
	query := `INSERT INTO "client_platform" (
		id,
		project_id,
		name,
		subdomain
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return pKey, err
	}

	_, err = r.db.Exec(ctx, query,
		uuid.String(),
		entity.ProjectId,
		entity.Name,
		entity.Subdomain,
	)

	pKey = &pb.ClientPlatformPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *clientPlatformRepo) GetByPK(ctx context.Context, pKey *pb.ClientPlatformPrimaryKey) (res *pb.ClientPlatform, err error) {
	res = &pb.ClientPlatform{}
	query := `SELECT
		id,
		project_id,
		name,
		subdomain
	FROM
		"client_platform"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ProjectId,
		&res.Name,
		&res.Subdomain,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *clientPlatformRepo) GetByPKDetailed(ctx context.Context, pKey *pb.ClientPlatformPrimaryKey) (res *pb.ClientPlatformDetailedResponse, err error) {
	res = &pb.ClientPlatformDetailedResponse{}
	query := `SELECT
	  id,
	  project_id,
	  name,
	  subdomain
	FROM
	  "client_platform"
	WHERE
	  id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ProjectId,
		&res.Name,
		&res.Subdomain,
	)

	if err != nil {
		return res, err
	}

	query = `SELECT
	  id,
	  client_platform_id,
	  parent_id,
	  name
	FROM
	  "permission"
	WHERE
	  client_platform_id = $1`

	permissionRows, err := r.db.Query(ctx, query, pKey.Id)
	if err != nil {
		return res, nil
	}
	defer permissionRows.Close()
	for permissionRows.Next() {
		var (
			tempPersmission pb.Permission
			name            sql.NullString
			parentId        sql.NullString
		)
		err = permissionRows.Scan(
			&tempPersmission.Id,
			&tempPersmission.ClientPlatformId,
			&parentId,
			&name,
		)
		if err != nil {
			return res, err
		}
		tempPersmission.ParentId = parentId.String
		tempPersmission.Name = name.String

		res.Permissions = append(res.Permissions, &tempPersmission)
	}

	query = `SELECT
	  client_platform_id,
	  COALESCE(path, ''),
	  COALESCE(method, ''),
	  COALESCE(requests, 0) AS requests
	FROM
	  "scope"
	WHERE client_platform_id = $1`

	scopeRows, err := r.db.Query(ctx, query, pKey.Id)
	if err != nil {
		return res, nil
	}
	defer scopeRows.Close()

	for scopeRows.Next() {
		var tempScope pb.Scope
		err = scopeRows.Scan(
			&tempScope.ClientPlatformId,
			&tempScope.Path,
			&tempScope.Method,
			&tempScope.Requests,
		)
		if err != nil {
			return res, err
		}
		res.Scopes = append(res.Scopes, &tempScope)
	}

	return res, nil
}

func (r *clientPlatformRepo) GetList(ctx context.Context, queryParam *pb.GetClientPlatformListRequest) (res *pb.GetClientPlatformListResponse, err error) {
	res = &pb.GetClientPlatformListResponse{}
	var arr []interface{}
	params := make(map[string]interface{})
	query := `SELECT
		id,
		project_id,
		name,
		subdomain
	FROM
		"client_platform"`
	filter := " WHERE true"
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
		filter += " AND ((name || subdomain) ILIKE ('%' || :search || '%'))"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "client_platform"` + filter
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
		obj := &pb.ClientPlatform{}
		err = rows.Scan(
			&obj.Id,
			&obj.ProjectId,
			&obj.Name,
			&obj.Subdomain,
		)
		if err != nil {
			return res, err
		}
		res.ClientPlatforms = append(res.ClientPlatforms, obj)
	}

	return res, nil
}

func (r *clientPlatformRepo) Update(ctx context.Context, entity *pb.UpdateClientPlatformRequest) (rowsAffected int64, err error) {
	query := `UPDATE "client_platform" SET
		name = :name,
		subdomain = :subdomain,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":        entity.Id,
		"name":      entity.Name,
		"subdomain": entity.Subdomain,
	}

	query, arr := helper.ReplaceQueryParams(query, params)
	result, err := r.db.Exec(ctx, query, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *clientPlatformRepo) Delete(ctx context.Context, pKey *pb.ClientPlatformPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "client_platform" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}
