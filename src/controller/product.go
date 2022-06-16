package controller

import (
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/model"
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func arrayToString(intArray []int) string {
	var result string
	for i, v := range intArray {
		result += strconv.Itoa(v)
		if i != len(intArray)-1 {
			result += ", "
		}
	}
	return result
}

func ProductGetTabList(ctx *gin.Context) {
	ids, names, err := model.Get().ProductGetTabList()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get tab list", err)
		return
	}
	var tabs []param.TabItem
	for i := range ids {
		tabs = append(tabs, param.TabItem{
			Id:   ids[i],
			Name: names[i],
		})
	}
	response.Success(ctx, param.ResProductGetTabList{
		Tabs: tabs,
	})
}
func ProductGetTabProducts(ctx *gin.Context) {
	req := param.ReqProductGetTabProducts{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}

	products, err := model.Get().ProductGetTabProducts(req.TabId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get products", err)
		return
	}

	var result []param.TabProductsItem
	for _, v := range products {
		result = append(result, param.TabProductsItem{
			Id:          int(v.ID),
			Price:       v.Price,
			Title:       v.Title,
			Description: v.Description,
			OwnerUserId: v.OwnerUserId,
			Stock:       v.Stock,
			IsDrop:      v.IsDrop,
		})
	}

	response.Success(ctx, param.ResProductGetTabProducts{
		Products: result,
	})
}
func ProductModifyTab(ctx *gin.Context) {
	req := param.ReqProductModifyTab{}
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

	if (profile.Permission & 2) == 0 { // 不是 root
		response.Error(ctx, http.StatusForbidden, "not a shop", nil)
		return
	}

	err = model.Get().ProductModifyTabProducts(req.TabId, model.ProductTab{
		Name:     req.Name,
		Products: arrayToString(req.Products),
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot modify tab", err)
		return
	}
	response.Success(ctx, "ok")
}
func ProductAddTab(ctx *gin.Context) {
	req := param.ReqProductAddTab{}
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

	if (profile.Permission & 4) == 0 { // 不是 root
		response.Error(ctx, http.StatusForbidden, "not a root", nil)
		return
	}

	err = model.Get().ProductAddTabProducts(model.ProductTab{
		Name:     req.Name,
		Products: arrayToString(req.Products),
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot add tab", err)
		return
	}
	response.Success(ctx, "ok")
}
func ProductDeleteTab(ctx *gin.Context) {
	req := param.ReqProductDeleteTab{}
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

	if (profile.Permission & 4) == 0 { // 不是 root
		response.Error(ctx, http.StatusForbidden, "not a root", nil)
		return
	}

	err = model.Get().ProductDeleteTabProducts(req.TabId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot delete tab", err)
		return
	}
	response.Success(ctx, "ok")
}
func ProductGetHomeTab(ctx *gin.Context) {
	req := param.ReqProductGetHomeTab{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}
	tab, err := model.Get().ProductGetHomeTab(req.From, req.Length)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get home tab", err)
		return
	}
	var products []param.TabProductsItem
	for _, v := range tab {
		products = append(products, param.TabProductsItem{
			Id:          int(v.ID),
			Price:       v.Price,
			Title:       v.Title,
			Description: v.Description,
			OwnerUserId: v.OwnerUserId,
			Stock:       v.Stock,
			IsDrop:      v.IsDrop,
		})
	}
	/*for i := 0; i < 5; i++ {
		for j := i; j < len(tab); j += 5 {
			v := tab[j]
			products = append(products, param.TabProductsItem{
				Id:          int(v.ID),
				Price:       v.Price,
				Title:       v.Title,
				Description: v.Description,
				OwnerUserId: v.OwnerUserId,
				Stock:       v.Stock,
				IsDrop:      v.IsDrop,
			})
		}
	}*/
	response.Success(ctx, param.ResProductGetHomeTab{Products: products})
}
