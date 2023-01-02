package main

import (
	"fmt"
	"github.com/kamp-us/pano-service/internal/db"
	"github.com/kamp-us/pano-service/internal/models"
	"github.com/kamp-us/pano-service/internal/panoserver"
	panoapi "github.com/kamp-us/pano-service/rpc/pano"
	"github.com/kamp-us/pano-service/server"
	"log"
	"net/http"
)

func main() {
	dbClient, err := db.NewPostgreSQLConnection(db.PostgreSQLConfig{
		Host:     "192.168.1.103",
		Port:     5555,
		Username: "pgtest",
		Password: "pgtest",
		DbName:   "pgtest1",
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

	http.ListenAndServe(":80", mux)
}
