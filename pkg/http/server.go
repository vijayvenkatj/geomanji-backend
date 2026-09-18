package http

import (
	"net/http"

	"github.com/vijayvenkatj/geomanji-backend/pkg/http/controllers"
)

func NewServer(addr string, h *controllers.Handler) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: NewRouter(h),
	}
}
