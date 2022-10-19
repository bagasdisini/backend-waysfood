package database

import (
	"backend/models"
	"backend/pkg/mysql"
	"fmt"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Category{}, &models.User{}, &models.Product{}, &models.Transaction{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed!")
	}

	fmt.Println("Migration Successful!")
}
