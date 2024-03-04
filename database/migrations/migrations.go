package migrations

import (
	"gorm.io/gorm"
	"github.com/dejauls/task-5-pbi-fullstack-developer--Rajaul-Bani-Safar-/models"
)


func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Photo{})
}
