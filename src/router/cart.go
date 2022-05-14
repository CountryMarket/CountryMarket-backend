package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitCartRouter(g *gin.RouterGroup) {
	g.GET("/userProducts", middleware.JWTMiddleware(), controller.CartGetUserProducts)
	g.GET("/inCart", middleware.JWTMiddleware(), controller.CartGetInCart)
	g.POST("/getCart", middleware.JWTMiddleware(), controller.CartGetCart)

	g.POST("/addProduct", middleware.JWTMiddleware(), controller.CartAddProduct)
	g.POST("/reduceProduct", middleware.JWTMiddleware(), controller.CartReduceProduct)
	g.POST("/modifyProduct", middleware.JWTMiddleware(), controller.CartModifyProduct)
}
