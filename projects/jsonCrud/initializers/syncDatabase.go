package initializers

import "jsonCrud/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Post{})
}
