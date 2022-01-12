package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type clientPlatformRepo struct {
	db *sqlx.DB
}

func NewClientPlatformRepo(db *sqlx.DB) storage.ClientPlatformRepoI {
	return &clientPlatformRepo{
		db: db,
	}
}

func (r *clientPlatformRepo) Create(entity *pb.CreateClientPlatformRequest) (pKey *pb.ClientPlatformPrimaryKey, err error) {
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

	_, err = r.db.Exec(query,
		uuid,
		entity.ProjectId,
		entity.Name,
		entity.Subdomain,
	)

	pKey = &pb.ClientPlatformPrimaryKey{
		Id: uuid.String(),
	}

	return pKey, err
}

func (r *clientPlatformRepo) GetByPK(pKey *pb.ClientPlatformPrimaryKey) (res *pb.ClientPlatform, err error) {
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

	row, err := r.db.Query(query, pKey.Id)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		err = row.Scan(
			&res.Id,
			&res.ProjectId,
			&res.Name,
			&res.Subdomain,
		)
		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *clientPlatformRepo) GetList(queryParam *pb.GetClientPlatformListRequest) (res *pb.GetClientPlatformListResponse, err error) {
	res = &pb.GetClientPlatformListResponse{}
	params := make(map[string]interface{})
	query := `SELECT
		id,
		project_id,
		name,
		subdomain
	FROM
		"client_platform"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(queryParam.Search) > 0 {
		params["search"] = queryParam.Search
		filter += " AND (name ILIKE '%' || :search || '%')"
		filter += " AND (subdomain ILIKE '%' || :search || '%')"
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

func (r *clientPlatformRepo) Update(entity *pb.UpdateClientPlatformRequest) (rowsAffected int64, err error) {
	query := `UPDATE "client_platform" SET
		project_id = :project_id,
		name = :name,
		subdomain = :subdomain,
		updated_at = now()
	WHERE
		id = :id`

	params := map[string]interface{}{
		"id":         entity.Id,
		"project_id": entity.ProjectId,
		"name":       entity.Name,
		"subdomain":  entity.Subdomain,
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

func (r *clientPlatformRepo) Delete(pKey *pb.ClientPlatformPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "client_platform" WHERE id = $1`

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
