package models

import "gorm.io/gorm"

type Music struct {
	gorm.Model

	Title  string
	Artist string
}
