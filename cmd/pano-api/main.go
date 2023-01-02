package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kamp-us/pano-service/internal/panoserver"

	"github.com/kamp-us/pano-service/internal/db"
	panoapi "github.com/kamp-us/pano-service/rpc/pano"
	"github.com/kamp-us/pano-service/server"
)

func main() {
	dbClient, err := db.NewPostgreSQLConnection(db.PostgreSQLConfig{
		Host:     "192.168.1.103",
		Port:     5555,
		Username: "pgtest",
		Password: "pgtest",
		DbName:   "pgtest",
	})

	if err != nil {
		log.Fatal("error while creating a db connection pool", err)
	}

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
