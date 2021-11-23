package database

import (
	"github.com/miftahulhidayati/rest-api-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysqlDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/order_by?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(models.Person{})
	return db
}