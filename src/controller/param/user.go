package param

type ReqUserLogin struct {
	NickName  string `form:"nickName" json:"nickName" binding:"required"`
	AvatarUrl string `form:"avatarUrl" json:"avatarUrl" binding:"required"`
	Code      string `form:"code" json:"code" binding:"required"`
}
type ReqModifyPermission struct {
	UserId     int `form:"userId" json:"userId" binding:"required"`
	Permission int `form:"permission" json:"permission" binding:"required"`
}
type ReqGetToken struct {
	OpenId     string `form:"openid" json:"openid" binding:"required"`
	SessionKey string `form:"session_key" json:"session_key"`
}

type ResUserGetProfile struct {
	NickName    string
	AvatarUrl   string
	PhoneNumber string
	Permission  int
}
