package handlers

import (
	"context"
	"upm/udevs_go_auth_service/api/http"

	"upm/udevs_go_auth_service/genproto/auth_service"

	"upm/udevs_go_auth_service/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateClientPlatform godoc
// @ID create_client_platform
// @Router /client-platform [POST]
// @Summary Create ClientPlatform
// @Description Create ClientPlatform
// @Tags ClientPlatform
// @Accept json
// @Produce json
// @Param client-platform body auth_service.CreateClientPlatformRequest true "CreateClientPlatformRequestBody"
// @Success 201 {object} http.Response{data=auth_service.ClientPlatform} "ClientPlatform data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateClientPlatform(c *gin.Context) {
	var client_platform auth_service.CreateClientPlatformRequest

	err := c.ShouldBindJSON(&client_platform)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().CreateClientPlatform(
		context.Background(),
		&client_platform,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetClientPlatformList godoc
// @ID get_client_platform_list
// @Router /client-platform [GET]
// @Summary Get ClientPlatform List
// @Description  Get ClientPlatform List
// @Tags ClientPlatform
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetClientPlatformListResponse} "GetClientPlatformListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetClientPlatformList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.ClientService().GetClientPlatformList(
		context.Background(),
		&auth_service.GetClientPlatformListRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetClientPlatformByID godoc
// @ID get_client_platform_by_id
// @Router /client-platform/{client-platform-id} [GET]
// @Summary Get ClientPlatform By ID
// @Description Get ClientPlatform By ID
// @Tags ClientPlatform
// @Accept json
// @Produce json
// @Param client-platform-id path string true "client-platform-id"
// @Success 200 {object} http.Response{data=auth_service.ClientPlatform} "ClientPlatformBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetClientPlatformByID(c *gin.Context) {
	client_platformID := c.Param("client-platform-id")

	if !util.IsValidUUID(client_platformID) {
		h.handleResponse(c, http.InvalidArgument, "client_platform id is an invalid uuid")
		return
	}

	resp, err := h.services.ClientService().GetClientPlatformByID(
		context.Background(),
		&auth_service.ClientPlatformPrimaryKey{
			Id: client_platformID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateClientPlatform godoc
// @ID update_client_platform
// @Router /client-platform [PUT]
// @Summary Update ClientPlatform
// @Description Update ClientPlatform
// @Tags ClientPlatform
// @Accept json
// @Produce json
// @Param client-platform body auth_service.UpdateClientPlatformRequest true "UpdateClientPlatformRequestBody"
// @Success 200 {object} http.Response{data=auth_service.ClientPlatform} "ClientPlatform data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateClientPlatform(c *gin.Context) {
	var client_platform auth_service.UpdateClientPlatformRequest

	err := c.ShouldBindJSON(&client_platform)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().UpdateClientPlatform(
		context.Background(),
		&client_platform,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteClientPlatform godoc
// @ID delete_client_platform
// @Router /client-platform/{client-platform-id} [DELETE]
// @Summary Delete ClientPlatform
// @Description Get ClientPlatform
// @Tags ClientPlatform
// @Accept json
// @Produce json
// @Param client-platform-id path string true "client-platform-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteClientPlatform(c *gin.Context) {
	client_platformID := c.Param("client-platform-id")

	if !util.IsValidUUID(client_platformID) {
		h.handleResponse(c, http.InvalidArgument, "client_platform id is an invalid uuid")
		return
	}

	resp, err := h.services.ClientService().DeleteClientPlatform(
		context.Background(),
		&auth_service.ClientPlatformPrimaryKey{
			Id: client_platformID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// CreateClientType godoc
// @ID create_client_type
// @Router /client-type [POST]
// @Summary Create ClientType
// @Description Create ClientType
// @Tags ClientType
// @Accept json
// @Produce json
// @Param client-type body auth_service.CreateClientTypeRequest true "CreateClientTypeRequestBody"
// @Success 201 {object} http.Response{data=auth_service.ClientType} "ClientType data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateClientType(c *gin.Context) {
	var client_type auth_service.CreateClientTypeRequest

	err := c.ShouldBindJSON(&client_type)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().CreateClientType(
		context.Background(),
		&client_type,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetClientTypeList godoc
// @ID get_client_type_list
// @Router /client-type [GET]
// @Summary Get ClientType List
// @Description  Get ClientType List
// @Tags ClientType
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetClientTypeListResponse} "GetClientTypeListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetClientTypeList(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.ClientService().GetClientTypeList(
		context.Background(),
		&auth_service.GetClientTypeListRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetClientTypeByID godoc
// @ID get_client_type_by_id
// @Router /client-type/{client-type-id} [GET]
// @Summary Get ClientType By ID
// @Description Get ClientType By ID
// @Tags ClientType
// @Accept json
// @Produce json
// @Param client-type-id path string true "client-type-id"
// @Success 200 {object} http.Response{data=auth_service.CompleteClientType} "ClientTypeBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetClientTypeByID(c *gin.Context) {
	client_typeID := c.Param("client-type-id")

	if !util.IsValidUUID(client_typeID) {
		h.handleResponse(c, http.InvalidArgument, "client_type id is an invalid uuid")
		return
	}

	resp, err := h.services.ClientService().GetClientTypeByID(
		context.Background(),
		&auth_service.ClientTypePrimaryKey{
			Id: client_typeID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateClientType godoc
// @ID update_client_type
// @Router /client-type [PUT]
// @Summary Update ClientType
// @Description Update ClientType
// @Tags ClientType
// @Accept json
// @Produce json
// @Param client-type body auth_service.UpdateClientTypeRequest true "UpdateClientTypeRequestBody"
// @Success 200 {object} http.Response{data=auth_service.ClientType} "ClientType data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateClientType(c *gin.Context) {
	var client_type auth_service.UpdateClientTypeRequest

	err := c.ShouldBindJSON(&client_type)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().UpdateClientType(
		context.Background(),
		&client_type,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteClientType godoc
// @ID delete_client_type
// @Router /client-type/{client-type-id} [DELETE]
// @Summary Delete ClientType
// @Description Get ClientType
// @Tags ClientType
// @Accept json
// @Produce json
// @Param client-type-id path string true "client-type-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteClientType(c *gin.Context) {
	client_typeID := c.Param("client-type-id")

	if !util.IsValidUUID(client_typeID) {
		h.handleResponse(c, http.InvalidArgument, "client_type id is an invalid uuid")
		return
	}

	resp, err := h.services.ClientService().DeleteClientType(
		context.Background(),
		&auth_service.ClientTypePrimaryKey{
			Id: client_typeID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AddClient godoc
// @ID create_client
// @Router /client [POST]
// @Summary Create Client
// @Description Create Client
// @Tags Client
// @Accept json
// @Produce json
// @Param client body auth_service.AddClientRequest true "AddClientRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Client} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddClient(c *gin.Context) {
	var client auth_service.AddClientRequest

	err := c.ShouldBindJSON(&client)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().AddClient(
		context.Background(),
		&client,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetClientMatrix godoc
// @ID get_client_matrix
// @Router /client/{project-id} [GET]
// @Summary Get Client Matrix
// @Description Get Client Matrix
// @Tags Client
// @Accept json
// @Produce json
// @Param project-id path string true "project-id"
// @Success 200 {object} http.Response{data=auth_service.GetClientMatrixResponse} "GetClientMatrixBody"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetClientMatrix(c *gin.Context) {
	project_id := c.Param("project-id")

	resp, err := h.services.ClientService().GetClientMatrix(
		context.Background(),
		&auth_service.GetClientMatrixRequest{
			ProjectId: project_id,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateClient godoc
// @ID update_client
// @Router /client [PUT]
// @Summary Update Client
// @Description Update Client
// @Tags Client
// @Accept json
// @Produce json
// @Param client body auth_service.UpdateClientRequest true "UpdateClientRequestBody"
// @Success 200 {object} http.Response{data=auth_service.Client} "Client data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateClient(c *gin.Context) {
	var client auth_service.UpdateClientRequest

	err := c.ShouldBindJSON(&client)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().UpdateClient(
		context.Background(),
		&client,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// RemoveClient godoc
// @ID remove_client
// @Router /client [DELETE]
// @Summary Delete Client
// @Description Get Client
// @Tags Client
// @Accept json
// @Produce json
// @Param remove-client body auth_service.ClientPrimaryKey true "RemoveClientBody"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveClient(c *gin.Context) {
	var removeClient auth_service.ClientPrimaryKey

	err := c.ShouldBindJSON(&removeClient)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	if !util.IsValidUUID(removeClient.ClientPlatformId) {
		h.handleResponse(c, http.InvalidArgument, "client platform id is an invalid uuid")
		return
	}

	if !util.IsValidUUID(removeClient.ClientTypeId) {
		h.handleResponse(c, http.InvalidArgument, "client type id is an invalid uuid")
		return
	}

	resp, err := h.services.ClientService().RemoveClient(
		context.Background(),
		&removeClient,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AddRelation godoc
// @ID create_relation
// @Router /relation [POST]
// @Summary Create Relation
// @Description Create Relation
// @Tags Relation
// @Accept json
// @Produce json
// @Param relation body auth_service.AddRelationRequest true "AddRelationRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Relation} "Relation data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddRelation(c *gin.Context) {
	var relation auth_service.AddRelationRequest

	err := c.ShouldBindJSON(&relation)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().AddRelation(
		context.Background(),
		&relation,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// UpdateRelation godoc
// @ID update_relation
// @Router /relation [PUT]
// @Summary Update Relation
// @Description Update Relation
// @Tags Relation
// @Accept json
// @Produce json
// @Param relation body auth_service.UpdateRelationRequest true "UpdateRelationRequestBody"
// @Success 200 {object} http.Response{data=auth_service.Relation} "Relation data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateRelation(c *gin.Context) {
	var relation auth_service.UpdateRelationRequest

	err := c.ShouldBindJSON(&relation)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().UpdateRelation(
		context.Background(),
		&relation,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// RemoveRelation godoc
// @ID delete_relation
// @Router /relation/{relation-id} [DELETE]
// @Summary Delete Relation
// @Description Get Relation
// @Tags Relation
// @Accept json
// @Produce json
// @Param relation-id path string true "relation-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveRelation(c *gin.Context) {
	relationID := c.Param("relation-id")

	if !util.IsValidUUID(relationID) {
		h.handleResponse(c, http.InvalidArgument, "relation id is an invalid uuid")
		return
	}

	resp, err := h.services.ClientService().RemoveRelation(
		context.Background(),
		&auth_service.RelationPrimaryKey{
			Id: relationID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AddUserInfoField godoc
// @ID create_user_info_field
// @Router /user-info-field [POST]
// @Summary Create UserInfoField
// @Description Create UserInfoField
// @Tags UserInfoField
// @Accept json
// @Produce json
// @Param user-info-field body auth_service.AddUserInfoFieldRequest true "AddUserInfoFieldRequestBody"
// @Success 201 {object} http.Response{data=auth_service.UserInfoField} "UserInfoField data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddUserInfoField(c *gin.Context) {
	var user_info_field auth_service.AddUserInfoFieldRequest

	err := c.ShouldBindJSON(&user_info_field)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().AddUserInfoField(
		context.Background(),
		&user_info_field,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// UpdateUserInfoField godoc
// @ID update_user_info_field
// @Router /user-info-field [PUT]
// @Summary Update UserInfoField
// @Description Update UserInfoField
// @Tags UserInfoField
// @Accept json
// @Produce json
// @Param user-info-field body auth_service.UpdateUserInfoFieldRequest true "UpdateUserInfoFieldRequestBody"
// @Success 200 {object} http.Response{data=auth_service.UserInfoField} "UserInfoField data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUserInfoField(c *gin.Context) {
	var user_info_field auth_service.UpdateUserInfoFieldRequest

	err := c.ShouldBindJSON(&user_info_field)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.ClientService().UpdateUserInfoField(
		context.Background(),
		&user_info_field,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// RemoveUserInfoField godoc
// @ID delete_user_info_field
// @Router /user-info-field/{user-info-field-id} [DELETE]
// @Summary Delete UserInfoField
// @Description Get UserInfoField
// @Tags UserInfoField
// @Accept json
// @Produce json
// @Param user-info-field-id path string true "user-info-field-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveUserInfoField(c *gin.Context) {
	userInfoFieldID := c.Param("user-info-field-id")

	if !util.IsValidUUID(userInfoFieldID) {
		h.handleResponse(c, http.InvalidArgument, "user info field id is an invalid uuid")
		return
	}

	resp, err := h.services.ClientService().RemoveUserInfoField(
		context.Background(),
		&auth_service.UserInfoFieldPrimaryKey{
			Id: userInfoFieldID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
