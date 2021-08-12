package database

import (
	"gorm.io/gorm"

	"example.com/hello/models"
)
import "gorm.io/driver/mysql"

var DB *gorm.DB
func Connect(){
	// making a connection to mysql
	connection, err := gorm.Open(mysql.Open("root:mypass@/login"), &gorm.Config{})

	if err != nil{
		// panic will stop the app if an error happens
		panic("could not connect to database")}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
