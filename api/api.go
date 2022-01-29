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
	r.GET("/sphere/:sphere_id", h.GetSphereByID)
	r.PUT("/sphere", h.UpdateSphere)
	r.DELETE("/sphere/:sphere_id", h.DeleteSphere)

	r.POST("/position", h.CreatePosition)
	r.GET("/position", h.GetPositionList)
	r.GET("/position/:position_id", h.GetPositionByID)
	r.PUT("/position", h.UpdatePosition)
	r.DELETE("/position/:position_id", h.DeletePosition)
	r.POST("/position/level", h.AddLevelItem)
	r.PUT("/position/level", h.UpdateLevelItem)
	r.DELETE("/position/:position_id/level/:level_id", h.RemoveLevelItem)

	// CLIENT SERVICE
	r.POST("/client_platform", h.CreateClientPlatform)
	r.GET("/client_platform", h.GetClientPlatformList)
	r.GET("/client_platform/:client_platform_id", h.GetClientPlatformByID)
	r.PUT("/client_platform", h.UpdateClientPlatform)
	r.DELETE("/client_platform/:client_platform_id", h.DeleteClientPlatform)

	r.POST("/client_type", h.CreateClientType)
	r.GET("/client_type", h.GetClientTypeList)
	r.GET("/client_type/:client_type_id", h.GetClientTypeByID)
	r.PUT("/client_type", h.UpdateClientType)
	r.DELETE("/client_type/:client_type_id", h.DeleteClientType)

	r.POST("/client", h.AddClient)
	r.PUT("/client", h.UpdateClient)
	r.DELETE("/client/:client_platform_id/:client_type_id", h.RemoveClient)

	r.POST("/relation", h.AddRelation)
	r.PUT("/relation", h.UpdateRelation)
	r.DELETE("/relation/:relation_id", h.RemoveRelation)

	r.POST("/user_info_field", h.AddUserInfoField)
	r.PUT("/user_info_field", h.UpdateUserInfoField)
	r.DELETE("/user_info_field/:user_info_field_id", h.RemoveUserInfoField)

	// PERMISSION SERVICE
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
	r.DELETE("/role-permission", h.RemoveRolePermission)

	r.POST("/user", h.CreateUser)
	r.GET("/user", h.GetUserList)
	r.GET("/user/:user-id", h.GetUserByID)
	r.PUT("/user", h.UpdateUser)
	r.DELETE("/user/:user-id", h.DeleteUser)

	r.POST("/user-relation", h.AddUserRelation)
	r.DELETE("/user-relation", h.RemoveUserRelation)

	r.POST("/upsert-user-info", h.UpsertUserInfo)

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
