package router

import (
	"github.com/CountryMarket/CountryMarket-backend/controller"
	"github.com/CountryMarket/CountryMarket-backend/middleware"
	"github.com/gin-gonic/gin"
)

func InitCommentRouter(g *gin.RouterGroup) {
	g.POST("/add", middleware.JWTMiddleware(), controller.CommentAddComment)
	g.GET("/product", controller.CommentGetProductComment)
	g.POST("/delete", middleware.JWTMiddleware(), controller.CommentDeleteComment)
}
