package controller

import (
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/model"
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShopAddProduct(ctx *gin.Context) {
	req := param.ReqShopAddProduct{}
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

	if (profile.Permission & 2) == 0 { // 不是商户账号
		response.Error(ctx, http.StatusForbidden, "not a seller", nil)
		return
	}

	err = model.Get().ShopAddProduct(model.Product{
		OwnerUserId:   int(profile.ID),
		Price:         req.Price,
		Title:         req.Title,
		Description:   req.Description,
		PictureNumber: req.PictureNumber,
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot add product", err)
		return
	}

	response.Success(ctx, "ok")
}
func ShopUpdateProduct(ctx *gin.Context) {
	req := param.ReqShopUpdateProduct{}
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

	if (profile.Permission & 2) == 0 { // 不是商户账号
		response.Error(ctx, http.StatusForbidden, "not a seller", nil)
		return
	}

	err = model.Get().ShopUpdateProduct(model.Product{
		OwnerUserId:   int(profile.ID),
		Price:         req.Price,
		Title:         req.Title,
		Description:   req.Description,
		PictureNumber: req.PictureNumber,
	}, req.Id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot update product", err)
		return
	}

	response.Success(ctx, "ok")
}
func ShopGetProduct(ctx *gin.Context) {
	req := param.ReqShopGetProduct{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	product, err := model.Get().ShopGetProduct(req.Id)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get product", err)
		return
	}
	response.Success(ctx, param.ResShopGetProduct{
		Id:            int(product.ID),
		OwnerUserId:   product.OwnerUserId, // user 表中的 id 列
		Price:         product.Price,
		Title:         product.Title,
		Description:   product.Description,
		PictureNumber: product.PictureNumber,
	})
}
func ShopGetOwnerProducts(ctx *gin.Context) {
	req := param.ReqShopGetOwnerProducts{}
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

	if (profile.Permission & 2) == 0 { // 不是商户账号
		response.Error(ctx, http.StatusForbidden, "not a seller", nil)
		return
	}

	products, err := model.Get().ShopGetOwnerProducts(int(profile.ID), req.From, req.Length)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get products", err)
		return
	}

	var resProducts []param.ResShopGetProduct
	for _, v := range products {
		resProducts = append(resProducts, param.ResShopGetProduct{
			Id:            int(v.ID),
			OwnerUserId:   v.OwnerUserId, // user 表中的 id 列
			Price:         v.Price,
			Title:         v.Title,
			Description:   v.Description,
			PictureNumber: v.PictureNumber,
		})
	}

	response.Success(ctx, param.ResShopGetOwnerProducts{Products: resProducts})
}
