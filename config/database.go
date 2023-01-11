package config

import (
	"fmt"

	"github.com/ndzf/fiber-gorm-mysql.git/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI = "root:@tcp(localhost:3306)/fiber_crud?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() {
	var err error

	Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Dog{})
	fmt.Println("Migrate success")
}
