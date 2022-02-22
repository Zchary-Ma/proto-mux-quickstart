package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/zchary-ma/proto-mux-template/app/server/handler"
)

func main() {
	// ctx := context.Background()

	// Listen and serve HTTP.
	host := ""
	port := "8080"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	routers := handler.NewRouter([]mux.MiddlewareFunc{})
	srv := &http.Server{
		Handler: routers,
		Addr:    host + ":" + port,
	}

	log.Printf("starting server listening at %s:%s ...", host, port)
	log.Fatal(srv.ListenAndServe())
}
