package param

type ReqProductGetTabProducts struct {
	TabId int `form:"tabId" json:"tabId" binding:"required"`
}
type ReqProductModifyTab struct {
	TabId    int    `form:"tabId" json:"tabId" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required"`
	Products []int  `form:"products" json:"products" binding:"required"`
}
type ReqProductAddTab struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Products []int  `form:"products" json:"products" binding:"required"`
}
type ReqProductDeleteTab struct {
	TabId int `form:"tabId" json:"tabId" binding:"required"`
}

type TabItem struct {
	Id   int
	Name string
}
type ResProductGetTabList struct {
	Tabs []TabItem
}
type TabProductsItem struct {
	Id          int
	Price       float64
	Title       string
	Description string
	OwnerUserId int
	IsDrop      bool
}
type ResProductGetTabProducts struct {
	Products []TabProductsItem
}
