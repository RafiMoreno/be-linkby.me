package models

import "gorm.io/gorm"

type Profile struct{
	gorm.Model
	DisplayName string
	PrimaryColor string `gorm:"default:#A44646"`
	SecondaryColor string `gorm:"default:#FFFFFF"`
	Description string
	DisplayPicture string
}