package param

type ReqSearch struct {
	Key string `form:"key" json:"key" binding:"required"`
}

type ResSearch struct {
	Products []ResShopGetProduct
}
