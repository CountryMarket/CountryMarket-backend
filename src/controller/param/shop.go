package param

type ReqShopAddProduct struct {
	Price               float64 `form:"price" json:"price" binding:"required"`
	Title               string  `form:"title" json:"title" binding:"required"`
	Description         string  `form:"description" json:"description" binding:"required"`
	PictureNumber       int     `form:"pictureNumber" json:"pictureNumber"`
	Stock               int     `form:"stock" json:"stock"`
	Detail              string  `form:"detail" json:"detail" binding:"required"`
	DetailPictureNumber int     `form:"detailPictureNumber" json:"detailPictureNumber"`
}
type ReqShopUpdateProduct struct {
	Price               float64 `form:"price" json:"price" binding:"required"`
	Title               string  `form:"title" json:"title" binding:"required"`
	Description         string  `form:"description" json:"description" binding:"required"`
	PictureNumber       int     `form:"pictureNumber" json:"pictureNumber"`
	Stock               int     `form:"stock" json:"stock"`
	Detail              string  `form:"detail" json:"detail"`
	DetailPictureNumber int     `form:"detailPictureNumber" json:"detailPictureNumber"`
	Id                  int     `form:"id" json:"id" binding:"required"`
}
type ReqShopGetProduct struct {
	Id int `form:"id" json:"id" binding:"required"`
}
type ReqShopGetOwnerProducts struct {
	From   int `form:"from" json:"from"`
	Length int `form:"length" json:"length" binding:"required"`
}

type ResShopGetProduct struct {
	Id                  int
	OwnerUserId         int // user 表中的 id 列
	Price               float64
	Title               string
	Description         string
	PictureNumber       int
	Stock               int
	Detail              string
	DetailPictureNumber int
	IsDrop              bool
}
type ResShopGetOwnerProducts struct {
	Products []ResShopGetProduct
}
