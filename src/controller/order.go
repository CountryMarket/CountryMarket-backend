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
	"time"
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

	/*// 找到商品各自的商家
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
	}*/
	var productSlice [][]int
	productSlice = append(productSlice, req.ProductIds)
	// 生产订单
	resId, err := model.Get().OrderGenerateOrder(productSlice, int(profile.ID), profile.PhoneNumber, req.TransportationPrice, model.Address{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
	}, req.Message)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot generate product", err)
		return
	}
	response.Success(ctx, gin.H{
		"order_id": resId,
	})
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

	ordersRaw, err := model.Get().OrderGetUserOrder(int(profile.ID), req.Length, req.From, req.Status)
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

			product, err := model.Get().ShopGetProduct(id)
			if err != nil {
				return
			}

			pcs = append(pcs, param.ProductAndCount{Products: param.ProductItem{
				Id:            int(product.ID),
				OwnerUserId:   product.OwnerUserId,
				Price:         product.Price,
				Title:         product.Title,
				Description:   product.Description,
				PictureNumber: product.PictureNumber,
				Stock:         product.Stock,
				IsDrop:        product.IsDrop,
			}, Count: count})
		}

		orders = append(orders, param.Order{
			OrderId:             int(v.ID),
			OwnerUserId:         v.OwnerUserId,
			OwnerShopUserId:     v.OwnerShopUserId,
			UserPhoneNumber:     v.UserPhoneNumber,
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
			TrackingNumber:      strings.Split(v.TrackingNumber, " "),
			Message:             v.Message,
			ShopMessage:         v.ShopMessage,
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

		product, err := model.Get().ShopGetProduct(id)
		if err != nil {
			return
		}

		pcs = append(pcs, param.ProductAndCount{Products: param.ProductItem{
			Id:            int(product.ID),
			OwnerUserId:   product.OwnerUserId,
			Price:         product.Price,
			Title:         product.Title,
			Description:   product.Description,
			PictureNumber: product.PictureNumber,
			Stock:         product.Stock,
			IsDrop:        product.IsDrop,
		}, Count: count})
	}

	response.Success(ctx, param.Order{
		OrderId:             int(order.ID),
		OwnerUserId:         order.OwnerUserId,
		OwnerShopUserId:     order.OwnerShopUserId,
		UserPhoneNumber:     order.UserPhoneNumber,
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
		TrackingNumber:      strings.Split(order.TrackingNumber, " "),
		Message:             order.Message,
		ShopMessage:         order.ShopMessage,
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

	ordersRaw, err := model.Get().OrderGetShopOrder(int(profile.ID), req.Length, req.From, req.Status)
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

			product, err := model.Get().ShopGetProduct(id)
			if err != nil {
				return
			}

			pcs = append(pcs, param.ProductAndCount{Products: param.ProductItem{
				Id:            int(product.ID),
				OwnerUserId:   product.OwnerUserId,
				Price:         product.Price,
				Title:         product.Title,
				Description:   product.Description,
				PictureNumber: product.PictureNumber,
				Stock:         product.Stock,
				IsDrop:        product.IsDrop,
			}, Count: count})
		}

		orders = append(orders, param.Order{
			OrderId:             int(v.ID),
			OwnerUserId:         v.OwnerUserId,
			OwnerShopUserId:     v.OwnerShopUserId,
			UserPhoneNumber:     v.UserPhoneNumber,
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
			TrackingNumber:      strings.Split(v.TrackingNumber, " "),
			Message:             v.Message,
			ShopMessage:         v.ShopMessage,
		})
	}

	response.Success(ctx, param.ResOrderGetOrders{
		Orders: orders,
	})
}
func OrderDeleteOrder(ctx *gin.Context) {
	req := param.ReqOrderGetOneOrder{}
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

	err = model.Get().OrderDeleteOrder(int(profile.ID), req.OrderId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot delete order", err)
		return
	}
	response.Success(ctx, "ok")
}
func OrderChangeStatus(ctx *gin.Context) {
	req := param.ReqOrderChangeStatus{}
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

	var pt, vt time.Time
	pt = time.Unix(int64(req.PayTime), 0)
	vt = time.Unix(int64(req.VerifyTime), 0)
	err = model.Get().OrderChangeStatus(int(profile.ID), req.OrderId, req.Status, pt, vt, req.ShopMessage)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot change status", err)
		return
	}
	response.Success(ctx, "ok")
}
func OrderAddTrackingNumber(ctx *gin.Context) {
	req := param.ReqOrderAddTrackingNumber{}
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

	var TNs string
	for i, v := range req.TrackingNumber {
		TNs += v
		if i != len(req.TrackingNumber) {
			TNs += " "
		}
	}

	err = model.Get().OrderAddTrackingNumber(int(profile.ID), req.OrderId, TNs)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot add", err)
		return
	}
	response.Success(ctx, "ok")
}
