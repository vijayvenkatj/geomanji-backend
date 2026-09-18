package http

import (
	"net/http"

	"github.com/vijayvenkatj/geomanji-backend/pkg/http/controllers"
)

func NewRouter(h *controllers.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /geolocate", h.Geolocate)

	return mux
}
