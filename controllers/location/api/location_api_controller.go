package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"obas/config"
	locationDomain "obas/domain/location"
	locationIO "obas/io/location"
)

func LocationAPI(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/getforparent/{locationParentId}", GetLocationsForParent(app))
	r.Get("/getinstitutions/{institutionTypeId}", GetInstitutions(app))

	return r
}

func GetInstitutions(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/**to do on monday if all goes well thanks lord for the week may god remain the whole mighty**/
		//institutionType := chi.URLParam(r, "locationParentId")
		//var institutions []institutionDomain.Institution
		//locations, err := locationIO.GetLocationsForParent(locationParentId)
		//if err != nil {
		//	app.ErrorLog.Println(err.Error())
		//}
	}
}

func GetLocationsForParent(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locationParentId := chi.URLParam(r, "locationParentId")
		var locations []locationDomain.Location
		locations, err := locationIO.GetLocationsForParent(locationParentId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		render.JSON(w, r, locations)
	}
}
