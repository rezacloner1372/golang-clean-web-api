package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type header struct {
	UserId  string
	Browser string
}

type personData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=10"`
	LastName     string `json:"last_name" binding:"required,alpha,min=6,max=20"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}
func (h *TestHandler) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Test handler",
	})
}
func (h *TestHandler) User(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from User handler",
	})
}
func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from UserById handler",
		"id":      id,
	})
}
func (h *TestHandler) UserByUsername(c *gin.Context) {
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"message":  "Hello from UserByUsername handler",
		"username": username,
	})
}
func (h *TestHandler) Account(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Account handler",
		"id":      id,
	})
}
func (h *TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from AddUser handler",
	})
}

// find the user by id at header
func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	userId := c.GetHeader("UserId")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from HeaderBinder1 handler",
		"userId":  userId,
	})
}

// find the header struct
func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from HeaderBinder1 handler",
		"header":  header,
	})
}

func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from QueryBinder1 handler",
		"id":      id,
		"name":    name,
	})
}
func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from QueryBinder2 handler",
		"ids":     ids,
		"name":    name,
	})
}

func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from UriBinder handler",
		"ids":     id,
		"name":    name,
	})
}

func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from BodyBinder handler",
		"p":       p,
	})
}

// bind data from form
func (h *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	c.Bind(&p)
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from FormBinder handler",
		"person":  p,
	})
}

// bind data from file
func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from FileBinder handler",
		"file":    file.Filename,
	})
}
