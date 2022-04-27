package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(g *gin.RouterGroup) {
	g.GET("/code", controller.UserLogin)
}
