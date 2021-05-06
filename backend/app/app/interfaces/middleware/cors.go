package middleware

import (
	"net/http"
	"todo-list/app/config"
)

type CORSMiddleware interface {
	Handler(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error)
}

type corsMiddleware struct {
	conf *config.Config
}

func NewCORSMiddleware(conf *config.Config) CORSMiddleware {
	return &corsMiddleware{conf}
}

func (m *corsMiddleware) Handler(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, error) {
	w.Header().Set("Access-Control-Allow-Origin", m.conf.Server.AllowOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	return w, r, nil
}