package handlers

import (
	"upm/udevs_go_auth_service/api/http"

	"upm/udevs_go_auth_service/genproto/auth_service"

	"github.com/saidamir98/udevs_pkg/util"

	"github.com/gin-gonic/gin"

	structpb "google.golang.org/protobuf/types/known/structpb"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Param user body auth_service.CreateUserRequest true "CreateUserRequestBody"
// @Success 201 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var user auth_service.CreateUserRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().CreateUser(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetUserList godoc
// @ID get_user_list
// @Router /user [GET]
// @Summary Get User List
// @Description  Get User List
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Param client-platform-id query string false "client-platform-id"
// @Param client-type-id query string false "client-type-id"
// @Param project-id query string false "project-id"
// @Success 200 {object} http.Response{data=auth_service.GetUserListResponse} "GetUserListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserList(c *gin.Context) {
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

	resp, err := h.services.UserService().GetUserList(
		c.Request.Context(),
		&auth_service.GetUserListRequest{
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

// GetUserByID godoc
// @ID get_user_by_id
// @Router /user/{user-id} [GET]
// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Success 200 {object} http.Response{data=auth_service.User} "UserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserByID(c *gin.Context) {
	userID := c.Param("user-id")

	if !util.IsValidUUID(userID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.UserService().GetUserByID(
		c.Request.Context(),
		&auth_service.UserPrimaryKey{
			Id: userID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateUser godoc
// @ID update_user
// @Router /user [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param user body auth_service.UpdateUserRequest true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {
	var user auth_service.UpdateUserRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().UpdateUser(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteUser godoc
// @ID delete_user
// @Router /user/{user-id} [DELETE]
// @Summary Delete User
// @Description Get User
// @Tags User
// @Accept json
// @Produce json
// @Param user-id path string true "user-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("user-id")

	if !util.IsValidUUID(userID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.UserService().DeleteUser(
		c.Request.Context(),
		&auth_service.UserPrimaryKey{
			Id: userID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AddUserRelation godoc
// @ID add_user_relation
// @Router /user-relation [POST]
// @Summary Create UserRelation
// @Description Create UserRelation
// @Tags UserRelation
// @Accept json
// @Produce json
// @Param user-relation body auth_service.AddUserRelationRequest true "AddUserRelationRequestBody"
// @Success 201 {object} http.Response{data=auth_service.UserRelation} "UserRelation data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddUserRelation(c *gin.Context) {
	var user_relation auth_service.AddUserRelationRequest

	err := c.ShouldBindJSON(&user_relation)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().AddUserRelation(
		c.Request.Context(),
		&user_relation,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// RemoveUserRelation godoc
// @ID delete_user_relation
// @Router /user-relation [DELETE]
// @Summary Delete UserRelation
// @Description Get UserRelation
// @Tags UserRelation
// @Accept json
// @Produce json
// @Param user-relation body auth_service.UserRelationPrimaryKey true "UserRelationPrimaryKeyBody"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveUserRelation(c *gin.Context) {
	var user_relation auth_service.UserRelationPrimaryKey

	err := c.ShouldBindJSON(&user_relation)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().RemoveUserRelation(
		c.Request.Context(),
		&user_relation,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// // UpsertUserInfo godoc
// // @ID upsert_user_info
// // @Router /upsert-user-info/{user-id} [POST]
// // @Summary Upsert UserInfo
// // @Description Upsert UserInfo
// // @Tags UpsertUserInfo
// // @Accept json
// // @Produce json
// // @Param data body structpb.Struct true "UpsertUserInfoRequestBody"
// // @Param user-id path string true "user-id"
// // @Success 201 {object} http.Response{data=auth_service.Role} "Role data"
// // @Response 400 {object} http.Response{data=string} "Bad Request"
// // @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpsertUserInfo(c *gin.Context) {
	var data structpb.Struct
	userID := c.Param("user-id")

	if !util.IsValidUUID(userID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&data)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().UpsertUserInfo(
		c.Request.Context(),
		&auth_service.UpsertUserInfoRequest{
			UserId: userID,
			Data:   &data,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// UpdateUser godoc
// @ID reset_password
// @Router /user/reset-password [PUT]
// @Summary Update User
// @Description Reset Password
// @Tags User
// @Accept json
// @Produce json
// @Param reset_password body auth_service.ResetPasswordRequest true "ResetPasswordRequestBody"
// @Success 200 {object} http.Response{data=auth_service.LoginResponse} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) ResetPassword(c *gin.Context) {
	var user auth_service.ResetPasswordRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().ResetPassword(
		c.Request.Context(),
		&user,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	login, err := h.services.SessionService().Login(
		c.Request.Context(),
		&auth_service.LoginRequest{
			Username: resp.GetLogin(),
			Password: user.GetPassword(),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, login)
}

// UpdateUser godoc
// @ID send_message_to_user_email
// @Router /user/send-message [POST]
// @Summary Send Message To User
// @Description Send Message to User Email
// @Tags User
// @Accept json
// @Produce json
// @Param send_message body auth_service.SendMessageToEmailRequest true "SendMessageToEmailRequestBody"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) SendMessageToUserEmail(c *gin.Context) {
	var customerMessage auth_service.SendMessageToEmailRequest

	err := c.ShouldBindJSON(&customerMessage)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().SendMessageToEmail(
		c.Request.Context(),
		&customerMessage,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
