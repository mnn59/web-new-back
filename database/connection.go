package database

import (
	"github.com/mnn59/web-new-back/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect()  {
	connection, err := gorm.Open(mysql.Open("root:@/go_test"), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}