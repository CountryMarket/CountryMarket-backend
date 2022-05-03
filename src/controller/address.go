package controller

import (
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/model"
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddressAddAddress(ctx *gin.Context) {
	req := param.ReqAddressAddAddress{}
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

	err = model.Get().AddressAddAddress(model.Address{
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		OwnerUserId: int(profile.ID),
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot add address", err)
		return
	}
	response.Success(ctx, "ok")
}
func AddressModifyAddress(ctx *gin.Context) {
	req := param.ReqAddressModifyAddress{}
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

	address, err := model.Get().AddressGetOneAddress(req.AddressId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get address", err)
		return
	}
	if address.OwnerUserId != (int)(profile.ID) {
		response.Error(ctx, http.StatusForbidden, "not your address", err)
		return
	}

	err = model.Get().AddressModifyAddress(req.AddressId, model.Address{
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		OwnerUserId: int(profile.ID),
	})
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot modify address", err)
		return
	}
	response.Success(ctx, "ok")
}
func AddressDeleteAddress(ctx *gin.Context) {
	req := param.ReqAddressDeleteAddress{}
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

	address, err := model.Get().AddressGetOneAddress(req.AddressId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get address", err)
		return
	}
	if address.OwnerUserId != (int)(profile.ID) {
		response.Error(ctx, http.StatusForbidden, "not your address", err)
		return
	}

	err = model.Get().AddressDeleteAddress(req.AddressId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot delete address", err)
		return
	}
	response.Success(ctx, "ok")
}
func AddressGetAddress(ctx *gin.Context) {
	openid, _ := util.GetClaimsFromJWT(ctx)

	profile, err := model.Get().UserGetProfile(openid)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get profile", err)
		return
	}

	addresses, err := model.Get().AddressGetAddress(int(profile.ID))
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get address", err)
		return
	}

	var result []param.AddressItem

	for _, v := range addresses {
		result = append(result, param.AddressItem{
			Name:        v.Name,
			Address:     v.Address,
			PhoneNumber: v.PhoneNumber,
			AddressId:   int(v.ID),
		})
	}

	response.Success(ctx, param.ResAddressGetAddress{
		Address: result,
	})
}
