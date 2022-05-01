package controller

import (
	"encoding/json"
	"fmt"
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/model"
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type getLogin struct {
	Appid     string `json:"appid"`
	Secret    string `json:"secret"`
	JSCode    string `json:"js_code"`
	GrantType string `json:"grant_type"`
}
type getResult struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	Errcode    int    `json:"errcode"`
	Errmsg     string `json:"errmsg"`
}

func UserLogin(ctx *gin.Context) {
	req := param.ReqUserLogin{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}

	// 获取 OpenId
	p := getLogin{
		os.Getenv("APPID"),
		os.Getenv("APPSECRET"),
		req.Code,
		"authorization_code",
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=%s",
		p.Appid, p.Secret, p.JSCode, p.GrantType)

	getString, err := util.HttpGet(url)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get openId", err)
		return
	}

	getJson := getResult{}
	if err := json.Unmarshal([]byte(getString), &getJson); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot unmarshal data", err)
		return
	}

	if getJson.Errcode != 0 {
		response.Error(ctx, http.StatusInternalServerError, getJson.Errmsg, err)
		return
	} else {
		encryptedOpenId, _, err := util.GenerateJWTToken(getJson.Openid, getJson.SessionKey)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, "cannot generate token", err)
			return
		}

		// 如果用户第一次登录，则将用户数据插入数据库，否则什么都不做
		err = model.Get().UserRegisterOrDoNothing(getJson.Openid, req.NickName, req.AvatarUrl)
		if err != nil {
			response.Error(ctx, http.StatusInternalServerError, "cannot check user database", err)
			return
		}

		response.Success(ctx, gin.H{
			"token": encryptedOpenId,
		})
	}

}
func UserValidate(ctx *gin.Context) {
	util.GetClaimsFromJWT(ctx)
	response.Success(ctx, "ok")
}
func UserGetProfile(ctx *gin.Context) {
	openid, _ := util.GetClaimsFromJWT(ctx)
	user, err := model.Get().UserGetProfile(openid)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get user database", err)
		return
	}
	response.Success(ctx, param.ResUserGetProfile{
		NickName:    user.NickName,
		AvatarUrl:   user.AvatarUrl,
		PhoneNumber: user.PhoneNumber,
		Permission:  user.Permission,
	})
}
