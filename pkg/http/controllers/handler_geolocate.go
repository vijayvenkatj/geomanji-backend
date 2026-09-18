package controllers

import "net/http"

// Handler is the parent struct for all HTTP handlers.
// Add dependencies (DB, clients, config, etc.) here as fields.
type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Geolocate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("geolocate"))
}
