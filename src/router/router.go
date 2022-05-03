package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.RouterGroup) {
	rUser := g.Group("/user")
	InitUserRouter(rUser)
	rShop := g.Group("/shop")
	InitShopRouter(rShop)
	rCart := g.Group("/cart")
	InitCartRouter(rCart)
	rAddress := g.Group("/address")
	InitAddressRouter(rAddress)
	rProduct := g.Group("/product")
	InitProductRouter(rProduct)
}
