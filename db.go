package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() {
	dsn := "host=localhost user=****** password=****** dbname=users port=**** sslmode=disable"

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't Connect to Postgres")
		panic(err)
	} else {
		fmt.Println("Connected to Postgres successfully!")
	}
	db.AutoMigrate(&User{})
}
