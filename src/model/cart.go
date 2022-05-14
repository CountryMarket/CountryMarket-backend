package model

import "gorm.io/gorm"

type Cart struct {
	OwnerUserId  int
	ProductId    int
	ProductCount int
	gorm.Model
}
type CartAndProduct struct {
	ProductId    int
	ProductCount int
	Price        float64
	Title        string
	Description  string
	OwnerUserId  int
	Stock        int
	IsDrop       bool
}

func (m *model) CartGetInCart(userId, productId int) (int, error) {
	var cart Cart
	err := m.db.Model(&Cart{}).Where("owner_user_id = ? AND product_id = ? AND product_count > 0", userId, productId).Take(&cart).Error
	return cart.ProductCount, err
}
func (m *model) CartGetUserProducts(userId, from, length int) ([]CartAndProduct, error) {
	var carts []CartAndProduct
	err := m.db.Model(&Cart{}).
		Select("cart.product_id, cart.product_count, product.price, product.title, product.description, product.owner_user_id, product.is_drop, product.stock").
		Joins("LEFT JOIN product ON cart.product_id = product.id").
		Where("cart.owner_user_id = ? AND cart.product_count > 0", userId).
		Limit(length).Offset(from).Scan(&carts).Error
	if err != nil {
		return []CartAndProduct{}, err
	}
	return carts, err
}
func (m *model) CartGetCart(userId int, ids []int) ([]CartAndProduct, error) {
	var carts []CartAndProduct
	err := m.db.Model(&Cart{}).
		Select("cart.product_id, cart.product_count, product.price, product.title, product.description, product.owner_user_id, product.is_drop, product.stock").
		Joins("LEFT JOIN product ON cart.product_id = product.id").
		Where("cart.owner_user_id = ? AND cart.product_id IN ? AND cart.product_count > 0", userId, ids).Scan(&carts).Error
	if err != nil {
		return []CartAndProduct{}, err
	}
	return carts, err
}
func (m *model) CartAddProduct(userId, productId int) (int, error) {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		var cart Cart
		result := m.db.Model(&Cart{}).Where("owner_user_id = ? AND product_id = ?", userId, productId).Take(&cart)
		err := result.Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if err == gorm.ErrRecordNotFound { // 没有找到记录，创建
			err = m.db.Model(&Cart{}).Create(&Cart{
				OwnerUserId:  userId,
				ProductId:    productId,
				ProductCount: 1,
			}).Error
			return err
		} else {
			err = result.Update("product_count", cart.ProductCount+1).Error
			return err
		}
	})
	if err != nil {
		return 0, err
	}
	var cart Cart
	err = m.db.Model(&Cart{}).Where("owner_user_id = ? AND product_id = ?", userId, productId).Take(&cart).Error
	if err != nil {
		return 0, err
	}
	return cart.ProductCount, err
}
func (m *model) CartReduceProduct(userId, productId int) (int, error) {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		var cart Cart
		result := m.db.Model(&Cart{}).Where("owner_user_id = ? AND product_id = ?", userId, productId).Take(&cart)
		err := result.Error
		if err != nil {
			return err
		}
		if cart.ProductCount <= 0 {
			return nil
		}
		err = result.Update("product_count", cart.ProductCount-1).Error
		return err
	})
	if err != nil {
		return 0, err
	}
	var cart Cart
	err = m.db.Model(&Cart{}).Where("owner_user_id = ? AND product_id = ?", userId, productId).Take(&cart).Error
	if err != nil {
		return 0, err
	}
	return cart.ProductCount, err
}
func (m *model) CartModifyProduct(userId, productId, modifyCount int) (int, error) {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		if modifyCount < 0 {
			return nil
		}
		var cart Cart
		result := m.db.Model(&Cart{}).Where("owner_user_id = ? AND product_id = ?", userId, productId).Take(&cart)
		err := result.Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if err == gorm.ErrRecordNotFound { // 没有找到记录，创建
			err = m.db.Model(&Cart{}).Create(&Cart{
				OwnerUserId:  userId,
				ProductId:    productId,
				ProductCount: modifyCount,
			}).Error
			return err
		} else {
			err = result.Update("product_count", modifyCount).Error
			return err
		}
	})
	if err != nil {
		return 0, err
	}
	var cart Cart
	err = m.db.Model(&Cart{}).Where("owner_user_id = ? AND product_id = ?", userId, productId).Take(&cart).Error
	if err != nil {
		return 0, err
	}
	return cart.ProductCount, err
}
