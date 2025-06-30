package initializers

import "jwtgo/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
