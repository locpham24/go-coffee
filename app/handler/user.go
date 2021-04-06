package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/locpham24/go-coffee/app/entity"
	"github.com/locpham24/go-coffee/app/form"
	"github.com/locpham24/go-coffee/app/response"
)

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
	var input form.RegisterPhoneNumber
	if err := c.Bind(&input); err != nil {
		return
	}

	userEntity := entity.UserEntity{}
	user, err := userEntity.Create(input)
	if err != nil {
		return
	}

	userView, err := response.PopulateUserView(user)
	if err != nil {
		return
	}
	c.JSON(200, userView)
}
