package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var dbInstance *gorm.DB

type Model interface {
	Close() // 关闭数据库连接
	Abort() // 终止操作，用于如事务的取消
	// user
	UserRegisterOrDoNothing(openid, nickName, avatarUrl string) error
	UserGetProfile(openid string) (User, error)
	UserModifyPermission(userId, permission int) error
	// address
	AddressGetOneAddress(addressId int) (Address, error)
	AddressAddAddress(address Address) error
	AddressModifyAddress(addressId int, address Address) error
	AddressDeleteAddress(addressId int) error
	AddressGetAddress(userId int) ([]Address, error)
	AddressGetDefaultAddress(userId int) (int, error)
	AddressModifyDefaultAddress(userId, addressId int) error
	// shop
	ShopAddProduct(product Product) (int, error)
	ShopUpdateProduct(product Product, productId int) error
	ShopGetProduct(productId int) (Product, error)
	ShopGetOwnerProducts(userId, from, length int) ([]Product, error)
	ShopDropProduct(userId, productId int) error
	ShopPutProduct(userId, productId int) error
	// cart
	CartGetUserProducts(userId, from, length int) ([]CartAndProduct, error)
	CartGetCart(userId int, ids []int) ([]CartAndProduct, error)
	CartGetInCart(userId, productId int) (int, error)
	CartAddProduct(userId, productId int) (int, error)
	CartReduceProduct(userId, productId int) (int, error)
	CartModifyProduct(userId, productId, modifyCount int) (int, error)
	// product
	ProductGetTabList() ([]int, []string, error)
	ProductGetTabProducts(tabId int) ([]Product, error)
	ProductModifyTabProducts(tabId int, productTab ProductTab) error
	ProductAddTabProducts(productTab ProductTab) error
	ProductDeleteTabProducts(tabId int) error
	ProductGetHomeTab(from, length int) ([]Product, error)
	// order
	OrderGenerateOrder(productsIds [][]int, userId int, phoneNumber string, transportationPrice float64, address Address, message string) ([]int, error)
	OrderGetOneOrder(orderId, userId int) (ProductOrder, error)
	OrderGetUserOrder(userId, length, from, status int) ([]ProductOrder, error)
	OrderGetShopOrder(userId, length, from, status int) ([]ProductOrder, error)
	OrderDeleteOrder(userId, orderId int) error
	OrderChangeStatus(userId, orderId, status int, payTime, verifyTime time.Time, shopMessage string) error
	OrderAddTrackingNumber(userId, orderId int, trackingNumber string) error
	// comment
	CommentGetProductComment(productId int) ([]UserComment, error)
	CommentAddComments(userId int, userComments []UserComment) error
	CommentAddComment(userId, productId int, comment string) error
	CommentDeleteComment(commentId int) error
	// search
	Search(key string) ([]Product, error)
}

type model struct {
	db    *gorm.DB
	abort bool
}

func init() {
	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	log.Println("start init MySQL with ", source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		log.Println("database open error, err=", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("database init error, err=", err.Error())
	}

	sqlDB.SetMaxIdleConns(100)          // 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(200)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置了连接可复用的最大时间

	dbInstance = db

	log.Println("MySQL init finished.")
}
func (m *model) Close() {
}
func (m *model) Abort() {
	m.abort = true
}
func Get() Model {
	return &model{
		dbInstance,
		false,
	}
}
