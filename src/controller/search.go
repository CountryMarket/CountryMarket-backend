package controller

import (
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/model"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Search(ctx *gin.Context) {
	req := param.ReqSearch{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	results, err := model.Get().Search(req.Key)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot search", err)
		return
	}
	var products []param.ResShopGetProduct
	for _, v := range results {
		products = append(products, param.ResShopGetProduct{
			Id:          int(v.ID),
			Price:       v.Price,
			Title:       v.Title,
			Description: v.Description,
			OwnerUserId: v.OwnerUserId,
			Stock:       v.Stock,
			IsDrop:      v.IsDrop,
		})
	}
	response.Success(ctx, param.ResSearch{
		Products: products,
	})
}
