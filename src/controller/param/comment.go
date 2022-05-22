package param

type ReqAddComment struct {
	Comments []ReqComment `form:"comments" json:"comments" binding:"required"`
}
type ReqComment struct {
	ProductId int    `form:"product_id" json:"product_id" binding:"required"`
	Comment   string `form:"comment" json:"comment" binding:"required"`
}
type ReqCommentGetProductComment struct {
	ProductId int `form:"product_id" json:"product_id" binding:"required"`
}
type ReqDeleteComment struct {
	CommentId int `form:"comment_id" json:"comment_id" binding:"required"`
}

type ResCommentGetProductComment struct {
	Comments []ResComment `json:"comments"`
}
type ResComment struct {
	UserId  int    `json:"user_id"`
	Comment string `json:"comment"`
}
