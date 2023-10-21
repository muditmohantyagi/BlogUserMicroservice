package migration

import (
	"fmt"

	"blog.com/config"
	"blog.com/model"
)

func Migrate() {
	db := config.GoConnect()
	db.AutoMigrate(&model.User{})
	fmt.Println("Migration successfull")
}
