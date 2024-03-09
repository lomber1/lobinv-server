package units

import (
	"github.com/go-chi/chi/v5"
)

func Router(r chi.Router) {
	r.Get("/", Get)
	r.Post("/", Post)
	r.Put("/{unitID}", Put)
	r.Delete("/{unitID}", Delete)
}
