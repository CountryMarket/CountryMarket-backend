package param

type ReqAddressAddAddress struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Address     string `form:"address" json:"address" binding:"required"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" binding:"required"`
}
type ReqAddressModifyAddress struct {
	AddressId   int    `form:"addressId" json:"addressId" binding:"required"`
	Name        string `form:"name" json:"name" binding:"required"`
	Address     string `form:"address" json:"address" binding:"required"`
	PhoneNumber string `form:"phoneNumber" json:"phoneNumber" binding:"required"`
}
type ReqAddressDeleteAddress struct {
	AddressId int `form:"addressId" json:"addressId" binding:"required"`
}
type ReqAddressModifyDefaultAddress struct {
	AddressId int `json:"address_id"`
}

type AddressItem struct {
	AddressId   int
	Name        string
	Address     string
	PhoneNumber string
}
type ResAddressGetAddress struct {
	Address []AddressItem
}
type ResAddressGetDefaultAddress struct {
	AddressId int `json:"address_id"`
}
