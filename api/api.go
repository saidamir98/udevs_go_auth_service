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
