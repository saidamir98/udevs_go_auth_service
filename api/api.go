package api

import (
	"upm/udevs_go_auth_service/api/docs"
	"upm/udevs_go_auth_service/api/handlers"
	"upm/udevs_go_auth_service/config"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetUpRouter godoc
// @description This is a api gateway
// @termsOfService https://udevs.io
func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	// docs.SwaggerInfo.Host = cfg.ServiceHost + cfg.HTTPPort
	docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	r.Use(customCORSMiddleware())

	r.GET("/ping", h.Ping)
	r.GET("/config", h.GetConfig)

	r.POST("/sphere", h.CreateSphere)
	r.GET("/sphere", h.GetSphereList)
	r.GET("/sphere/:sphere-id", h.GetSphereByID)
	r.PUT("/sphere", h.UpdateSphere)
	r.DELETE("/sphere/:sphere-id", h.DeleteSphere)

	r.POST("/position", h.CreatePosition)
	r.GET("/position", h.GetPositionList)
	r.GET("/position/:position-id", h.GetPositionByID)
	r.PUT("/position", h.UpdatePosition)
	r.DELETE("/position/:position-id", h.DeletePosition)
	r.POST("/position/level", h.AddLevelItem)
	r.PUT("/position/level", h.UpdateLevelItem)
	r.DELETE("/position/:position-id/level/:level-id", h.RemoveLevelItem)

	// CLIENT SERVICE
	r.POST("/client-platform", h.CreateClientPlatform)
	r.GET("/client-platform", h.GetClientPlatformList)
	r.GET("/client-platform/:client-platform-id", h.GetClientPlatformByID)
	r.GET("/client-platform-detailed/:client-platform-id", h.GetClientPlatformByIDDetailed)
	r.PUT("/client-platform", h.UpdateClientPlatform)
	r.DELETE("/client-platform/:client-platform-id", h.DeleteClientPlatform)

	r.POST("/client-type", h.CreateClientType)
	r.GET("/client-type", h.GetClientTypeList)
	r.GET("/client-type/:client-type-id", h.GetClientTypeByID)
	r.PUT("/client-type", h.UpdateClientType)
	r.DELETE("/client-type/:client-type-id", h.DeleteClientType)

	r.POST("/client", h.AddClient)
	r.GET("/client/:project-id", h.GetClientMatrix)
	r.PUT("/client", h.UpdateClient)
	r.DELETE("/client", h.RemoveClient)

	r.POST("/relation", h.AddRelation)
	r.PUT("/relation", h.UpdateRelation)
	r.DELETE("/relation/:relation-id", h.RemoveRelation)

	r.POST("/user-info-field", h.AddUserInfoField)
	r.PUT("/user-info-field", h.UpdateUserInfoField)
	r.DELETE("/user-info-field/:user-info-field-id", h.RemoveUserInfoField)

	// PERMISSION SERVICE
	r.GET("/role/:role-id", h.GetRoleByID)
	r.GET("/role", h.GetRolesList)
	r.POST("/role", h.AddRole)
	r.PUT("/role", h.UpdateRole)
	r.DELETE("/role/:role-id", h.RemoveRole)

	r.POST("/permission", h.CreatePermission)
	r.GET("/permission", h.GetPermissionList)
	r.GET("/permission/:permission-id", h.GetPermissionByID)
	r.PUT("/permission", h.UpdatePermission)
	r.DELETE("/permission/:permission-id", h.DeletePermission)

	r.POST("/upsert-scope", h.UpsertScope)

	r.POST("/permission-scope", h.AddPermissionScope)
	r.DELETE("/permission-scope", h.RemovePermissionScope)

	r.POST("/role-permission", h.AddRolePermission)
	r.POST("/role-permission/many", h.AddRolePermissions)
	r.DELETE("/role-permission", h.RemoveRolePermission)

	r.POST("/user", h.CreateUser)
	r.GET("/user", h.GetUserList)
	r.GET("/user/:user-id", h.GetUserByID)
	r.PUT("/user", h.UpdateUser)
	r.DELETE("/user/:user-id", h.DeleteUser)
	r.PUT("/user/reset-password", h.ResetPassword)
	r.POST("/user/send-message", h.SendMessageToUserEmail)

	r.POST("/integration", h.CreateIntegration)
	r.GET("/integration", h.GetIntegrationList)
	r.GET("/integration/:integration-id", h.GetIntegrationByID)
	r.DELETE("/integration/:integration-id", h.DeleteIntegration)
	r.GET("/integration/:integration-id/session", h.GetIntegrationSessions)
	r.POST("/integration/:integration-id/session", h.AddSessionToIntegration)
	r.GET("/integration/:integration-id/session/:session-id", h.GetIntegrationToken)

	r.POST("/user-relation", h.AddUserRelation)
	r.DELETE("/user-relation", h.RemoveUserRelation)

	r.POST("/upsert-user-info/:user-id", h.UpsertUserInfo)

	r.POST("/login", h.Login)
	r.DELETE("/logout", h.Logout)
	r.PUT("/refresh", h.RefreshToken)
	r.POST("/has-acess", h.HasAccess)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
