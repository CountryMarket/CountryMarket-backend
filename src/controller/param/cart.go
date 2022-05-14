package param

type ReqCartGetUserProducts struct {
	From   int `form:"from" json:"from"`
	Length int `form:"length" json:"length" binding:"required"`
}
type ReqCartProductId struct {
	ProductId int `form:"productId" json:"productId" binding:"required"`
}
type ReqCartModifyProduct struct {
	ProductId   int `form:"productId" json:"productId" binding:"required"`
	ModifyCount int `form:"modifyCount" json:"modifyCount"`
}
type ReqCartGetCart struct {
	ProductIds []int `form:"product_ids" json:"product_ids"`
}

type ResCartGetInCart struct {
	InCart bool
	Count  int
}
type CartItem struct {
	Id          int
	Count       int
	Price       float64
	Title       string
	Description string
	OwnerUserId int
	Stock       int
	IsDrop      bool
}
type ResCartGetUserProducts struct {
	Products []CartItem
}
