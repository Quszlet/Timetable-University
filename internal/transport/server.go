package server

import (
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func Init(r *mux.Router) (server *http.Server) {
	return &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}