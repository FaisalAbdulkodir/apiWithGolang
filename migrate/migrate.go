package main

import (
	"example/musicAPI/initializers"
	"example/musicAPI/models"
)

func init() {
	initializers.ConnectDB()
	initializers.LoadEnv()
}

func main() {
	// db.AutoMigrate(&Product{})
	initializers.DB.AutoMigrate(&models.Music{})
}
