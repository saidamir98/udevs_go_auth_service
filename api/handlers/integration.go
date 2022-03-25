package handlers

import (
	"upm/udevs_go_auth_service/api/http"

	"upm/udevs_go_auth_service/genproto/auth_service"

	"upm/udevs_go_auth_service/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateIntegration godoc
// @ID create_Integration
// @Router /integration [POST]
// @Summary Create Integration
// @Description Create Integration
// @Tags Integration
// @Accept json
// @Produce json
// @Param Integration body auth_service.CreateIntegrationRequest true "CreateIntegrationRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Integration} "Integration data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateIntegration(c *gin.Context) {
	var Integration auth_service.CreateIntegrationRequest

	err := c.ShouldBindJSON(&Integration)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.IntegrationService().CreateIntegration(
		c.Request.Context(),
		&Integration,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetIntegrationList godoc
// @ID get_integration_list
// @Router /integration [GET]
// @Summary Get Integration List
// @Description  Get Integration List
// @Tags Integration
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param client-platform-id query string false "client-platform-id"
// @Param client-type-id query string false "client-type-id"
// @Param project-id query string false "project-id"
// @Success 200 {object} http.Response{data=auth_service.GetIntegrationListResponse} "GetIntegrationListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetIntegrationList(c *gin.Context) {
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

	println("a")
	resp, err := h.services.IntegrationService().GetIntegrationList(
		c.Request.Context(),
		&auth_service.GetIntegrationListRequest{
			Limit:            int32(limit),
			Offset:           int32(offset),
			Search:           c.Query("search"),
			ClientPlatformId: c.Query("client-platform-id"),
			ClientTypeId:     c.Query("client-type-id"),
			ProjectId:        c.Query("project-id"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}
	println("b")

	h.handleResponse(c, http.OK, resp)
}

// GetIntegrationByID godoc
// @ID get_Integration_by_id
// @Router /integration/{Integration-id} [GET]
// @Summary Get Integration By ID
// @Description Get Integration By ID
// @Tags Integration
// @Accept json
// @Produce json
// @Param Integration-id path string true "Integration-id"
// @Success 200 {object} http.Response{data=auth_service.Integration} "IntegrationBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetIntegrationByID(c *gin.Context) {
	IntegrationID := c.Param("Integration-id")

	if !util.IsValidUUID(IntegrationID) {
		h.handleResponse(c, http.InvalidArgument, "Integration id is an invalid uuid")
		return
	}

	resp, err := h.services.IntegrationService().GetIntegrationByID(
		c.Request.Context(),
		&auth_service.IntegrationPrimaryKey{
			Id: IntegrationID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateIntegration godoc
// @ID update_Integration
// @Router /integration [PUT]
// @Summary Update Integration
// @Description Update Integration
// @Tags Integration
// @Accept json
// @Produce json
// @Param Integration body auth_service.UpdateIntegrationRequest true "UpdateIntegrationRequestBody"
// @Success 200 {object} http.Response{data=auth_service.Integration} "Integration data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateIntegration(c *gin.Context) {
	var Integration auth_service.UpdateIntegrationRequest

	err := c.ShouldBindJSON(&Integration)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.IntegrationService().UpdateIntegration(
		c.Request.Context(),
		&Integration,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteIntegration godoc
// @ID delete_Integration
// @Router /integration/{Integration-id} [DELETE]
// @Summary Delete Integration
// @Description Get Integration
// @Tags Integration
// @Accept json
// @Produce json
// @Param Integration-id path string true "Integration-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteIntegration(c *gin.Context) {
	IntegrationID := c.Param("Integration-id")

	if !util.IsValidUUID(IntegrationID) {
		h.handleResponse(c, http.InvalidArgument, "Integration id is an invalid uuid")
		return
	}

	resp, err := h.services.IntegrationService().DeleteIntegration(
		c.Request.Context(),
		&auth_service.IntegrationPrimaryKey{
			Id: IntegrationID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
