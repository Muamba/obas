package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	domain "obas/domain/location"
	locationIO "obas/io/location"
)

func LocationManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", LocationManagementHandler(app))

	r.Post("/type/add", AddLocationTypeHandler(app))
	r.Post("/add", AddLocationHandler(app))
	r.Post("/type/edit", EditLocationTypeHandler(app))
	r.Get("/delete/location/{resetkey}", DeleteLocationHandler(app))
	return r
}

func EditLocationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())

		r.ParseForm()
		locationId := r.PostFormValue("LocationId")
		locationName := r.PostFormValue("Name")
		latitude := r.PostFormValue("Latitude")
		longitude := r.PostFormValue("Longitude")
		locationType := r.PostFormValue("locationType")
		locationParent := r.PostFormValue("locationParent")
		location := domain.Location{locationId, locationType, locationName, latitude, longitude, locationParent}

		fmt.Println(location)
		_, err := locationIO.UpdateLocation(location)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func DeleteLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		locationId := chi.URLParam(r, "resetkey")
		_ = app.Session.Destroy(r.Context())
		fmt.Print(locationId)
		locationObject, err := locationIO.GetLocation(locationId)
		fmt.Print(locationObject)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if locationObject.LocationId != "" {
			_, err := locationIO.DeleteLocation(locationObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func AddLocationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())
		/***
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		*/

		r.ParseForm()
		locationTypeName := r.PostFormValue("locationTypeName")
		locationType := domain.LocationType{"", locationTypeName, ""}
		app.InfoLog.Println("Location type to save: ", locationType)
		savedLocationType, err := locationIO.CreateLocationType(locationType)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}

		app.Session.Put(r.Context(), "tab", "tab1")

		app.InfoLog.Println("Create location type response is ", savedLocationType)
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func AddLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())
		//email := app.Session.GetString(r.Context(), "userId")
		//token := app.Session.GetString(r.Context(), "token")
		//if email == "" || token == "" {
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		r.ParseForm()
		locationName := r.PostFormValue("locationName")
		latitude := r.PostFormValue("Latitude")
		longitude := r.PostFormValue("Longitude")
		locationType := r.PostFormValue("LocationType")
		locationParent := r.PostFormValue("LocationParent")
		location := domain.Location{"", locationType, locationName, latitude, longitude, locationParent}
		app.InfoLog.Println("Location to save: ", location)
		savedLocation, err := locationIO.CreateLocation(location)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		app.InfoLog.Println("Create location response is ", savedLocation)
		http.Redirect(w, r, "/support/management/location", 301)
	}
}

func LocationManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tab1 string
		var tab2 string
		tab := app.Session.GetString(r.Context(), "tab")

		if tab == "tab1" {
			tab1 = "active show"
			tab2 = ""
		} else if tab == "tab2" {
			tab2 = "active show"
			tab1 = ""
		} else {
			tab1 = "active show"
			tab2 = ""
		}

		var locations []domain.Location
		locationTypes, err := locationIO.GetLocationTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			locations, err = locationIO.GetLocations()
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}

		type PageData struct {
			LocationTypes []domain.LocationType
			Locations     []domain.Location
			Tab1          string
			Tab2          string
		}

		data := PageData{locationTypes, locations, tab1, tab2}

		files := []string{
			app.Path + "content/tech/tech_admin_loc.html",
			app.Path + "content/tech/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
