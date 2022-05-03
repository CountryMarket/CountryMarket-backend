package model

import (
	"errors"
	"gorm.io/gorm"
)

type Address struct {
	Name        string
	Address     string
	PhoneNumber string
	OwnerUserId int
	gorm.Model
}

func (m *model) AddressGetOneAddress(addressId int) (Address, error) {
	var address Address
	err := m.db.Model(&Address{}).Where("id = ?", addressId).Take(&address).Error
	return address, err
}
func (m *model) AddressAddAddress(address Address) error {
	var count int64
	err := m.db.Model(&Address{}).Where("owner_user_id = ?", address.OwnerUserId).Count(&count).Error
	if err != nil {
		return err
	}
	if count >= 5 {
		return errors.New("more than 5 addresses")
	}
	return m.db.Model(&Address{}).Create(&address).Error
}
func (m *model) AddressModifyAddress(addressId int, address Address) error {
	return m.db.Model(&Address{}).Where("id = ?", addressId).Updates(address).Error
}
func (m *model) AddressDeleteAddress(addressId int) error {
	return m.db.Delete(&Address{}, addressId).Error
}
func (m *model) AddressGetAddress(userId int) ([]Address, error) {
	var addresses []Address
	return addresses, m.db.Model(&Address{}).Where("owner_user_id = ?", userId).Scan(&addresses).Error
}
