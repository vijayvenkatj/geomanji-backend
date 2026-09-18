package main

import (
	"log"

	"github.com/joho/godotenv"
	geomanjihttp "github.com/vijayvenkatj/geomanji-backend/pkg/http"
	"github.com/vijayvenkatj/geomanji-backend/pkg/http/controllers"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on system env")
	}
}

func main() {
	h := controllers.NewHandler()
	srv := geomanjihttp.NewServer(":8080", h)

	log.Printf("listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
