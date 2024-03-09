package ping

import (
	"github.com/go-chi/chi/v5"
	"lobinv-server/api/middleware"
	"net/http"
)

func Router(r chi.Router) {
	r.Use(middleware.ContentTypeText)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Pong!"))
	})
}
