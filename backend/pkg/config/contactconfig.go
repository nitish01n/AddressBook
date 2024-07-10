package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:Nitish@26@tcp(localhost:3306)/phonebook")

	if err != nil {
		panic(err.Error())
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
