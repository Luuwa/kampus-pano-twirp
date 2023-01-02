package backend

import (
	"context"
	"github.com/kamp-us/pano-service/internal/models"
	"gorm.io/gorm"
)

type Backender interface {
	GetPost(ctx context.Context, id string) (*models.Post, error)
}

type PostgreSQLBackend struct {
	DB *gorm.DB
}

func NewPostgreSQLBackend(db *gorm.DB) Backender {
	return &PostgreSQLBackend{
		DB: db,
	}
}
