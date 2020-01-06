package institutions

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/config"
	"obas/controllers/institutions/api"
)

func Institutions(app *config.Env) http.Handler {
	mux := chi.NewMux()

	mux.Mount("/api", api.InstitutionAPI(app))

	return mux
}