package controllers

import (
	"example/musicAPI/initializers"
	"example/musicAPI/models"

	"github.com/gin-gonic/gin"
)

// Create
func MusicCreate(c *gin.Context) {

	var body struct {
		Title  string
		Artist string
	}

	c.Bind(&body)

	post := models.Music{Title: body.Title, Artist: body.Artist}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"music": post,
	})
}

// Read All
func MusicRead(c *gin.Context) {
	var musics []models.Music
	initializers.DB.Find(&musics)

	c.JSON(200, gin.H{
		"musics": musics,
	})
}

// Read By Id
func MusicReadById(c *gin.Context) {

	id := c.Param("id")

	var music []models.Music
	initializers.DB.First(&music, id)

	c.JSON(200, gin.H{
		"music": music,
	})
}

// Update
func UpdateMusic(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title  string
		Artist string
	}

	c.Bind(&body)

	var music []models.Music
	initializers.DB.First(&music, id)

	initializers.DB.Model(&music).Updates(models.Music{
		Title:  body.Title,
		Artist: body.Artist,
	})

	c.JSON(200, gin.H{
		"music": music,
	})

}

// Delete
func DeleteMusic(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Music{}, id)

	c.Status(200)
}
