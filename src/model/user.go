package model

type Data struct { // for test
	Openid     string `gorm:"column:openid"`
	TestString string `gorm:"column:test_string"`
}

func (m *model) Test(openid string) (Data, error) { // for test
	var d = Data{}
	err := m.db.Table("test").Where("openid = ?", openid).First(&d).Error
	return d, err
}
