package main

import (
	"example/musicAPI/controllers"
	"example/musicAPI/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/musics", controllers.MusicCreate)
	r.GET("/musics", controllers.MusicRead)
	r.GET("/music/:id", controllers.MusicReadById)
	r.PUT("/music/:id", controllers.UpdateMusic)
	r.DELETE("/music/:id", controllers.DeleteMusic)

	r.Run()
}
