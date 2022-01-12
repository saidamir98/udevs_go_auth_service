package postgres

import (
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/storage"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type clientTypeRepo struct {
	db *sqlx.DB
}

func NewClientTypeRepo(db *sqlx.DB) storage.ClientTypeRepoI {
	return &clientTypeRepo{
		db: db,
	}
}

func (r *clientTypeRepo) Create(entity *pb.CreateClientTypeRequest) (pKey *pb.ClientTypePrimaryKey, err error) {
	query := `INSERT INTO "client_type" (
		id,
		name,
		confirm_by,
		self_register,
		self_recover
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

	_, err = r.db.Exec(query,
		uuid,
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

func (r *clientTypeRepo) GetByPK(pKey *pb.ClientTypePrimaryKey) (res *pb.ClientType, err error) {
	res = &pb.ClientType{}
	query := `SELECT
		id,
		name,
		confirm_by,
		self_register,
		self_recover
	FROM
		"client_type"
	WHERE
		id = $1`

	row, err := r.db.Query(query, pKey.Id)
	if err != nil {
		return res, err
	}
	defer row.Close()

	if row.Next() {
		var confirmBy string

		err = row.Scan(
			&res.Id,
			&res.Name,
			&confirmBy,
			&res.SelfRegister,
			&res.SelfRecover,
		)

		res.ConfirmBy = pb.ConfirmStrategies(pb.ConfirmStrategies_value[confirmBy])

		if err != nil {
			return res, err
		}
	} else {
		return res, storage.ErrorNotFound
	}

	return res, nil
}

func (r *clientTypeRepo) GetList(queryParam *pb.GetClientTypeListRequest) (res *pb.GetClientTypeListResponse, err error) {
	res = &pb.GetClientTypeListResponse{}
	params := make(map[string]interface{})
	query := `SELECT
		id,
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
		obj := &pb.ClientType{}
		var confirmBy string
		err = rows.Scan(
			&obj.Id,
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

func (r *clientTypeRepo) Update(entity *pb.UpdateClientTypeRequest) (rowsAffected int64, err error) {
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

func (r *clientTypeRepo) Delete(pKey *pb.ClientTypePrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "client_type" WHERE id = $1`

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
