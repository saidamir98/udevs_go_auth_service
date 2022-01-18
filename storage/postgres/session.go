package postgres

import (
	"time"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type sessionRepo struct {
	db *sqlx.DB
}

func NewSessionRepo(db *sqlx.DB) storage.SessionRepoI {
	return &sessionRepo{
		db: db,
	}
}

func (r *sessionRepo) Create(entity *pb.CreateSessionRequest) (pKey *pb.SessionPrimaryKey, err error) {
	query := `INSERT INTO "session" (
		id,
		project_id,
		client_platform_id,
		client_type_id,
		user_id,
		role_id,
		ip,
		data,
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
		$9
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
		entity.UserId,
		entity.RoleId,
		entity.Ip,
		entity.Data,
		entity.ExpiresAt,
	)

	pKey = &pb.SessionPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *sessionRepo) GetByPK(pKey *pb.SessionPrimaryKey) (res *pb.Session, err error) {
	res = &pb.Session{}
	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		user_id,
		role_id,
		ip,
		data,
		expires_at,
		created_at,
		updated_at
	FROM
		"session"
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
			&res.UserId,
			&res.RoleId,
			&res.Ip,
			&res.Data,
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

func (r *sessionRepo) GetList(queryParam *pb.GetSessionListRequest) (res *pb.GetSessionListResponse, err error) {
	res = &pb.GetSessionListResponse{}
	params := make(map[string]interface{})
	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		user_id,
		role_id,
		ip,
		data,
		expires_at,
		created_at,
		updated_at
	FROM
		"session"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND ((ip) ILIKE ('%' || :search || '%'))"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "session"` + filter
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
		obj := &pb.Session{}
		err = rows.Scan(
			&obj.Id,
			&obj.ProjectId,
			&obj.ClientPlatformId,
			&obj.ClientTypeId,
			&obj.UserId,
			&obj.RoleId,
			&obj.Ip,
			&obj.Data,
			&obj.ExpiresAt,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		res.Sessions = append(res.Sessions, obj)
	}

	return res, nil
}

func (r *sessionRepo) Update(entity *pb.UpdateSessionRequest) (rowsAffected int64, err error) {
	query := `UPDATE "session" SET
		project_id = :project_id,
		client_platform_id = :client_platform_id,
		client_type_id = :client_type_id,
		user_id = :user_id,
		role_id = :role_id,
		ip = :ip,
		data = :data,
		expires_at = :expires_at,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":                 entity.Id,
		"project_id":         entity.ProjectId,
		"client_platform_id": entity.ClientPlatformId,
		"client_type_id":     entity.ClientTypeId,
		"user_id":            entity.UserId,
		"role_id":            entity.RoleId,
		"ip":                 entity.Ip,
		"data":               entity.Data,
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

func (r *sessionRepo) Delete(pKey *pb.SessionPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "session" WHERE id = $1`

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

func (r *sessionRepo) DeleteExpiredUserSessions(userID string) (rowsAffected int64, err error) {
	query := `DELETE FROM "session" WHERE user_id = $1 AND expires_at < $2`

	result, err := r.db.Exec(query, userID, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return 0, err
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, err
}

func (r *sessionRepo) GetSessionListByUserID(userID string) (res *pb.GetSessionListResponse, err error) {
	res = &pb.GetSessionListResponse{}

	query := `SELECT
		id,
		project_id,
		client_platform_id,
		client_type_id,
		user_id,
		role_id,
		ip,
		data,
		expires_at,
		created_at,
		updated_at
	FROM
		"session"
	WHERE user_id = $1`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &pb.Session{}
		err = rows.Scan(
			&obj.Id,
			&obj.ProjectId,
			&obj.ClientPlatformId,
			&obj.ClientTypeId,
			&obj.UserId,
			&obj.RoleId,
			&obj.Ip,
			&obj.Data,
			&obj.ExpiresAt,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		res.Sessions = append(res.Sessions, obj)
	}

	return res, nil
}
