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
// @Router /integration/{integration-id}/session [GET]
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

	h.handleResponse(c, http.OK, resp)
}

// GetIntegrationSessions godoc
// @ID get_integration_sessions
// @Router /integration/{integration-id}/session [GET]
// @Summary Get Integration Sessions
// @Description  Get Integration Sessions
// @Tags Integration
// @Accept json
// @Produce json
// @Param integration-id path string true "integration-id"
// @Success 200 {object} http.Response{data=auth_service.GetIntegrationSessionsResponse} "GetIntegrationSessionsResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetIntegrationSessions(c *gin.Context) {
	resp, err := h.services.IntegrationService().GetIntegrationSessions(
		c.Request.Context(),
		&auth_service.IntegrationPrimaryKey{
			Id: c.Param("integration-id"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetIntegrationToken godoc
// @ID get_integration_token
// @Router /integration/{integration-id}/session [POST]
// @Summary Add Session To Integration
// @Description Add Session To Integration
// @Tags Integration
// @Accept json
// @Produce json
// @Param integration-id path string true "integration-id"
// @Param addSessionToIntegration body auth_service.AddSessionToIntegrationRequest true "AddSessionToIntegrationRequestBody"
// @Success 201 {object} http.Response{data=auth_service.AddSessionToIntegrationResponse} "Add Session To Integration Response"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddSessionToIntegration(c *gin.Context) {
	var login auth_service.AddSessionToIntegrationRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	integrationID := c.Param("integration-id")
	if !util.IsValidUUID(integrationID) {
		h.handleResponse(c, http.InvalidArgument, "integration id is an invalid uuid")
		return
	}
	login.IntegrationId = integrationID

	resp, err := h.services.IntegrationService().AddSessionToIntegration(
		c.Request.Context(),
		&login,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetIntegrationByID godoc
// @ID get_Integration_by_id
// @Router /integration/{integration-id} [GET]
// @Summary Get Integration By ID
// @Description Get Integration By ID
// @Tags Integration
// @Accept json
// @Produce json
// @Param integration-id path string true "integration-id"
// @Success 200 {object} http.Response{data=auth_service.Integration} "IntegrationBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetIntegrationByID(c *gin.Context) {
	IntegrationID := c.Param("integration-id")

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

// DeleteIntegration godoc
// @ID delete_Integration
// @Router /integration/{integration-id} [DELETE]
// @Summary Delete Integration
// @Description Delete Integration
// @Tags Integration
// @Accept json
// @Produce json
// @Param integration-id path string true "Integration-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteIntegration(c *gin.Context) {
	IntegrationID := c.Param("integration-id")

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

// GetIntegrationByID godoc
// @ID get_integration_token
// @Router /integration/{integration-id}/session/{session-id} [GET]
// @Summary Get Integration Token
// @Description Get Integration Token
// @Tags Integration
// @Accept json
// @Produce json
// @Param integration-id path string true "integration-id"
// @Param session-id path string true "session-id"
// @Success 200 {object} http.Response{data=auth_service.Token} "IntegrationBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetIntegrationToken(c *gin.Context) {
	IntegrationID := c.Param("integration-id")
	SessionID := c.Param("session_id")

	if !util.IsValidUUID(IntegrationID) {
		h.handleResponse(c, http.InvalidArgument, "Integration id is an invalid uuid")
		return
	}

	resp, err := h.services.IntegrationService().GetIntegrationToken(
		c.Request.Context(),
		&auth_service.GetIntegrationTokenRequest{
			IntegrationId: IntegrationID,
			SessionId:     SessionID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteIntegration godoc
// @ID delete_session_from_integration
// @Router /integration/{integration-id}/session/{session-id} [DELETE]
// @Summary Delete Session From Integration
// @Description Delete Session From Integration
// @Tags Integration
// @Accept json
// @Produce json
// @Param integration-id path string true "Integration-id"
// @Param session-id path string true "session-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveSessionFromIntegration(c *gin.Context) {
	integrationID := c.Param("integration-id")
	sessionID := c.Param("session-id")

	if !util.IsValidUUID(integrationID) {
		h.handleResponse(c, http.InvalidArgument, "Integration id is an invalid uuid")
		return
	}

	if !util.IsValidUUID(sessionID) {
		h.handleResponse(c, http.InvalidArgument, "Session id is an invalid uuid")
		return
	}

	resp, err := h.services.IntegrationService().DeleteSessionFromIntegration(
		c.Request.Context(),
		&auth_service.GetIntegrationTokenRequest{
			IntegrationId: integrationID,
			SessionId:     sessionID,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
