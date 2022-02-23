package postgres

import (
	"context"
	"database/sql"
	"upm/udevs_go_auth_service/config"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/pkg/util"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.UserRepoI {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, entity *pb.CreateUserRequest) (pKey *pb.UserPrimaryKey, err error) {
	query := `INSERT INTO "user" (
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		name,
		photo_url,
		phone,
		email,
		login,
		password,
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
		$10,
		$11,
		$12,
		$13
	)`

	uuid, err := uuid.NewRandom()
	if err != nil {
		return pKey, err
	}

	_, err = r.db.Exec(ctx, query,
		uuid,
		entity.ProjectId,
		entity.ClientPlatformId,
		entity.ClientTypeId,
		entity.RoleId,
		entity.Name,
		entity.PhotoUrl,
		entity.Phone,
		entity.Email,
		entity.Login,
		entity.Password,
		entity.Active,
		entity.ExpiresAt,
	)

	pKey = &pb.UserPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *userRepo) GetByPK(ctx context.Context, pKey *pb.UserPrimaryKey) (res *pb.User, err error) {
	res = &pb.User{}
	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		name,
		photo_url,
		phone,
		email,
		login,
		password,
		active,
		TO_CHAR(expires_at, ` + config.DatabaseQueryTimeLayout + `) AS expires_at,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"user"
	WHERE
		id = $1`

	err = r.db.QueryRow(ctx, query, pKey.Id).Scan(
		&res.Id,
		&res.ProjectId,
		&res.ClientPlatformId,
		&res.ClientTypeId,
		&res.RoleId,
		&res.Name,
		&res.PhotoUrl,
		&res.Phone,
		&res.Email,
		&res.Login,
		&res.Password,
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

func (r *userRepo) GetListByPKs(ctx context.Context, pKeys *pb.UserPrimaryKeyList) (res *pb.GetUserListResponse, err error) {
	res = &pb.GetUserListResponse{}
	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		name,
		photo_url,
		phone,
		email,
		login,
		password,
		active,
		expires_at,
		created_at,
		updated_at
	FROM
		"user"
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

		user := &pb.User{}
		err = rows.Scan(
			&user.Id,
			&user.ProjectId,
			&user.ClientPlatformId,
			&user.ClientTypeId,
			&user.RoleId,
			&user.Name,
			&user.PhotoUrl,
			&user.Phone,
			&user.Email,
			&user.Login,
			&user.Password,
			&active,
			&expiresAt,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return res, err
		}

		if active.Valid {
			user.Active = active.Int32
		}

		if expiresAt.Valid {
			user.ExpiresAt = expiresAt.String
		}

		if createdAt.Valid {
			user.CreatedAt = createdAt.String
		}

		if updatedAt.Valid {
			user.UpdatedAt = updatedAt.String
		}

		res.Users = append(res.Users, user)
	}

	return res, nil
}

func (r *userRepo) GetList(ctx context.Context, queryParam *pb.GetUserListRequest) (res *pb.GetUserListResponse, err error) {
	res = &pb.GetUserListResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		name,
		photo_url,
		phone,
		email,
		login,
		password,
		active,
		expires_at,
		created_at,
		updated_at
	FROM
		"user"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND ((name || phone || email || login) ILIKE ('%' || :search || '%'))"
	}

	if len(queryParam.ClientPlatformId) > 0 {
		params["client_platform_id"] = queryParam.ClientPlatformId
		filter += " AND client_platform_id = :client_platform_id"
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

	cQ := `SELECT count(1) FROM "user"` + filter
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
		obj := &pb.User{}
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
			&obj.Name,
			&obj.PhotoUrl,
			&obj.Phone,
			&obj.Email,
			&obj.Login,
			&obj.Password,
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

		res.Users = append(res.Users, obj)
	}

	return res, nil
}

func (r *userRepo) Update(ctx context.Context, entity *pb.UpdateUserRequest) (rowsAffected int64, err error) {
	query := `UPDATE "user" SET
		project_id = :project_id,
		client_platform_id = :client_platform_id,
		client_type_id = :client_type_id,
		role_id = :role_id,
		name = :name,
		photo_url = :photo_url,
		phone = :phone,
		email = :email,
		login = :login,
		active = :active,
		expires_at = :expires_at,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":                 entity.Id,
		"project_id":         entity.ProjectId,
		"client_platform_id": entity.ClientPlatformId,
		"client_type_id":     entity.ClientTypeId,
		"role_id":            entity.RoleId,
		"name":               entity.Name,
		"photo_url":          entity.PhotoUrl,
		"phone":              entity.Phone,
		"email":              entity.Email,
		"login":              entity.Login,
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

func (r *userRepo) Delete(ctx context.Context, pKey *pb.UserPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "user" WHERE id = $1`

	result, err := r.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (res *pb.User, err error) {
	res = &pb.User{}

	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		role_id,
		name,
		photo_url,
		phone,
		email,
		login,
		password,
		active,
		TO_CHAR(expires_at, ` + config.DatabaseQueryTimeLayout + `) AS expires_at,
		TO_CHAR(created_at, ` + config.DatabaseQueryTimeLayout + `) AS created_at,
		TO_CHAR(updated_at, ` + config.DatabaseQueryTimeLayout + `) AS updated_at
	FROM
		"user"
	WHERE`

	if util.IsValidEmail(username) {
		query = query + ` email = $1`
	} else if util.IsValidPhone(username) {
		query = query + ` phone = $1`
	} else {
		query = query + ` login = $1`
	}

	err = r.db.QueryRow(ctx, query, username).Scan(
		&res.Id,
		&res.ProjectId,
		&res.ClientPlatformId,
		&res.ClientTypeId,
		&res.RoleId,
		&res.Name,
		&res.PhotoUrl,
		&res.Phone,
		&res.Email,
		&res.Login,
		&res.Password,
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
