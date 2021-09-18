package dao

import (
	"cloudcost-demo/model"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func InitMySQL() (err error) {
	dsn := "root:mistral@tcp(172.16.32.173:3306)/cloudcost?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		err = fmt.Errorf("connection DB timeout err(%v)", err)
		return
	}
	db.AutoMigrate(&model.Auth{})
	return db.DB().Ping()
}

func Close() {
	db.Close()
}
