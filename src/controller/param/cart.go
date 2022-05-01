package param

type ReqCartGetUserProducts struct {
	From   int `form:"from" json:"from"`
	Length int `form:"length" json:"length" binding:"required"`
}
type ReqCartProductId struct {
	ProductId int `form:"productId" json:"product_id" binding:"required"`
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
}
type ResCartGetUserProducts struct {
	Products []CartItem
}
