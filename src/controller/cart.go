package controller

import (
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/model"
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CartGetUserProducts(ctx *gin.Context) {
	req := param.ReqCartGetUserProducts{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	openid, _ := util.GetClaimsFromJWT(ctx)

	profile, err := model.Get().UserGetProfile(openid)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get profile", err)
		return
	}

	cartAndProducts, err := model.Get().CartGetUserProducts(int(profile.ID), req.From, req.Length)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get products", err)
		return
	}

	var resCarts []param.CartItem

	for _, v := range cartAndProducts {
		resCarts = append(resCarts, param.CartItem{
			Id:          v.ProductId,
			Count:       v.ProductCount,
			Price:       v.Price,
			Title:       v.Title,
			Description: v.Description,
			OwnerUserId: v.OwnerUserId,
			Stock:       v.Stock,
			IsDrop:      v.IsDrop,
		})
	}

	response.Success(ctx, param.ResCartGetUserProducts{Products: resCarts})
}
func CartGetInCart(ctx *gin.Context) {
	req := param.ReqCartProductId{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	openid, _ := util.GetClaimsFromJWT(ctx)

	profile, err := model.Get().UserGetProfile(openid)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get profile", err)
		return
	}

	count, err := model.Get().CartGetInCart(int(profile.ID), req.ProductId)
	if err != nil && err != gorm.ErrRecordNotFound {
		response.Error(ctx, http.StatusInternalServerError, "cannot get cart records", err)
		return
	}
	if err == gorm.ErrRecordNotFound {
		response.Success(ctx, param.ResCartGetInCart{
			InCart: false,
			Count:  0,
		})
		return
	}
	response.Success(ctx, param.ResCartGetInCart{
		InCart: true,
		Count:  int(count),
	})
}
func CartAddProduct(ctx *gin.Context) {
	req := param.ReqCartProductId{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	openid, _ := util.GetClaimsFromJWT(ctx)

	profile, err := model.Get().UserGetProfile(openid)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get profile", err)
		return
	}

	count, err := model.Get().CartAddProduct(int(profile.ID), req.ProductId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot add product or get count", err)
		return
	}

	response.Success(ctx, gin.H{
		"count": count,
	})
}
func CartReduceProduct(ctx *gin.Context) {
	req := param.ReqCartProductId{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	openid, _ := util.GetClaimsFromJWT(ctx)

	profile, err := model.Get().UserGetProfile(openid)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get profile", err)
		return
	}

	count, err := model.Get().CartReduceProduct(int(profile.ID), req.ProductId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot reduce product or get count", err)
		return
	}

	response.Success(ctx, gin.H{
		"count": count,
	})
}
func CartModifyProduct(ctx *gin.Context) {
	req := param.ReqCartModifyProduct{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	openid, _ := util.GetClaimsFromJWT(ctx)

	profile, err := model.Get().UserGetProfile(openid)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get profile", err)
		return
	}

	count, err := model.Get().CartModifyProduct(int(profile.ID), req.ProductId, req.ModifyCount)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot add product or get count", err)
		return
	}

	response.Success(ctx, gin.H{
		"count": count,
	})
}
func CartGetCart(ctx *gin.Context) {
	req := param.ReqCartGetCart{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	openid, _ := util.GetClaimsFromJWT(ctx)

	profile, err := model.Get().UserGetProfile(openid)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get profile", err)
		return
	}

	cartAndProducts, err := model.Get().CartGetCart(int(profile.ID), req.ProductIds)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get product", err)
		return
	}

	var resCarts []param.CartItem

	for _, v := range cartAndProducts {
		resCarts = append(resCarts, param.CartItem{
			Id:          v.ProductId,
			Count:       v.ProductCount,
			Price:       v.Price,
			Title:       v.Title,
			Description: v.Description,
			OwnerUserId: v.OwnerUserId,
			Stock:       v.Stock,
			IsDrop:      v.IsDrop,
		})
	}

	response.Success(ctx, param.ResCartGetUserProducts{Products: resCarts})
}
