package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"obas/config"
	institutionDomain "obas/domain/institutions"
	institutionIO "obas/io/institutions"
)

func InstitutionAPI(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/getInstitutionsInLocation/{locationId}", GetInstitutionsInLocationHandler(app))
	r.Get("/getInstitutionsByType/{institutionTypeId}", GetInstitutionByTypeHandler(app))
	r.Get("/getInstitutionsByTypenLocation/{institutionTypeId}/{locationId}", GetInstitutionsByTypenLocationHandler(app))

	return r
}

func GetInstitutionsByTypenLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		institutionTypeId := chi.URLParam(r, "institutionTypeId")
		locationId := chi.URLParam(r, "locationId")
		var institutions []institutionDomain.Institution
		institutionLocations := getInstitutionsByLocation(locationId, app)
		if len(institutionLocations) > 0 {
			for _, institutionLocation := range institutionLocations {
				institution, err := institutionIO.GetInstitution(institutionLocation.InstitutionId)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				} else {
					if institution.InstitutionTypeId == institutionTypeId {
						institutions = append(institutions, institution)
					}
				}
			}
		}
		render.JSON(w, r, institutions)
	}
}

func GetInstitutionByTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		institutionTypeId := chi.URLParam(r, "institutionTypeId")
		institutions, err := institutionIO.GetInstitutionsByType(institutionTypeId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		render.JSON(w, r, institutions)
	}
}

func GetInstitutionsInLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locationId := chi.URLParam(r, "locationId")
		var institutions []institutionDomain.Institution
		institutionLocations := getInstitutionsByLocation(locationId, app)
		if len(institutionLocations) > 0 {
			for _, institutionLocation := range institutionLocations {
				institution, err := institutionIO.GetInstitution(institutionLocation.InstitutionId)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				} else {
					institutions = append(institutions, institution)
				}
			}
		}
		render.JSON(w, r, institutions)
	}
}

func getInstitutionsByLocation(locationId string, app *config.Env) []institutionDomain.InstitutionLocation {
	var institutionLocations []institutionDomain.InstitutionLocation
	institutionLocations, err := institutionIO.GetInstitutionsInLocation(locationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
	}
	return institutionLocations
}
