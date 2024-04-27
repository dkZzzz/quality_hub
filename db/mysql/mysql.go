package mysql

import (
	"github.com/dkZzzz/quality_hub/config"
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
	user := config.Cfg.MysqlUser
	password := config.Cfg.MysqlPassword
	host := config.Cfg.MysqlHost
	port := config.Cfg.MysqlPort
	dbname := config.Cfg.MysqlDatabase
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	db.AutoMigrate(&User{}, &Project{}, &Report{}, &Issue{}, &Advice{})
	return db, err
}
