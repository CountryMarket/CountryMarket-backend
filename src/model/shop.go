package model

import "gorm.io/gorm"

type Product struct {
	OwnerUserId         int // user 表中的 id 列
	Price               float64
	Title               string
	Description         string
	PictureNumber       int
	IsDrop              bool
	Stock               int
	Detail              string
	DetailPictureNumber int
	gorm.Model
}

func (m *model) ShopAddProduct(product Product) (int, error) {
	err := m.db.Model(&Product{}).Create(&product).Error
	if err != nil {
		return 0, err
	}
	return int(product.ID), nil
}
func (m *model) ShopUpdateProduct(product Product, productId int) error {
	err := m.db.Model(&Product{}).Where("id = ?", productId).Updates(product).Error
	if err != nil {
		return err
	}
	return nil
}
func (m *model) ShopGetProduct(productId int) (Product, error) {
	product := Product{}
	err := m.db.Model(&Product{}).Where("id = ?", productId).Take(&product).Error
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
func (m *model) ShopGetOwnerProducts(userId, from, length int) ([]Product, error) {
	var products []Product
	err := m.db.Model(&Product{}).Where("owner_user_id = ?", userId).Limit(length).Offset(from).Order("id DESC").Scan(&products).Error
	if err != nil {
		return []Product{}, err
	}
	return products, err
}
func (m *model) ShopDropProduct(userId, productId int) error {
	return m.db.Model(&Product{}).Where("id = ? AND owner_user_id = ?", productId, userId).Update("is_drop", 1).Error
}
func (m *model) ShopPutProduct(userId, productId int) error {
	return m.db.Model(&Product{}).Where("id = ? AND owner_user_id = ?", productId, userId).Update("is_drop", 0).Error
}
