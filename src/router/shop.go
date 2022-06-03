package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitShopRouter(g *gin.RouterGroup) {
	g.POST("/addProduct", middleware.JWTMiddleware(), controller.ShopAddProduct)
	g.POST("/updateProduct", middleware.JWTMiddleware(), controller.ShopUpdateProduct)

	g.GET("/product", controller.ShopGetProduct)
	g.GET("/ownerProducts", middleware.JWTMiddleware(), controller.ShopGetOwnerProducts)
	g.POST("/dropProduct", middleware.JWTMiddleware(), controller.ShopDropProduct)
	g.POST("/putProduct", middleware.JWTMiddleware(), controller.ShopPutProduct)
}
