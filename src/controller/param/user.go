package param

type ReqUserLogin struct {
	NickName  string `form:"nickName" json:"nick_name" binding:"required"`
	AvatarUrl string `form:"avatarUrl" json:"avatar_url" binding:"required"`
	Code      string `form:"code" json:"code" binding:"required"`
}
type ResUserGetProfile struct {
	NickName    string `json:"nickName"`
	AvatarUrl   string `json:"avatarUrl"`
	PhoneNumber string `json:"phoneNumber"`
	Permission  int    `json:"permission"`
}
