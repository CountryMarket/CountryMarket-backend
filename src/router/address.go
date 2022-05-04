package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitAddressRouter(g *gin.RouterGroup) {
	g.POST("/addAddress", middleware.JWTMiddleware(), controller.AddressAddAddress)
	g.POST("/modifyAddress", middleware.JWTMiddleware(), controller.AddressModifyAddress)
	g.POST("/deleteAddress", middleware.JWTMiddleware(), controller.AddressDeleteAddress)
	g.GET("/address", middleware.JWTMiddleware(), controller.AddressGetAddress)
	g.GET("/default", middleware.JWTMiddleware(), controller.AddressGetDefaultAddress)
	g.POST("/modifyDefault", middleware.JWTMiddleware(), controller.AddressModifyDefaultAddress)
}
