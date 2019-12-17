package location

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/config"
	"obas/controllers/location/api"
)

func Location(app *config.Env) http.Handler {
	mux := chi.NewMux()

	mux.Mount("/api", api.LocationAPI(app))

	return mux
}
