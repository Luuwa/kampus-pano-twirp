package backend

import (
	"context"
	"github.com/gosimple/slug"
	"github.com/kamp-us/pano-service/internal/models"
)

func (b *PostgreSQLBackend) CreatePost(ctx context.Context, title string, site string, content string, userID string) (*models.Post, error) {

	post := models.Post{
		Title:   title,
		Site:    site,
		Slug:    slug.MakeLang(title, "tr"),
		Content: content,
		UserID:  userID,
	}

	result := b.DB.Create(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	return &post, nil
}
