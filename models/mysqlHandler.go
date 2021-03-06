package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GormConnect() *gorm.DB {
	dblink := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true",
		beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpassword"),
		beego.AppConfig.String("dbhost"),
		beego.AppConfig.String("dbname"),
	)
	dbms := beego.AppConfig.String("dbdriver")
	db, err := gorm.Open(dbms, dblink)
	if err != nil {
		panic(err.Error())
	}
	return db
}
