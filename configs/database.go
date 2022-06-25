package configs

import (
	"fmt"
	"go-todo-apps/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitiateDB() *gorm.DB {
	UserName := "root"
	Password := "085283480788"
	HostName := "db"
	Port := "3306"
	DbName := "todo-app"
	connStr := UserName + ":" + Password + "@tcp(" + HostName + ":" + Port + ")/" + DbName
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", UserName, Password, HostName, Port, DbName)
	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println(connStr)

	DB = db
	MigrateDatabase()

	return db
}

func MigrateDatabase() {
	DB.AutoMigrate(models.Todo{})
}
