package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	geomanjihttp "github.com/vijayvenkatj/geomanji-backend/pkg/http"
	"github.com/vijayvenkatj/geomanji-backend/pkg/http/controllers"
	"github.com/vijayvenkatj/geomanji-backend/pkg/integrations"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on system env")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h := controllers.NewHandler(integrations.NewGeomanjiClient())
	srv := geomanjihttp.NewServer(":"+port, h)

	log.Printf("listening on %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
