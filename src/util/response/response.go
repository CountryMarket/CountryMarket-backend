package response

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
	Info    string      `json:"info"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Data:    data,
		Success: true,
		Info:    "",
	})
}

func Error(ctx *gin.Context, status int, info string, err error) {
	if err != nil {
		log.Print(err)
	}
	ctx.JSON(status, Response{
		Data:    nil,
		Success: false,
		Info:    info,
	})
}
