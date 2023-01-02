package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Title   string
	Site    string
	Slug    string `gorm:"index"`
	Content string
	UserID  string `gorm:"index"`
}

func AutoMigrate(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	if err := db.AutoMigrate(&Post{}); err != nil {
		return err
	}
	return nil
}
