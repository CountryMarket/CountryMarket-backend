package main

import (
	"github.com/CountryMarket/CountryMarket-backend/config"
	"github.com/CountryMarket/CountryMarket-backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	g := r.Group(config.C.App.Root)
	router.InitRouter(g)

	err := r.Run(config.C.App.Addr)
	if err != nil {
		panic(err)
	}
}
