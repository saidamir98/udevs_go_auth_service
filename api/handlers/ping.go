package handlers

import (
	"upm/udevs_go_auth_service/api/http"
	"upm/udevs_go_auth_service/config"

	"github.com/gin-gonic/gin"
)

// Ping godoc
// @ID ping
// @Router /ping [GET]
// @Summary returns "pong" message
// @Description this returns "pong" messsage to show service is working
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=string} "Response data"
// @Failure 500 {object} http.Response{}
func (h *Handler) Ping(c *gin.Context) {
	h.handleResponse(c, http.OK, "pong")
}

// GetConfig godoc
// @ID get_config
// @Router /config [GET]
// @Summary get config data on the debug mode
// @Description show service config data when the service environment set to debug mode
// @Accept json
// @Produce json
// @Success 200 {object} http.Response{data=config.Config} "Response data"
// @Failure 400 {object} http.Response{}
func (h *Handler) GetConfig(c *gin.Context) {
	switch h.cfg.Environment {
	case config.DebugMode:
		h.handleResponse(c, http.OK, h.cfg)
		return
	case config.TestMode:
		h.handleResponse(c, http.OK, h.cfg.Environment)
		return
	case config.ReleaseMode:
		h.handleResponse(c, http.OK, "private data")
		return
	}

	h.handleResponse(c, http.BadEnvironment, "wrong environment value passed")
}
