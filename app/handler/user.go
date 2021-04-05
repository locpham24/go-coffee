package handler

import "github.com/gin-gonic/gin"

type UserRouter interface {
	Hello(c *gin.Context)
	RegisterPhone(c *gin.Context)
}

func NewUserRouter() UserRouter {
	return &UserHandler{}
}

type UserHandler struct {
}

func (h *UserHandler) Hello(c *gin.Context) {
	c.JSON(200, "Hello")
}

func (h *UserHandler) RegisterPhone(c *gin.Context) {

}
