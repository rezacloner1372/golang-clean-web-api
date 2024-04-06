package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rezacloner1372/golang-clean-web-api/api/helper"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

/*
func (h *HealthHandler) HealthGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, i am Get handlers for Health",
	})
}
func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, i am Post handlers for Health",
	})
}
func (h *HealthHandler) HealthPostById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("Hello, i am Post handlers for Health with id: %s", id))
}
*/

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Working!", true, 0))
}
