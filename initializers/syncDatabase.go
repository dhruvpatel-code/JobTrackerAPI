package initializers

import "github.com/dhruvpatel-code/JobTrackerAPI/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
