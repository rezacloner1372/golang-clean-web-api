package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rezacloner1372/golang-clean-web-api/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.HealthGet)
	r.POST("/", handler.HealthPost)
	r.POST("/:id", handler.HealthPostById)
}
