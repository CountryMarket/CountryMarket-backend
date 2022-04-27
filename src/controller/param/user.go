package param

type ReqUserLogin struct {
	Code string `form:"code" json:"code" binding:"required"`
}
