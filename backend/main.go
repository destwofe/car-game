package main

import (
	"fmt"
	"toycar_backend/database"
	"toycar_backend/models"
	"toycar_backend/routers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := database.DbURL(database.BuildDBConfig())
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	database.DB = db
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.CarGame{})

	r := routers.SetupRouter()
	r.Run()
}
