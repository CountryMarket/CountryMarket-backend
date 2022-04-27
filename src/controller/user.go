package controller

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"log"
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

	log.Print(url)

	getString, err := util.HttpGet(url)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get openId", nil)
		return
	}

	getJson := getResult{}
	if err := json.Unmarshal([]byte(getString), &getJson); err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot unmarshal data", nil)
		return
	}

	log.Print(getJson)

	if getJson.Errcode != 0 {
		response.Error(ctx, http.StatusInternalServerError, getJson.Errmsg, nil)
		return
	} else {
		encryptedOpenId := util.AesEncryptCBC([]byte(getJson.Openid), []byte(getJson.SessionKey))
		response.Success(ctx, gin.H{
			"openid": hex.EncodeToString(encryptedOpenId),
		})
	}

}
