package initializers

import (
	"example/be-linkby.me/models"
)

func SyncDatabase() {
	db.AutoMigrate(&models.User{})
}
