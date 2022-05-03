package param

type ReqUserLogin struct {
	NickName  string `form:"nickName" json:"nick_name" binding:"required"`
	AvatarUrl string `form:"avatarUrl" json:"avatar_url" binding:"required"`
	Code      string `form:"code" json:"code" binding:"required"`
}
type ReqModifyPermission struct {
	UserId     int `form:"userId" json:"user_id" binding:"required"`
	Permission int `form:"permission" json:"permission" binding:"required"`
}

type ResUserGetProfile struct {
	NickName    string
	AvatarUrl   string
	PhoneNumber string
	Permission  int
}
