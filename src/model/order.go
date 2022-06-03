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
	UserPhoneNumber     string
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
	ShopMessage         string
	gorm.Model
}

func (m *model) OrderGenerateOrder(productsIds [][]int,
	userId int, phoneNumber string, transportationPrice float64,
	address Address, message string) ([]int, error) {
	var resId []int
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
				// 判断商品库存是否充足
				if product.Stock < cart.ProductCount {
					return errors.New(fmt.Sprintf("stock not enough for product(id = %d)", product.ID))
				}
				totalPrice += product.Price * (float64)(cart.ProductCount)
				// 构造 productAndCount
				productAndCount += strconv.Itoa(v2) + ","
				productAndCount += strconv.Itoa(cart.ProductCount)
				if j != len(v)-1 {
					productAndCount += ","
				}
				// 从库存中删除
				err = m.ShopUpdateProduct(Product{
					Stock: product.Stock - cart.ProductCount,
				}, v2)
				if err != nil {
					return err
				}
				// 从购物车中删除
				_, err = m.CartModifyProduct(userId, v2, 0)
				if err != nil {
					return err
				}
			}

			p := ProductOrder{
				OwnerUserId:         userId,
				OwnerShopUserId:     ownerShopUserId,
				UserPhoneNumber:     phoneNumber,
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
			}

			err := m.db.Model(&ProductOrder{}).Create(&p).Error

			if err != nil {
				return err
			}

			resId = append(resId, (int)(p.ID))

		}
		return nil
	})
	return resId, err
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
	if status != 0 {
		fmtStr := fmt.Sprintf("%s = ? AND now_status = ?", name)
		var orders []ProductOrder
		err := m.db.Model(&ProductOrder{}).Where(fmtStr, userId, status).
			Limit(length).Offset(from).Order("id DESC").Scan(&orders).Error
		return orders, err
	} else {
		fmtStr := fmt.Sprintf("%s = ?", name)
		var orders []ProductOrder
		err := m.db.Model(&ProductOrder{}).Where(fmtStr, userId).
			Limit(length).Offset(from).Order("id DESC").Scan(&orders).Error
		return orders, err
	}
}
func (m *model) OrderDeleteOrder(userId, orderId int) error {
	return m.db.Model(&ProductOrder{}).Where("owner_user_id = ?", userId).Delete(&ProductOrder{}, orderId).Error
}
func (m *model) OrderChangeStatus(userId, orderId, status int, payTime, verifyTime time.Time, shopMessage string) error {
	p := ProductOrder{
		NowStatus:   status,
		PayTime:     payTime,
		VerifyTime:  verifyTime,
		ShopMessage: shopMessage,
	}
	if shopMessage == "" {
		p = ProductOrder{
			NowStatus:  status,
			PayTime:    payTime,
			VerifyTime: verifyTime,
		}
	}
	return m.db.Model(&ProductOrder{}).Where("id = ?", orderId).Updates(p).Error
}
func (m *model) OrderAddTrackingNumber(userId, orderId int, trackingNumber string) error {
	return m.db.Model(&ProductOrder{}).Where("id = ?", orderId).Update("tracking_number", trackingNumber).Error
}
