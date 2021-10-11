package migrations

import (
	"github.com/TheNeoCarvalho/go-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
