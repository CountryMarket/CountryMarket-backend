package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	g.GET("/code", controller.UserLogin)
	g.GET("/test", middleware.JWTMiddleware(), controller.UserJWTTest) // for test
}
