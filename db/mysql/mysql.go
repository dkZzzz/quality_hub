package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func Init() {
	InitDB()
}

func InitDB() (*gorm.DB, error) {
	dsn := "root:20021210@tcp(127.0.0.1:3306)/codecctest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	db.AutoMigrate(&User{})
	return db, err
}
