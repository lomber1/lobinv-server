package resources

import (
	"github.com/go-chi/chi/v5"
	"lobinv-server/api/resources/ping"
	"lobinv-server/api/resources/units"
)

func Router(r chi.Router) {
	r.Route("/ping", ping.Router)
	r.Route("/units", units.Router)
}
