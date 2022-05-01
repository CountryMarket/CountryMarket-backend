package model

import "gorm.io/gorm"

type Product struct {
	OwnerUserId   int // user 表中的 id 列
	Price         float64
	Title         string
	Description   string
	PictureNumber int
	gorm.Model
}

func (m *model) ShopAddProduct(product Product) error {
	err := m.db.Model(&Product{}).Create(&product).Error
	if err != nil {
		return err
	}
	return nil
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
	err := m.db.Model(&Product{}).Where("owner_user_id = ?", userId).Limit(length).Offset(from).Scan(&products).Error
	if err != nil {
		return []Product{}, err
	}
	return products, err
}
