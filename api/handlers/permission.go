package handlers

import (
	"upm/udevs_go_auth_service/api/http"

	"upm/udevs_go_auth_service/genproto/auth_service"

	"github.com/saidamir98/udevs_pkg/util"

	"github.com/gin-gonic/gin"
)

// AddRole godoc
// @ID create_role
// @Router /role [POST]
// @Summary Create Role
// @Description Create Role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body auth_service.AddRoleRequest true "AddRoleRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Role} "Role data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddRole(c *gin.Context) {
	var role auth_service.AddRoleRequest

	err := c.ShouldBindJSON(&role)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().AddRole(
		c.Request.Context(),
		&role,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetRoleById godoc
// @ID get_role_by_id
// @Router /role/{role-id} [GET]
// @Summary Get Role By ID
// @Description Get Role By ID
// @Tags Role
// @Accept json
// @Produce json
// @Param role-id path string true "role-id"
// @Success 200 {object} http.Response{data=auth_service.CompleteClientType} "ClientTypeBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetRoleByID(c *gin.Context) {
	roleId := c.Param("role-id")
	if !util.IsValidUUID(roleId) {
		h.handleResponse(c, http.InvalidArgument, "role id is an invalid uuid")
		return
	}

	resp, err := h.services.PermissionService().GetRoleById(c.Request.Context(), &auth_service.RolePrimaryKey{
		Id: roleId,
	})

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// GetRolesList godoc
// @ID get_roles_list
// @Router /role [GET]
// @Summary Get Roles List
// @Description  Get Roles List
// @Tags Role
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param client-platform-id query string false "client-platform-id"
// @Param client-type-id query string false "client-type-id"
// @Success 200 {object} http.Response{data=auth_service.GetRolesResponse} "GetRolesListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetRolesList(c *gin.Context) {
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

	resp, err := h.services.PermissionService().GetRolesList(
		c.Request.Context(),
		&auth_service.GetRolesListRequest{
			Offset:           uint32(offset),
			Limit:            uint32(limit),
			ClientPlatformId: c.Query("client-platform-id"),
			ClientTypeId:     c.Query("client-type-id"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdateRole godoc
// @ID update_role
// @Router /role [PUT]
// @Summary Update Role
// @Description Update Role
// @Tags Role
// @Accept json
// @Produce json
// @Param role body auth_service.UpdateRoleRequest true "UpdateRoleRequestBody"
// @Success 200 {object} http.Response{data=auth_service.Role} "Role data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateRole(c *gin.Context) {
	var role auth_service.UpdateRoleRequest

	err := c.ShouldBindJSON(&role)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().UpdateRole(
		c.Request.Context(),
		&role,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// RemoveRole godoc
// @ID delete_role
// @Router /role/{role-id} [DELETE]
// @Summary Delete Role
// @Description Get Role
// @Tags Role
// @Accept json
// @Produce json
// @Param role-id path string true "role-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveRole(c *gin.Context) {
	roleID := c.Param("role-id")

	if !util.IsValidUUID(roleID) {
		h.handleResponse(c, http.InvalidArgument, "role id is an invalid uuid")
		return
	}

	resp, err := h.services.PermissionService().RemoveRole(
		c.Request.Context(),
		&auth_service.RolePrimaryKey{
			Id: roleID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// CreatePermission godoc
// @ID create_permission
// @Router /permission [POST]
// @Summary Create Permission
// @Description Create Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission body auth_service.CreatePermissionRequest true "CreatePermissionRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Permission} "Permission data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreatePermission(c *gin.Context) {
	var permission auth_service.CreatePermissionRequest

	err := c.ShouldBindJSON(&permission)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().CreatePermission(
		c.Request.Context(),
		&permission,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetPermissionList godoc
// @ID get_permission_list
// @Router /permission [GET]
// @Summary Get Permission List
// @Description  Get Permission List
// @Tags Permission
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=auth_service.GetPermissionListResponse} "GetPermissionListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPermissionList(c *gin.Context) {
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

	resp, err := h.services.PermissionService().GetPermissionList(
		c.Request.Context(),
		&auth_service.GetPermissionListRequest{
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

// GetPermissionByID godoc
// @ID get_permission_by_id
// @Router /permission/{permission-id} [GET]
// @Summary Get Permission By ID
// @Description Get Permission By ID
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission-id path string true "permission-id"
// @Success 200 {object} http.Response{data=auth_service.GetPermissionByIDResponse} "PermissionBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetPermissionByID(c *gin.Context) {
	permissionID := c.Param("permission-id")

	if !util.IsValidUUID(permissionID) {
		h.handleResponse(c, http.InvalidArgument, "permission id is an invalid uuid")
		return
	}

	resp, err := h.services.PermissionService().GetPermissionByID(
		c.Request.Context(),
		&auth_service.PermissionPrimaryKey{
			Id: permissionID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// UpdatePermission godoc
// @ID update_permission
// @Router /permission [PUT]
// @Summary Update Permission
// @Description Update Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission body auth_service.UpdatePermissionRequest true "UpdatePermissionRequestBody"
// @Success 200 {object} http.Response{data=auth_service.Permission} "Permission data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePermission(c *gin.Context) {
	var permission auth_service.UpdatePermissionRequest

	err := c.ShouldBindJSON(&permission)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().UpdatePermission(
		c.Request.Context(),
		&permission,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeletePermission godoc
// @ID delete_permission
// @Router /permission/{permission-id} [DELETE]
// @Summary Delete Permission
// @Description Get Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission-id path string true "permission-id"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeletePermission(c *gin.Context) {
	permissionID := c.Param("permission-id")

	if !util.IsValidUUID(permissionID) {
		h.handleResponse(c, http.InvalidArgument, "permission id is an invalid uuid")
		return
	}

	resp, err := h.services.PermissionService().DeletePermission(
		c.Request.Context(),
		&auth_service.PermissionPrimaryKey{
			Id: permissionID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// UpsertScope godoc
// @ID upsert_scope
// @Router /upsert-scope [POST]
// @Summary Upsert Scope
// @Description Upsert Scope
// @Tags Scope
// @Accept json
// @Produce json
// @Param upsert-scope body auth_service.UpsertScopeRequest true "UpsertScopeRequestBody"
// @Success 201 {object} http.Response{data=auth_service.Role} "Role data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpsertScope(c *gin.Context) {
	var upsert_scope auth_service.UpsertScopeRequest

	err := c.ShouldBindJSON(&upsert_scope)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().UpsertScope(
		c.Request.Context(),
		&upsert_scope,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetScopesList godoc
// @ID get_scopes_list
// @Router /scope [GET]
// @Summary Get Scopes List
// @Description  Get Scopes List
// @Tags Scope
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param client-platform-id query string true "client-platform-id"
// @Param search query string false "search"
// @Param order_by query string false "order_by"
// @Param order_type query string false "order_type"
// @Success 200 {object} http.Response{data=auth_service.GetScopesResponse} "GetScopesListResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetScopesList(c *gin.Context) {
	clientPlatformID := c.Query("client-platform-id")
	if !util.IsValidUUID(clientPlatformID) {
		h.handleResponse(c, http.InvalidArgument, "Client platform id is an invalid uuid")
		return
	}

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

	resp, err := h.services.PermissionService().GetScopeList(
		c.Request.Context(),
		&auth_service.GetScopeListRequest{
			Offset:           uint32(offset),
			Limit:            uint32(limit),
			Search:           c.Query("search"),
			OrderBy:          c.Query("order_by"),
			OrderType:        c.Query("order_type"),
			ClientPlatformId: clientPlatformID,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// AddPermissionScope godoc
// @ID add_permission_scope
// @Router /permission-scope [POST]
// @Summary Create PermissionScope
// @Description Create PermissionScope
// @Tags PermissionScope
// @Accept json
// @Produce json
// @Param permission-scope body auth_service.AddPermissionScopeRequest true "AddPermissionScopeRequestBody"
// @Success 201 {object} http.Response{data=auth_service.PermissionScope} "PermissionScope data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddPermissionScope(c *gin.Context) {
	var permission_scope auth_service.AddPermissionScopeRequest

	err := c.ShouldBindJSON(&permission_scope)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().AddPermissionScope(
		c.Request.Context(),
		&permission_scope,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// RemovePermissionScope godoc
// @ID delete_permission_scope
// @Router /permission-scope [DELETE]
// @Summary Delete PermissionScope
// @Description Get PermissionScope
// @Tags PermissionScope
// @Accept json
// @Produce json
// @Param permission-scope body auth_service.PermissionScopePrimaryKey true "PermissionScopePrimaryKeyBody"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemovePermissionScope(c *gin.Context) {
	var permission_scope auth_service.PermissionScopePrimaryKey

	err := c.ShouldBindJSON(&permission_scope)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().RemovePermissionScope(
		c.Request.Context(),
		&permission_scope,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AddRolePermission godoc
// @ID add_role_permission
// @Router /role-permission [POST]
// @Summary Create RolePermission
// @Description Create RolePermission
// @Tags RolePermission
// @Accept json
// @Produce json
// @Param role-permission body auth_service.AddRolePermissionRequest true "AddRolePermissionRequestBody"
// @Success 201 {object} http.Response{data=auth_service.RolePermission} "RolePermission data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddRolePermission(c *gin.Context) {
	var role_permission auth_service.AddRolePermissionRequest

	err := c.ShouldBindJSON(&role_permission)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().AddRolePermission(
		c.Request.Context(),
		&role_permission,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// AddRolePermission godoc
// @ID add_role_permissions
// @Router /role-permission/many [POST]
// @Summary Create RolePermissions
// @Description Create RolePermissions
// @Tags RolePermission
// @Accept json
// @Produce json
// @Param role-permission body auth_service.AddRolePermissionsRequest true "AddRolePermissionsRequestBody"
// @Success 201 {object} http.Response{data=auth_service.AddRolePermissionsResponse} "RolePermission Added Amount"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) AddRolePermissions(c *gin.Context) {
	var role_permissions auth_service.AddRolePermissionsRequest

	err := c.ShouldBindJSON(&role_permissions)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().AddRolePermissions(
		c.Request.Context(),
		&role_permissions,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// RemoveRolePermission godoc
// @ID delete_role_permission
// @Router /role-permission [DELETE]
// @Summary Delete RolePermission
// @Description Get RolePermission
// @Tags RolePermission
// @Accept json
// @Produce json
// @Param role-permission body auth_service.RolePermissionPrimaryKey true "RolePermissionPrimaryKeyBody"
// @Success 204
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) RemoveRolePermission(c *gin.Context) {
	var role_permission auth_service.RolePermissionPrimaryKey

	err := c.ShouldBindJSON(&role_permission)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.PermissionService().RemoveRolePermission(
		c.Request.Context(),
		&role_permission,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}

// AddPermissionScope godoc
// @ID permission_generated
// @Router /permission_generated [POST]
// @Summary Generate Permission
// @Description Generate Permission
// @Tags Permission
// @Accept json
// @Produce json
// @Param permission-scope body auth_service.PermissionGenerated true "AddPermissionScopeRequestBody"
// @Success 201 {object} http.Response{data=auth_service.PermissionScope} "PermissionScope data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) PermissionGeneratedPermission(c *gin.Context) {

}
