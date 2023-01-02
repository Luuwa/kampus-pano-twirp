package server

import (
	"context"
	api "github.com/kamp-us/pano-service/rpc/pano"
	"github.com/kamp-us/pano-service/server/helper"
	"github.com/twitchtv/twirp"
)

func (s *PanoAPIServer) CreatePost(ctx context.Context, req *api.CreatePostRequest) (*api.Post, error) {
	if err := validateCreatePostRequest(req); err != nil {
		return nil, err
	}

	post, err := s.backend.CreatePost(ctx, req.Title, req.Site, req.Content, req.UserID)
	if err != nil {
		return nil, twirp.InternalErrorWith(err)
	}

	return helper.ConvertToPostModel(post), nil
}

func validateCreatePostRequest(req *api.CreatePostRequest) error {
	if req.Title == "" {
		return twirp.RequiredArgumentError("title")
	}
	if req.Site == "" {
		return twirp.RequiredArgumentError("site")
	}
	if req.Content == "" {
		return twirp.RequiredArgumentError("content")
	}
	if len(req.UserID) == 0 {
		return twirp.RequiredArgumentError("userid")
	}
	return nil
}
