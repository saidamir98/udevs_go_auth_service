package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type permissionRepo struct {
	db *sqlx.DB
}

func NewPermissionRepo(db *sqlx.DB) storage.PermissionRepoI {
	return &permissionRepo{
		db: db,
	}
}

func (r *permissionRepo) Create(ctx context.Context, entity *pb.CreatePermissionRequest) (pKey *pb.PermissionPrimaryKey, err error) {
	query := `INSERT INTO "permission" (
		id,
		client_platform_id,
		parent_id,
		name
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

	if entity.ParentId == uuid.String() {
		err = storage.ErrorTheSameId
		return
	}

	var nullStr *string
	if len(entity.ParentId) > 0 {
		nullStr = &entity.ParentId
	}

	_, err = r.db.ExecContext(ctx, query,
		uuid,
		entity.ClientPlatformId,
		nullStr,
		entity.Name,
	)

	pKey = &pb.PermissionPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *permissionRepo) GetByPK(ctx context.Context, pKey *pb.PermissionPrimaryKey) (res *pb.Permission, err error) {
	res = &pb.Permission{}
	query := `SELECT
		id,
		client_platform_id,
		parent_id,
		name
	FROM
		"permission"
	WHERE
		id = $1`

	row, err := r.db.QueryContext(ctx, query, pKey.Id)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		var nullableStr *string
		err = row.Scan(
			&res.Id,
			&res.ClientPlatformId,
			&nullableStr,
			&res.Name,
		)
		if nullableStr != nil {
			res.ParentId = *nullableStr
		}
		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *permissionRepo) GetList(ctx context.Context, queryParam *pb.GetPermissionListRequest) (res *pb.GetPermissionListResponse, err error) {
	res = &pb.GetPermissionListResponse{}
	params := make(map[string]interface{})
	query := `SELECT
		id,
		client_platform_id,
		parent_id,
		name
	FROM
		"permission"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

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

	cQ := `SELECT count(1) FROM "permission"` + filter
	row, err := r.db.NamedQueryContext(ctx, cQ, params)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&res.Count,
		)
		if err != nil {
			return res, err
		}
	}

	q := query + filter + order + arrangement + offset + limit
	rows, err := r.db.NamedQueryContext(ctx, q, params)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.Permission{}
		var nullableStr *string
		err = rows.Scan(
			&obj.Id,
			&obj.ClientPlatformId,
			&nullableStr,
			&obj.Name,
		)
		if nullableStr != nil {
			obj.ParentId = *nullableStr
		}

		if err != nil {
			return res, err
		}
		res.Permissions = append(res.Permissions, obj)
	}

	return res, nil
}

func (r *permissionRepo) Update(ctx context.Context, entity *pb.UpdatePermissionRequest) (rowsAffected int64, err error) {
	if entity.Id == entity.ParentId {
		err = storage.ErrorTheSameId
		return
	}

	query := `UPDATE "permission" SET
		client_platform_id = :client_platform_id,
		parent_id = :parent_id,
		name = :name,
		updated_at = now()
	WHERE
		id = :id`

	var nullStr *string
	if len(entity.ParentId) > 0 {
		nullStr = &entity.ParentId
	}

	params := map[string]interface{}{
		"id":                 entity.Id,
		"client_platform_id": entity.ClientPlatformId,
		"parent_id":          nullStr,
		"name":               entity.Name,
	}

	result, err := r.db.NamedExecContext(ctx, query, params)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *permissionRepo) Delete(ctx context.Context, pKey *pb.PermissionPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "permission" WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}
