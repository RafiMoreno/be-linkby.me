package models

import "gorm.io/gorm"

type Profile struct{
	gorm.Model
	DisplayName string	`json:"display_name"`
	PrimaryColor string `json:"primary_color" gorm:"default:#A44646"`
	SecondaryColor string `json:"secondary_color" gorm:"default:#FFFFFF"`
	Description string `json:"description"`
	DisplayPicture string `json:"display_picture"`
}