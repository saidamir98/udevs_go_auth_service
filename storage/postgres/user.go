package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/util"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) storage.UserRepoI {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(entity *pb.CreateUserRequest) (pKey *pb.UserPrimaryKey, err error) {
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

	_, err = r.db.Exec(query,
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

func (r *userRepo) GetByPK(pKey *pb.UserPrimaryKey) (res *pb.User, err error) {
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
		expires_at,
		created_at,
		updated_at
	FROM
		"user"
	WHERE
		id = $1`

	row, err := r.db.Query(query, pKey.Id)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
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
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *userRepo) GetList(queryParam *pb.GetUserListRequest) (res *pb.GetUserListResponse, err error) {
	res = &pb.GetUserListResponse{}
	params := make(map[string]interface{})
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
	row, err := r.db.NamedQuery(cQ, params)
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
	rows, err := r.db.NamedQuery(q, params)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.User{}
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
			&obj.Active,
			&obj.ExpiresAt,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		res.Users = append(res.Users, obj)
	}

	return res, nil
}

func (r *userRepo) Update(entity *pb.UpdateUserRequest) (rowsAffected int64, err error) {
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

func (r *userRepo) Delete(pKey *pb.UserPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "user" WHERE id = $1`

	result, err := r.db.Exec(query, pKey.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *userRepo) GetByUsername(username string) (res *pb.User, err error) {
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
		expires_at,
		created_at,
		updated_at
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

	row, err := r.db.Query(query, username)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
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
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}
