package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vijayvenkatj/geomanji-backend/pkg/integrations"
)

type Handler struct {
	Geomanji *integrations.GeomanjiClient
}

func NewHandler(geomanji *integrations.GeomanjiClient) *Handler {
	return &Handler{Geomanji: geomanji}
}

func (h *Handler) Geolocate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("geolocate"))
}

func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("img")
	if err != nil {
		http.Error(w, "missing img file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	result, err := h.Geomanji.Geolocate(header.Filename, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

type SearchMapsResult struct {
	integrations.GeolocateResult
	MapsLink string `json:"MapsLink"`
}

func (h *Handler) SearchMaps(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("img")
	if err != nil {
		http.Error(w, "missing img file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	result, err := h.Geomanji.Geolocate(header.Filename, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SearchMapsResult{
		GeolocateResult: *result,
		MapsLink:        integrations.GoogleMapsLink(result.Latitude, result.Longitude),
	})
}
