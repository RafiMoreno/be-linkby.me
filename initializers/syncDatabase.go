package initializers

import (
	"github.com/RafiMoreno/be-linkby.me/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Profile{}, &models.Link{})
}
