package controller

import (
	"github.com/CountryMarket/CountryMarket-backend/controller/param"
	"github.com/CountryMarket/CountryMarket-backend/model"
	"github.com/CountryMarket/CountryMarket-backend/util"
	"github.com/CountryMarket/CountryMarket-backend/util/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentAddComment(ctx *gin.Context) {
	req := param.ReqAddComment{}
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

	var comments []model.UserComment
	for _, v := range req.Comments {
		comments = append(comments, model.UserComment{
			UserId:    int(profile.ID),
			ProductId: v.ProductId,
			Comment:   v.Comment,
		})
	}

	err = model.Get().CommentAddComments(int(profile.ID), comments)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot add comment", err)
		return
	}

	response.Success(ctx, "ok")
}
func CommentGetProductComment(ctx *gin.Context) {
	req := param.ReqCommentGetProductComment{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, "bad request", err)
		return
	}

	comments, err := model.Get().CommentGetProductComment(req.ProductId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot get comments", err)
		return
	}

	var resComments []param.ResComment
	for _, v := range comments {
		resComments = append(resComments, param.ResComment{
			UserId:  v.UserId,
			Comment: v.Comment,
		})
	}

	response.Success(ctx, param.ResCommentGetProductComment{
		Comments: resComments,
	})
}
func CommentDeleteComment(ctx *gin.Context) {
	req := param.ReqDeleteComment{}
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

	err = model.Get().CommentDeleteComment(req.CommentId)
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, "cannot delete comment", err)
		return
	}

	response.Success(ctx, "ok")
}
