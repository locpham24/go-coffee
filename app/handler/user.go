package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/locpham24/go-coffee/app/entity"
	"github.com/locpham24/go-coffee/app/form"
	"github.com/locpham24/go-coffee/app/response"
	"github.com/locpham24/go-coffee/infra"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserRouter interface {
	Hello(c *gin.Context)
	RegisterPhone(c *gin.Context)
	LoginPhone(c *gin.Context)
	Get(c *gin.Context)
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
		infra.GetLogging().Log(logrus.ErrorLevel, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userEntity := entity.UserEntity{}
	user, err := userEntity.Create(input)
	if err != nil {
		infra.GetLogging().Log(logrus.ErrorLevel, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userView, err := response.PopulateUserView(user)
	if err != nil {
		infra.GetLogging().Log(logrus.ErrorLevel, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(200, userView)
}

func (h *UserHandler) LoginPhone(c *gin.Context) {
	var input form.LoginPhoneNumber
	if err := c.Bind(&input); err != nil {
		infra.GetLogging().Log(logrus.ErrorLevel, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userEntity := entity.UserEntity{}
	tokens, err := userEntity.LoginPhone(input)
	if err != nil {
		infra.GetLogging().Log(logrus.ErrorLevel, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(200, tokens)
}

func (h *UserHandler) Get(c *gin.Context) {
	userIdGinKey := c.MustGet(UserGinKey)
	userIdStr := fmt.Sprintf("%v", userIdGinKey)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		infra.GetLogging().Log(logrus.ErrorLevel, err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userEntity := entity.UserEntity{}
	user, err := userEntity.GetById(userId)
	if err != nil {
		infra.GetLogging().Log(logrus.ErrorLevel, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userView, err := response.PopulateUserView(user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(200, userView)
}
