package configs

import (
	"fmt"
	"go-todo-apps/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitiateDB() *gorm.DB {
	var UserName = "root"
	var Password = "085283480788"
	var HostName = "127.0.0.1"
	var Port = "3306"
	var DbName = "todo-app"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", UserName, Password, HostName, Port, DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	MigrateDatabase()

	return db
}

func MigrateDatabase() {
	DB.AutoMigrate(models.Todo{})
}
