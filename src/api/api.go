package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rezacloner1372/golang-clean-web-api/api/middlewares"
	"github.com/rezacloner1372/golang-clean-web-api/api/routers"
	"github.com/rezacloner1372/golang-clean-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()

	// val, ok := binding.Validator.Engine().(*validator.Validate)
	// if ok {
	// 	val.RegisterValidation("mobile", validation.IranianMobileNumberValidator)
	// }

	r.Use(gin.Logger(), gin.Recovery() /*middlewares.TestMiddleware()*/, middlewares.LimitByRequest())
	api := r.Group("/api")

	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)
	}
	v2 := api.Group("/v2/")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}
