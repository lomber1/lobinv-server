package api

import (
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"lobinv-server/api/middleware"
	"lobinv-server/api/resources"
	"time"
)

func NewMainRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(chiMiddleware.Timeout(5 * time.Second))
	r.Use(middleware.ContentTypeJSON)

	r.Route("/api", resources.Router)

	return r
}
