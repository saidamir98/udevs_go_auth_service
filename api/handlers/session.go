package handlers

import (
	"context"
	"upm/udevs_go_auth_service/api/http"

	"upm/udevs_go_auth_service/genproto/auth_service"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Login
// @Description Login
// @Tags Session
// @Accept json
// @Produce json
// @Param login body auth_service.LoginRequest true "LoginRequestBody"
// @Success 201 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Login(c *gin.Context) {
	var login auth_service.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SessionService().Login(
		context.Background(),
		&login,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// Logout godoc
// @ID logout
// @Router /logout [DELETE]
// @Summary Logout User
// @Description Logout User
// @Tags Session
// @Accept json
// @Produce json
// @Param data body auth_service.LogoutRequest true "LogoutRequest"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Logout(c *gin.Context) {
	var logout auth_service.LogoutRequest

	err := c.ShouldBindJSON(&logout)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SessionService().Logout(
		context.Background(),
		&logout,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// RefreshToken godoc
// @ID refresh
// @Router /refresh [PUT]
// @Summary Refresh Token
// @Description Refresh Token
// @Tags Session
// @Accept json
// @Produce json
// @Param user body auth_service.RefreshTokenRequest true "RefreshTokenRequestBody"
// @Success 200 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RefreshToken(c *gin.Context) {
	var user auth_service.RefreshTokenRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SessionService().RefreshToken(
		context.Background(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// HasAccess godoc
// @ID has_access
// @Router /has-access [POST]
// @Summary Has Access
// @Description Has Access
// @Tags Session
// @Accept json
// @Produce json
// @Param has-access body auth_service.HasAccessRequest true "HasAccessRequestBody"
// @Success 201 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) HasAccess(c *gin.Context) {
	var login auth_service.HasAccessRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SessionService().HasAccess(
		context.Background(),
		&login,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}
