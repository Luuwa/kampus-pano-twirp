package backend

import (
	"context"
	"github.com/kamp-us/pano-service/internal/models"
)

func (b *PostgreSQLBackend) GetPost(ctx context.Context, id string) (*models.Post, error) {
	post := models.Post{}
	result := b.DB.First(&post, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	query := b.DB.Preload("Categories").Find(&post)
	if query.Error != nil {
		return nil, query.Error
	}

	return &post, nil
}
