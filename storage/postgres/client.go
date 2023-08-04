package postgres

import (
	"context"
	pb "upm/udevs_go_auth_service/genproto/auth_service"
	"upm/udevs_go_auth_service/pkg/helper"
	"upm/udevs_go_auth_service/storage"

	"github.com/saidamir98/udevs_pkg/util"

	"github.com/jackc/pgx/v4/pgxpool"
)

type clientRepo struct {
	db *pgxpool.Pool
}

func NewClientRepo(db *pgxpool.Pool) storage.ClientRepoI {
	return &clientRepo{
		db: db,
	}
}

func (r *clientRepo) Add(ctx context.Context, projectID string, entity *pb.AddClientRequest) (err error) {
	query := `INSERT INTO "client" (
		project_id,
		client_platform_id,
		client_type_id,
		login_strategy
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`

	_, err = r.db.Exec(ctx, query,
		projectID,
		entity.ClientPlatformId,
		entity.ClientTypeId,
		entity.LoginStrategy.String(),
	)

	return err
}

func (r *clientRepo) GetByPK(ctx context.Context, pKey *pb.ClientPrimaryKey) (res *pb.Client, err error) {
	res = &pb.Client{}
	var loginStrategy string
	query := `SELECT
		project_id,
		client_platform_id,
		client_type_id,
		login_strategy
	FROM
		"client"
	WHERE
		client_platform_id = $1 AND client_type_id = $2`

	err = r.db.QueryRow(ctx, query, pKey.ClientPlatformId, pKey.ClientTypeId).Scan(
		&res.ProjectId,
		&res.ClientPlatformId,
		&res.ClientTypeId,
		&loginStrategy,
	)

	res.LoginStrategy = pb.LoginStrategies(pb.LoginStrategies_value[loginStrategy])

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *clientRepo) Update(ctx context.Context, entity *pb.UpdateClientRequest) (rowsAffected int64, err error) {
	query := `UPDATE "client" SET
		login_strategy = :login_strategy,
		updated_at = now()
	WHERE
		client_platform_id = :client_platform_id AND client_type_id = :client_type_id`

	params := map[string]interface{}{
		"client_platform_id": entity.ClientPlatformId,
		"client_type_id":     entity.ClientTypeId,
		"login_strategy":     entity.LoginStrategy.String(),
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := r.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *clientRepo) Remove(ctx context.Context, pKey *pb.ClientPrimaryKey) (rowsAffected int64, err error) {
	query := `DELETE FROM "client" WHERE client_platform_id = $1 AND client_type_id = $2`

	result, err := r.db.Exec(ctx, query, pKey.ClientPlatformId, pKey.ClientTypeId)
	if err != nil {
		return 0, err
	}

	rowsAffected = result.RowsAffected()

	return rowsAffected, err
}

func (r *clientRepo) GetList(ctx context.Context, queryParam *pb.GetClientListRequest) (res *pb.GetClientListResponse, err error) {
	res = &pb.GetClientListResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		project_id,
		client_platform_id,
		client_type_id,
		login_strategy
	FROM
		"client"`
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
		filter += " AND (login_strategy ILIKE '%' || :search || '%')"
	}

	if queryParam.Offset > 0 {
		params["offset"] = queryParam.Offset
		offset = " OFFSET :offset"
	}

	if queryParam.Limit > 0 {
		params["limit"] = queryParam.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "client"` + filter
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
		obj := &pb.Client{}
		var loginStrategy string
		err = rows.Scan(
			&obj.ProjectId,
			&obj.ClientPlatformId,
			&obj.ClientTypeId,
			&loginStrategy,
		)
		if err != nil {
			return res, err
		}
		obj.LoginStrategy = pb.LoginStrategies(pb.LoginStrategies_value[loginStrategy])
		res.Clients = append(res.Clients, obj)
	}

	return res, nil
}

func (r *clientRepo) GetMatrix(ctx context.Context, req *pb.GetClientMatrixRequest) (res *pb.GetClientMatrixResponse, err error) {
	if !util.IsValidUUID(req.ProjectId) {
		return res, storage.ErrorProjectId
	}
	res = &pb.GetClientMatrixResponse{}

	queryClientPlatform := `SELECT
		id,
		project_id,
		name,
		subdomain
	FROM
		"client_platform"
	WHERE
		project_id = $1`

	clientPlatformRows, err := r.db.Query(ctx, queryClientPlatform, req.ProjectId)
	if err != nil {
		return res, err
	}
	defer clientPlatformRows.Close()

	for clientPlatformRows.Next() {
		obj := &pb.ClientPlatform{}
		err = clientPlatformRows.Scan(
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

	queryClientType := `SELECT
		id,
		project_id,
		name,
		confirm_by,
		self_register,
		self_recover
	FROM
		"client_type"
	WHERE
		project_id = $1`

	clientTypeRows, err := r.db.Query(ctx, queryClientType, req.ProjectId)
	if err != nil {
		return res, err
	}
	defer clientTypeRows.Close()

	for clientTypeRows.Next() {
		obj := &pb.ClientType{}
		var confirmBy string
		err = clientTypeRows.Scan(
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

	queryClient := `SELECT
		project_id,
		client_platform_id,
		client_type_id,
		login_strategy
	FROM
		"client"
	WHERE
		project_id = $1`

	clientRows, err := r.db.Query(ctx, queryClient, req.ProjectId)
	if err != nil {
		return res, err
	}
	defer clientRows.Close()

	for clientRows.Next() {
		obj := &pb.Client{}
		var loginStrategy string
		err = clientRows.Scan(
			&obj.ProjectId,
			&obj.ClientPlatformId,
			&obj.ClientTypeId,
			&loginStrategy,
		)
		if err != nil {
			return res, err
		}
		obj.LoginStrategy = pb.LoginStrategies(pb.LoginStrategies_value[loginStrategy])
		res.Clients = append(res.Clients, obj)
	}

	return
}
