package model

import "gorm.io/gorm"

type User struct {
	Openid      string
	Permission  int
	PhoneNumber string
	NickName    string
	AvatarUrl   string
	gorm.Model
}

func (m *model) UserRegisterOrDoNothing(openid, nickName, avatarUrl string) error {
	var user = User{}
	err := m.db.Transaction(func(tx *gorm.DB) error {
		// 找有没有
		err := m.db.Model(&User{}).Where("openid = ?", openid).Take(&user).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err // 发生错误
		}
		if err == gorm.ErrRecordNotFound { // 未找到，添加
			err = m.db.Model(&User{}).Create(&User{
				Openid:      openid,
				Permission:  1,
				PhoneNumber: "",
				NickName:    nickName,
				AvatarUrl:   avatarUrl,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil // 提交事务
	})
	if err != nil {
		return err
	}
	return nil
}
func (m *model) UserGetProfile(openid string) (User, error) {
	var user = User{}
	err := m.db.Model(&User{}).Where("openid = ?", openid).Take(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func (m *model) UserModifyPermission(userId, permission int) error {
	err := m.db.Model(&User{}).Where("id = ?", userId).Update("permission", permission).Error
	return err
}
