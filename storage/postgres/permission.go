package postgres

import (
	"context"
	"database/sql"
	"strings"
	"upm/udevs_go_auth_service/genproto/auth_service"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type permissionRepo struct {
	db *pgxpool.Pool
}

func NewPermissionRepo(db *pgxpool.Pool) storage.PermissionRepoI {
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

	_, err = r.db.Exec(ctx, query,
		uuid.String(),
		entity.ClientPlatformId,
		nullStr,
		entity.Name,
	)

	pKey = &pb.PermissionPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *permissionRepo) GetByPK(ctx context.Context, pKey *pb.PermissionPrimaryKey) (res *pb.GetPermissionByIDResponse, err error) {
	res = &pb.GetPermissionByIDResponse{}
	var nullableStr *string
	query := `SELECT
		id,
		client_platform_id,
		parent_id,
		name
	FROM
		"permission"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ClientPlatformId,
		&nullableStr,
		&res.Name,
	)
	if err != nil {
		return res, err
	}
	if nullableStr != nil {
		res.ParentId = *nullableStr
	}

	queryPermissionScopes := `SELECT
		permission_id,
		client_platform_id,
		path,
		method
	FROM
		"permission_scope"
	WHERE
		permission_id = $1`

	permissionScopesRows, err := r.db.Query(ctx, queryPermissionScopes, res.Id)
	if err != nil {
		return res, err
	}
	defer permissionScopesRows.Close()

	for permissionScopesRows.Next() {
		obj := &pb.PermissionScope{}
		err = permissionScopesRows.Scan(
			&obj.PermissionId,
			&obj.ClientPlatformId,
			&obj.Path,
			&obj.Method,
		)
		if err != nil {
			return res, err
		}
		res.PermissionScopes = append(res.PermissionScopes, obj)
	}

	return res, nil
}

func (r *permissionRepo) GetList(ctx context.Context, queryParam *pb.GetPermissionListRequest) (res *pb.GetPermissionListResponse, err error) {
	res = &pb.GetPermissionListResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
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

func (r *permissionRepo) GetListByRoleId(ctx context.Context, roleID string) (res []*pb.Permission, err error) {
	var (
		permissionMap = make(map[string]*pb.Permission)
	)
	res = []*pb.Permission{}

	query := `SELECT
		p.id,
		p.client_platform_id,
		p.parent_id,
		p.name
	FROM 
		"role_permission"
	AS
		rp
	INNER JOIN
			"permission" 
		AS 
			p
	ON
		p.id = rp.permission_id
	WHERE
		rp.role_id = $1`

	rows, err := r.db.Query(ctx, query, roleID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		permission := &pb.Permission{}
		var parentID *string
		err = rows.Scan(
			&permission.Id,
			&permission.ClientPlatformId,
			&parentID,
			&permission.Name,
		)
		if parentID != nil {
			permission.ParentId = *parentID
		}

		if permission.ParentId != "" {
			permission.Name = permissionMap[permission.GetParentId()].GetName() + "/" + permission.Name
		}

		permissionMap[permission.Id] = permission

		if err != nil {
			return res, err
		}
		res = append(res, permission)
	}
	return res, nil
}

func (r *permissionRepo) GetListByClientPlatformId(ctx context.Context, clientPlatformID string) (res []*pb.Permission, err error) {
	var (
		permissionMap = make(map[string]*pb.Permission)
	)
	res = []*pb.Permission{}

	query := `SELECT
		id,
		client_platform_id,
		parent_id,
		name
	FROM
		"permission"
	WHERE
		client_platform_id = $1
	ORDER BY created_at`

	rows, err := r.db.Query(ctx, query, clientPlatformID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		permission := &pb.Permission{}
		var parentID *string
		err = rows.Scan(
			&permission.Id,
			&permission.ClientPlatformId,
			&parentID,
			&permission.Name,
		)
		if parentID != nil {
			permission.ParentId = *parentID
		}

		if permission.ParentId != "" {
			permission.Name = permissionMap[permission.GetParentId()].GetName() + "/" + permission.Name
		}

		permissionMap[permission.Id] = permission

		if err != nil {
			return res, err
		}
		res = append(res, permission)
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

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := r.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *permissionRepo) Delete(ctx context.Context, pKey *pb.PermissionPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "permission" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *permissionRepo) GeneratePermission(ctx context.Context, req auth_service.PermissionGenerated) (res *auth_service.Permission, err error) {

	return nil, nil
}

func MakePermissions(permissions *auth_service.PermissionGenerated_Permission) []*auth_service.Permission {
	var (
		res []*auth_service.Permission
	)
	if permissions.Children == nil {
		return res
	}

	return res
}

func InsertScopes(scopes []*auth_service.PermissionGenerated_Permission_Scope, permissionID, clientPlatformID string, tx *sql.Tx) error {
	values := []interface{}{}
	query := `
		INSERT INTO "permission_scope" (
		permission_id,
		client_platform_id,
		path,
		method
		) VALUES `

	for _, v := range scopes {
		query += "(?, ?, ?, ?),"
		values = append(values,
			permissionID,
			clientPlatformID,
			v.GetUrl(),
			v.GetMethod(),
		)
	}

	query = strings.TrimSuffix(query, ",")
	query = helper.ReplaceSQL(query, "?")

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}

	return nil
}
