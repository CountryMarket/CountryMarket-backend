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
	IsDefault   bool
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
	if count >= 100 {
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
func (m *model) AddressGetDefaultAddress(userId int) (int, error) {
	var address Address
	err := m.db.Model(&Address{}).Where("is_default = ? AND owner_user_id = ?", 1, userId).Take(&address).Error
	return int(address.ID), err
}
func (m *model) AddressModifyDefaultAddress(userId, addressId int) error {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		var address Address
		db := m.db.Model(&Address{}).Where("is_default = ? AND owner_user_id = ?", 1, userId)
		err := db.Take(&address).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if err != gorm.ErrRecordNotFound {
			err = db.Update("is_default", 0).Error
			if err != nil {
				return err
			}
		}

		err = m.db.Model(&Address{}).Where("id = ? AND owner_user_id = ?", addressId, userId).Update("is_default", 1).Error
		return err
	})
	return err
}
