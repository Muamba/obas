package controllers

import (
	"github.com/go-chi/chi"
	"net/http"
	"obas/config"
	"obas/controllers/support/management"
)

func Support(app *config.Env) http.Handler {
	mux := chi.NewMux()

	mux.Handle("/", SupportHome(app))
	mux.Mount("/management/location", management.LocationManagement(app))
	mux.Mount("/management/institution", management.InstitutionManagement(app))
	mux.Mount("/management/user", management.UserManagement(app))
	mux.Mount("/management/academics", management.AcademicManagement(app))

	return mux
}
