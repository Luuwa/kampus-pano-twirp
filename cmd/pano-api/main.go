package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kamp-us/pano-service/internal/db"
	"github.com/kamp-us/pano-service/internal/models"
	backend "github.com/kamp-us/pano-service/internal/panoserver"
	panoapi "github.com/kamp-us/pano-service/rpc/pano"
	"github.com/kamp-us/pano-service/server"
)

func main() {
	dbClient, err := db.NewPostgreSQLConnection(db.PostgreSQLConfig{
		Host:     "top2.nearest.of.panoservice-db.internal",
		Port:     5432,
		Username: "panoservice",
		Password: "nvtTe5igBmN5cuv",
		DbName:   "panoservice",
	})

	if err != nil {
		log.Fatal("error while creating a db connection pool", err)
	}

	if err != nil {
		return
	}
	err = models.AutoMigrate(dbClient)
	if err != nil {
		return
	}
	fmt.Println(err)
	postgreSQLBackend := backend.NewPostgreSQLBackend(dbClient)

	s := server.NewPanoAPIServer(postgreSQLBackend)
	twirpHandler := panoapi.NewPanoServer(s)

	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	mux.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("OK"))
	}))

	http.ListenAndServe("0.0.0.0:8080", mux)
}
