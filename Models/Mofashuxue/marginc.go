package Mofashuxue

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var magicDb *gorm.DB

func init() {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)

	dbType = "mysql"
	dbName = "mofashuxue"
	user = "root"
	password = "123456"
	host = "127.0.0.1:3306"
	magicDb, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		fmt.Println("connected failed:", err)
	}

	magicDb.SingularTable(true)
	magicDb.LogMode(true)
	magicDb.DB().SetMaxIdleConns(20)
	magicDb.DB().SetMaxOpenConns(100)
}
