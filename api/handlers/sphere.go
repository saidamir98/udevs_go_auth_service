package handlers

import (
	"upm/udevs_go_auth_service/api/http"

	"upm/udevs_go_auth_service/genproto/settings_service"

	"upm/udevs_go_auth_service/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreateSphere godoc
// @ID create_sphere
// @Router /sphere [POST]
// @Summary Create Sphere
// @Description Create Sphere
// @Tags Sphere
// @Accept json
// @Produce json
// @Param sphere body settings_service.CreateSphereRequest true "CreateSphereRequestBody"
// @Success 201 {object} http.Response{data=settings_service.Sphere} "Sphere data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateSphere(c *gin.Context) {
	var sphere settings_service.CreateSphereRequest

	err := c.ShouldBindJSON(&sphere)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SphereService().Create(
		c.Request.Context(),
		&sphere,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetSphereList godoc
// @ID get_sphere_list
// @Router /sphere [GET]
// @Summary Get Sphere List
// @Description  Get Sphere List
// @Tags Sphere
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=settings_service.GetSphereListResponse} "GetSphereListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSphereList(c *gin.Context) {
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

	resp, err := h.services.SphereService().GetList(
		c.Request.Context(),
		&settings_service.GetSphereListRequest{
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

// GetSphereByID godoc
// @ID get_sphere_by_id
// @Router /sphere/{sphere-id} [GET]
// @Summary Get Sphere By ID
// @Description Get Sphere By ID
// @Tags Sphere
// @Accept json
// @Produce json
// @Param sphere-id path string true "sphere-id"
// @Success 200 {object} http.Response{data=settings_service.Sphere} "SphereBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetSphereByID(c *gin.Context) {
	sphereID := c.Param("sphere-id")

	if !util.IsValidUUID(sphereID) {
		h.handleResponse(c, http.InvalidArgument, "sphere id is an invalid uuid")
		return
	}

	resp, err := h.services.SphereService().GetByID(
		c.Request.Context(),
		&settings_service.SpherePrimaryKey{
			Id: sphereID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateSphere godoc
// @ID update_sphere
// @Router /sphere [PUT]
// @Summary Update Sphere
// @Description Update Sphere
// @Tags Sphere
// @Accept json
// @Produce json
// @Param sphere body settings_service.UpdateSphereRequest true "UpdateSphereRequestBody"
// @Success 200 {object} http.Response{data=settings_service.Sphere} "Sphere data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateSphere(c *gin.Context) {
	var sphere settings_service.UpdateSphereRequest

	err := c.ShouldBindJSON(&sphere)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.SphereService().Update(
		c.Request.Context(),
		&sphere,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteSphere godoc
// @ID delete_sphere
// @Router /sphere/{sphere-id} [DELETE]
// @Summary Delete Sphere
// @Description Get Sphere
// @Tags Sphere
// @Accept json
// @Produce json
// @Param sphere-id path string true "sphere-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteSphere(c *gin.Context) {
	sphereID := c.Param("sphere-id")

	if !util.IsValidUUID(sphereID) {
		h.handleResponse(c, http.InvalidArgument, "sphere id is an invalid uuid")
		return
	}

	resp, err := h.services.SphereService().Delete(
		c.Request.Context(),
		&settings_service.SpherePrimaryKey{
			Id: sphereID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
