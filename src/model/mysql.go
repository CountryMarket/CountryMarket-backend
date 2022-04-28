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
	Test(string) (Data, error) // for test
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
