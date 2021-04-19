package handler

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
	a := gin.New()
	setupApiV1(a)

	return a
}

type ResponseFromHandler struct {
	Code  int
	Data  interface{}
	Error error
}

func setupApiV1(app *gin.Engine) {
	userHandler := NewUserRouter()
	apiGroup := app.Group("v1")

	// api group
	GET(apiGroup, "/hello", userHandler.Hello)

	// auth group
	authGroup := apiGroup.Group("auth")
	POST(authGroup, "/register/phone", userHandler.RegisterPhone)
	POST(authGroup, "/phone/login", userHandler.LoginPhone)
}

func GET(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "GET", relativePath, f)
}

func POST(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "POST", relativePath, f)
}

func PUT(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "PUT", relativePath, f)
}

func DELETE(group *gin.RouterGroup, relativePath string, f func(*gin.Context)) {
	route(group, "DELETE", relativePath, f)
}

func route(group *gin.RouterGroup, method string, relativePath string, f func(*gin.Context)) {
	switch method {
	case "POST":
		group.POST(relativePath, f)
	case "GET":
		group.GET(relativePath, f)
	case "PUT":
		group.PUT(relativePath, f)
	case "DELETE":
		group.DELETE(relativePath, f)
	}
}
