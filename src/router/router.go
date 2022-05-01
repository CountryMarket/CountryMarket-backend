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
}
