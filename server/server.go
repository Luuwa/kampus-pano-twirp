package server

import (
	"github.com/kamp-us/pano-service/internal/panoserver"
)

type PanoAPIServer struct {
	backend backend.Backender
}

func NewPanoAPIServer(backend backend.Backender) *PanoAPIServer {
	return &PanoAPIServer{backend: backend}
}
