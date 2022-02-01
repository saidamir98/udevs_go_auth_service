package handlers

import (
	"context"
	"upm/udevs_go_auth_service/api/http"

	"upm/udevs_go_auth_service/genproto/settings_service"

	"upm/udevs_go_auth_service/pkg/util"

	"github.com/gin-gonic/gin"
)

// CreatePosition godoc
// @ID create_position
// @Router /position [POST]
// @Summary Create Position
// @Description Create Position
// @Tags Position
// @Accept json
// @Produce json
// @Param position body settings_service.CreatePositionRequest true "position"
// @Success 201 {object} http.Response{data=settings_service.Position} "PositionBody"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreatePosition(c *gin.Context) {
	var position settings_service.CreatePositionRequest

	err := c.ShouldBindJSON(&position)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PositionService().Create(
		context.Background(),
		&position,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetPositionList godoc
// @ID get_position_list
// @Router /position [GET]
// @Summary Get Position List
// @Description  Get Position List
// @Tags Position
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=settings_service.GetPositionListResponse} "GetPositionListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPositionList(c *gin.Context) {
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

	resp, err := h.services.PositionService().GetList(
		context.Background(),
		&settings_service.GetPositionListRequest{
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

// GetPositionByID godoc
// @ID get_position_by_id
// @Router /position/{position-id} [GET]
// @Summary Get Position By ID
// @Description Get Position By ID
// @Tags Position
// @Accept json
// @Produce json
// @Param position-id path string true "position-id"
// @Success 200 {object} http.Response{data=settings_service.Position} "PositionBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPositionByID(c *gin.Context) {
	positionID := c.Param("position-id")

	if !util.IsValidUUID(positionID) {
		h.handleResponse(c, http.InvalidArgument, "position id is an invalid uuid")
		return
	}

	resp, err := h.services.PositionService().GetByID(
		context.Background(),
		&settings_service.PositionPrimaryKey{
			Id: positionID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdatePosition godoc
// @ID update_position
// @Router /position [PUT]
// @Summary Update Position
// @Description Update Position
// @Tags Position
// @Accept json
// @Produce json
// @Param position body settings_service.UpdatePositionRequest true "position"
// @Success 200 {object} http.Response{data=settings_service.Position} "PositionBody"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePosition(c *gin.Context) {
	var position settings_service.UpdatePositionRequest

	err := c.ShouldBindJSON(&position)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PositionService().Update(
		context.Background(),
		&position,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeletePosition godoc
// @ID delete_position
// @Router /position/{position-id} [DELETE]
// @Summary Delete Position
// @Description Get Position
// @Tags Position
// @Accept json
// @Produce json
// @Param position-id path string true "position-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeletePosition(c *gin.Context) {
	positionID := c.Param("position-id")

	if !util.IsValidUUID(positionID) {
		h.handleResponse(c, http.InvalidArgument, "position id is an invalid uuid")
		return
	}

	resp, err := h.services.PositionService().Delete(
		context.Background(),
		&settings_service.PositionPrimaryKey{
			Id: positionID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AddLevelItem godoc
// @ID add_level_item
// @Router /position/level [POST]
// @Summary Add Level Item
// @Description Add Level Item
// @Tags Position
// @Accept json
// @Produce json
// @Param position body settings_service.AddLevelItemRequest true "AddLevelItemRequestBody"
// @Success 201 {object} http.Response{data=settings_service.Position} "PositionBody"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddLevelItem(c *gin.Context) {
	var levelItem settings_service.AddLevelItemRequest

	err := c.ShouldBindJSON(&levelItem)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PositionService().AddLevelItem(
		context.Background(),
		&levelItem,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// UpdateLevelItem godoc
// @ID update_level_item
// @Router /position/level [PUT]
// @Summary Update Level Item
// @Description Update Level Item
// @Tags Position
// @Accept json
// @Produce json
// @Param position body settings_service.UpdateLevelItemRequest true "UpdateLevelItemRequestBody"
// @Success 200 {object} http.Response{data=settings_service.Position} "PositionBody"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateLevelItem(c *gin.Context) {
	var levelItem settings_service.UpdateLevelItemRequest

	err := c.ShouldBindJSON(&levelItem)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PositionService().UpdateLevelItem(
		context.Background(),
		&levelItem,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// RemoveLevelItem godoc
// @ID remove_level_item
// @Router /position/{position-id}/level/{level-id} [DELETE]
// @Summary Remove Level Item
// @Description Remove Level Item
// @Tags Position
// @Accept json
// @Produce json
// @Param position-id path string true "position-id"
// @Param level-id path string true "level-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveLevelItem(c *gin.Context) {
	positionID := c.Param("position-id")
	if !util.IsValidUUID(positionID) {
		h.handleResponse(c, http.InvalidArgument, "position id is an invalid uuid")
		return
	}

	levelID := c.Param("level-id")
	if !util.IsValidUUID(levelID) {
		h.handleResponse(c, http.InvalidArgument, "level id is an invalid uuid")
		return
	}

	resp, err := h.services.PositionService().RemoveLevelItem(
		context.Background(),
		&settings_service.RemoveLevelItemRequest{
			Id:         levelID,
			PositionId: positionID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
