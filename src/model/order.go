package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type ProductOrder struct {
	OwnerUserId         int
	OwnerShopUserId     int
	NowStatus           int
	TotalPrice          float64
	TransportationPrice float64
	DiscountPrice       float64
	ProductAndCount     string
	PersonName          string
	PersonPhoneNumber   string
	PersonAddress       string
	PayTime             time.Time
	VerifyTime          time.Time
	TrackingNumber      string
	Message             string
	gorm.Model
}

func (m *model) OrderGenerateOrder(productsIds [][]int,
	userId int, transportationPrice float64,
	address Address, message string) error {

	err := m.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range productsIds {
			totalPrice := 0.0
			var ownerShopUserId int
			var productAndCount string
			for j, v2 := range v {
				// 找到商品
				product, err := m.ShopGetProduct(v2)
				if err != nil {
					return err
				}
				ownerShopUserId = product.OwnerUserId
				// 找到购物车中的 count
				var cart Cart
				err = m.db.Model(&Cart{}).Where("owner_user_id = ? AND product_id = ? AND product_count > 0", userId, v2).Take(&cart).Error
				if err != nil {
					return err
				}
				totalPrice += product.Price * (float64)(cart.ProductCount)
				// 构造 productAndCount
				productAndCount += strconv.Itoa(v2) + ","
				productAndCount += strconv.Itoa(cart.ProductCount)
				if j != len(v)-1 {
					productAndCount += ","
				}
			}

			err := m.db.Model(&ProductOrder{}).Create(&ProductOrder{
				OwnerUserId:         userId,
				OwnerShopUserId:     ownerShopUserId,
				NowStatus:           1,
				TotalPrice:          totalPrice + transportationPrice,
				TransportationPrice: transportationPrice,
				DiscountPrice:       0.0,
				ProductAndCount:     productAndCount,
				PersonName:          address.Name,
				PersonPhoneNumber:   address.PhoneNumber,
				PersonAddress:       address.Address,
				PayTime:             time.Unix(0, 0),
				VerifyTime:          time.Unix(0, 0),
				TrackingNumber:      "",
				Message:             message,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}
func (m *model) OrderGetOneOrder(orderId, userId int) (ProductOrder, error) {
	var order ProductOrder
	err := m.db.Model(&ProductOrder{}).Where("id = ?", orderId).Take(&order).Error
	if err != nil {
		return ProductOrder{}, err
	}
	if order.OwnerUserId != userId && order.OwnerShopUserId != userId {
		return ProductOrder{}, errors.New("not your order")
	}
	return order, nil
}
func (m *model) OrderGetUserOrder(userId, length, from, status int) ([]ProductOrder, error) {
	orders, err := m.orderGetSbOrder(userId, length, from, status, "owner_user_id")
	if err != nil {
		return []ProductOrder{}, err
	}
	return orders, err
}
func (m *model) OrderGetShopOrder(userId, length, from, status int) ([]ProductOrder, error) {
	orders, err := m.orderGetSbOrder(userId, length, from, status, "owner_shop_user_id")
	if err != nil {
		return []ProductOrder{}, err
	}
	return orders, err
}
func (m *model) orderGetSbOrder(userId, length, from, status int, name string) ([]ProductOrder, error) {
	var orders []ProductOrder
	err := m.db.Model(&ProductOrder{}).Where(fmt.Sprintf("%s = ? AND now_status = ?", name), userId, status).
		Limit(length).Offset(from).Scan(&orders).Error
	return orders, err
}
func (m *model) OrderDeleteOrder(userId, orderId int) error {
	return m.db.Model(&ProductOrder{}).Where("owner_user_id = ?", userId).Delete(&ProductOrder{}, orderId).Error
}
func (m *model) OrderChangeStatus(userId, orderId, status int, payTime, verifyTime time.Time) error {
	return m.db.Model(&ProductOrder{}).Where("id = ? AND owner_shop_user_id = ?", orderId, userId).Updates(ProductOrder{
		NowStatus:  status,
		PayTime:    payTime,
		VerifyTime: verifyTime,
	}).Error
}
func (m *model) OrderAddTrackingNumber(userId, orderId int, trackingNumber string) error {
	return m.db.Model(&ProductOrder{}).Where("id = ? AND owner_shop_user_id = ?", orderId, userId).Update("tracking_number", trackingNumber).Error
}
