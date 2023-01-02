package helper

import (
	"github.com/kamp-us/pano-service/internal/models"
	api "github.com/kamp-us/pano-service/rpc/pano"
)

func ConvertToPostModel(c *models.Post) *api.Post {
	return &api.Post{
		ID:      c.ID.String(),
		Title:   c.Title,
		Site:    c.Site,
		Slug:    c.Slug,
		Content: c.Content,
		UserID:  c.UserID,
	}
}
