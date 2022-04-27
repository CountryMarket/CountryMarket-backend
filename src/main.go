package main

import (
	"github.com/CountryMarket/CountryMarket-backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	g := r.Group("/api/v1")
	router.InitRouter(g)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
