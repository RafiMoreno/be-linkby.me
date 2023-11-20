package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Url        string `json:"url"`
	Title      string `json:"title"`
	ClickCount uint64 `json:"clickCount"`
	IconUrl    string `json:"iconUrl"`
}
