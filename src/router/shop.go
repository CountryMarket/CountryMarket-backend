package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitShopRouter(g *gin.RouterGroup) {
	g.POST("/addProduct", middleware.JWTMiddleware(), controller.ShopAddProduct)
	g.PUT("/updateProduct", middleware.JWTMiddleware(), controller.ShopUpdateProduct)

	g.GET("/product", controller.ShopGetProduct)
	g.GET("/ownerProducts", middleware.JWTMiddleware(), controller.ShopGetOwnerProducts)
}