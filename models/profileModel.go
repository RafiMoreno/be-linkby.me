package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	DisplayName    string `json:"displayName"`
	PrimaryColor   string `json:"primaryColor" gorm:"default:#A44646"`
	SecondaryColor string `json:"secondaryColor" gorm:"default:#FFFFFF"`
	Description    string `json:"description"`
	DisplayPicture string `json:"displayPicture"`
	Links          []Link `json:"-"`
}
