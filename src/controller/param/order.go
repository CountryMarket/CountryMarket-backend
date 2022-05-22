package param

type ReqOrderGenerateOrder struct {
	ProductIds          []int   `json:"product_ids" binding:"required"`
	TransportationPrice float64 `json:"transportation_price" binding:"required"`
	Name                string  `json:"name" binding:"required"`
	PhoneNumber         string  `json:"phone_number" binding:"required"`
	Address             string  `json:"address" binding:"required"`
	Message             string  `json:"message"`
}
type ReqOrderGetOrders struct {
	From   int `form:"from" json:"from"`
	Length int `form:"length" json:"length" binding:"required"`
	Status int `form:"status" json:"status"`
}
type ReqOrderGetOneOrder struct {
	OrderId int `form:"order_id" json:"order_id" binding:"required"`
}
type ReqOrderChangeStatus struct {
	OrderId     int    `json:"order_id" binding:"required"`
	Status      int    `json:"status" binding:"required"`
	PayTime     int    `json:"pay_time"`
	VerifyTime  int    `json:"verify_time"`
	ShopMessage string `json:"shop_message"`
}
type ReqOrderAddTrackingNumber struct {
	OrderId        int      `json:"order_id" binding:"required"`
	TrackingNumber []string `json:"tracking_number" binding:"required"`
}

type ProductItem struct {
	Id            int
	OwnerUserId   int // user 表中的 id 列
	Price         float64
	Title         string
	Description   string
	PictureNumber int
	Stock         int
	IsDrop        bool
}
type ResOrderGetOrders struct {
	Orders []Order `json:"orders"`
}
type ProductAndCount struct {
	Products ProductItem `json:"products"`
	Count    int         `json:"count"`
}
type Order struct {
	OrderId             int               `json:"order_id"`
	OwnerUserId         int               `json:"owner_user_id"`
	OwnerShopUserId     int               `json:"owner_shop_user_id"`
	UserPhoneNumber     string            `json:"user_phone_number"`
	NowStatus           int               `json:"now_status"`
	TotalPrice          float64           `json:"total_price"`
	TransportationPrice float64           `json:"transportation_price"`
	DiscountPrice       float64           `json:"discount_price"`
	ProductAndCount     []ProductAndCount `json:"product_and_count"`
	PersonName          string            `json:"person_name"`
	PersonPhoneNumber   string            `json:"person_phone_number"`
	PersonAddress       string            `json:"person_address"`
	OrderTime           int64             `json:"order_time"`
	PayTime             int64             `json:"pay_time"`
	VerifyTime          int64             `json:"verify_time"`
	TrackingNumber      []string          `json:"tracking_number"`
	Message             string            `json:"message"`
	ShopMessage         string            `json:"shop_message"`
}
