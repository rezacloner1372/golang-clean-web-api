package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}
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