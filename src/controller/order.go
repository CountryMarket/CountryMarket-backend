package controller

import (
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/model"
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func OrderGenerateOrder(ctx *gin.Context) {
	req := param.ReqOrderGenerateOrder{}
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

	// 找到商品各自的商家
	productMap := make(map[int][]int)
	for _, v := range req.ProductIds {
		product, err := model.Get().ShopGetProduct(v)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, "cannot get product", err)
			return
		}
		if productMap[product.OwnerUserId] == nil {
			productMap[product.OwnerUserId] = make([]int, 0)
		}
		productMap[product.OwnerUserId] = append(productMap[product.OwnerUserId], v)
	}
	// map 转为切片
	var productSlice [][]int
	for _, v := range productMap {
		productSlice = append(productSlice, v)
	}
	// 生产订单
	err = model.Get().OrderGenerateOrder(productSlice, int(profile.ID), req.TransportationPrice, model.Address{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}, req.Message)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot generate product", err)
		return
	}
	response.Success(ctx, "ok")
}
func OrderGetUserOrder(ctx *gin.Context) {
	req := param.ReqOrderGetOrders{}
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

	ordersRaw, err := model.Get().OrderGetUserOrder(int(profile.ID), req.Length, req.From)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get orders", err)
		return
	}

	var orders []param.Order
	for _, v := range ordersRaw {
		// 将字符数组转为真数组，按照顺序排列 id 和 count
		arrayRaw := strings.Split(v.ProductAndCount, ",")
		var pcs []param.ProductAndCount
		for i := 0; i < len(arrayRaw); i += 2 {
			id, _ := strconv.Atoi(arrayRaw[i])
			count, _ := strconv.Atoi(arrayRaw[i+1])
			pcs = append(pcs, param.ProductAndCount{ProductId: id, Count: count})
		}

		orders = append(orders, param.Order{
			OwnerUserId:         v.OwnerUserId,
			OwnerShopUserId:     v.OwnerShopUserId,
			NowStatus:           v.NowStatus,
			TotalPrice:          v.TotalPrice,
			TransportationPrice: v.TransportationPrice,
			DiscountPrice:       v.DiscountPrice,
			ProductAndCount:     pcs,
			PersonName:          v.PersonName,
			PersonPhoneNumber:   v.PersonPhoneNumber,
			PersonAddress:       v.PersonAddress,
			OrderTime:           v.CreatedAt.Unix(),
			PayTime:             v.PayTime.Unix(),
			VerifyTime:          v.VerifyTime.Unix(),
			TrackingNumber:      v.TrackingNumber,
			Message:             v.Message,
		})
	}

	response.Success(ctx, param.ResOrderGetOrders{
		Orders: orders,
	})
}
func OrderGetOneOrder(ctx *gin.Context) {
	req := param.ReqOrderGetOneOrder{}
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

	order, err := model.Get().OrderGetOneOrder(req.OrderId, int(profile.ID))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get order", err)
		return
	}

	// 将字符数组转为真数组，按照顺序排列 id 和 count
	arrayRaw := strings.Split(order.ProductAndCount, ",")
	var pcs []param.ProductAndCount
	for i := 0; i < len(arrayRaw); i += 2 {
		id, _ := strconv.Atoi(arrayRaw[i])
		count, _ := strconv.Atoi(arrayRaw[i+1])
		pcs = append(pcs, param.ProductAndCount{ProductId: id, Count: count})
	}

	response.Success(ctx, param.Order{
		OwnerUserId:         order.OwnerUserId,
		OwnerShopUserId:     order.OwnerShopUserId,
		NowStatus:           order.NowStatus,
		TotalPrice:          order.TotalPrice,
		TransportationPrice: order.TransportationPrice,
		DiscountPrice:       order.DiscountPrice,
		ProductAndCount:     pcs,
		PersonName:          order.PersonName,
		PersonPhoneNumber:   order.PersonPhoneNumber,
		PersonAddress:       order.PersonAddress,
		OrderTime:           order.CreatedAt.Unix(),
		PayTime:             order.PayTime.Unix(),
		VerifyTime:          order.VerifyTime.Unix(),
		TrackingNumber:      order.TrackingNumber,
		Message:             order.Message,
	})
}
func OrderGetShopOrder(ctx *gin.Context) {
	req := param.ReqOrderGetOrders{}
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

	if (profile.Permission & 4) == 0 {
		response.Error(ctx, http.StatusForbidden, "not a seller", err)
		return
	}

	ordersRaw, err := model.Get().OrderGetShopOrder(int(profile.ID), req.Length, req.From)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get orders", err)
		return
	}

	var orders []param.Order
	for _, v := range ordersRaw {
		// 将字符数组转为真数组，按照顺序排列 id 和 count
		arrayRaw := strings.Split(v.ProductAndCount, ",")
		var pcs []param.ProductAndCount
		for i := 0; i < len(arrayRaw); i += 2 {
			id, _ := strconv.Atoi(arrayRaw[i])
			count, _ := strconv.Atoi(arrayRaw[i+1])
			pcs = append(pcs, param.ProductAndCount{ProductId: id, Count: count})
		}

		orders = append(orders, param.Order{
			OwnerUserId:         v.OwnerUserId,
			OwnerShopUserId:     v.OwnerShopUserId,
			NowStatus:           v.NowStatus,
			TotalPrice:          v.TotalPrice,
			TransportationPrice: v.TransportationPrice,
			DiscountPrice:       v.DiscountPrice,
			ProductAndCount:     pcs,
			PersonName:          v.PersonName,
			PersonPhoneNumber:   v.PersonPhoneNumber,
			PersonAddress:       v.PersonAddress,
			OrderTime:           v.CreatedAt.Unix(),
			PayTime:             v.PayTime.Unix(),
			VerifyTime:          v.VerifyTime.Unix(),
			TrackingNumber:      v.TrackingNumber,
			Message:             v.Message,
		})
	}

	response.Success(ctx, param.ResOrderGetOrders{
		Orders: orders,
	})
}
