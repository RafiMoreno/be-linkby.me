package initializers

import (
	"example/be-linkby.me/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
