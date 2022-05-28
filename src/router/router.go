package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
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
	rOrder := g.Group("/order")
	InitOrderRouter(rOrder)
	rComment := g.Group("/comment")
	InitCommentRouter(rComment)

	g.POST("/search", controller.Search)
}
