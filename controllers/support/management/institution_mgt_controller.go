package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	institutionDomain "obas/domain/institutions"
	locationDomain "obas/domain/location"
	institutionIO "obas/io/institutions"
	"obas/util"
)

func InstitutionManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", InstitutionManagementHandler(app))
	r.Get("/type/delete/{formData}", DeleteTypeHandler(app))
	r.Post("/type/add", AddInstitutiontypeHandler(app))
	r.Post("/type/edit", EditTypeHandler(app))

	return r
}

func EditTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//email := app.Session.GetString(r.Context(), "userId")
		//token := app.Session.GetString(r.Context(), "token")
		//if email == "" || token == "" {
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		r.ParseForm()
		institutionTypeId := r.PostFormValue("Id")
		institutionTypeName := r.PostFormValue("Name")
		institutionTypeDescription := r.PostFormValue("Description")

		if institutionTypeId != "" || institutionTypeName != "" || institutionTypeDescription != "" {

			institutionTypeObject := institutionDomain.InstitutionTypes{institutionTypeId, institutionTypeName, institutionTypeDescription}
			fmt.Print(institutionTypeObject)
			_, err := institutionIO.DeleteInstitutionType(institutionTypeObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
			_, erro := institutionIO.CreateInstitutionType(institutionTypeObject)
			if erro != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func DeleteTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "formData")
		institutionObject, err := institutionIO.GetInstitutionType(resetKey)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if institutionObject.Id != "" {
			_, err := institutionIO.DeleteInstitutionType(institutionObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())

			}
		}
		http.Redirect(w, r, "/support/management/institution", 301)
	}

}

func AddInstitutiontypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//email := app.Session.GetString(r.Context(), "userId")
		//token := app.Session.GetString(r.Context(), "token")
		//if email == "" || token == "" {
		//	http.Redirect(w, r, "/login", 301)
		//	return
		//}
		r.ParseForm()
		institutionName := r.PostFormValue("Name")
		institutionDescription := r.PostFormValue("Description")

		institution := institutionDomain.InstitutionTypes{"", institutionName, institutionDescription}
		app.InfoLog.Println("Institution to save: ", institution)
		savedInstitutionTypes, err := institutionIO.CreateInstitutionType(institution)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Create location response is ", savedInstitutionTypes)
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func InstitutionManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type InstitutionHolder struct {
			InstitutionName string
			InstitutionType string
		}

		var institutions []institutionDomain.Institution
		var institutionsHolder []InstitutionHolder

		institutionTypes, err := institutionIO.GetInstitutionTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			institutions, err = institutionIO.GetInstitutions()
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				for _, institution := range institutions {
					institutionTypeName := getInstitutionTypeName(institution, institutionTypes)
					institutionsHolder = append(institutionsHolder, InstitutionHolder{institution.Name, institutionTypeName})
				}
			}
		}
		provinces, _ := util.GetProvinces()

		type PageData struct {
			InstitutionTypes   []institutionDomain.InstitutionTypes
			InstitutionsHolder []InstitutionHolder
			Provinces          []locationDomain.Location
		}

		data := PageData{institutionTypes, institutionsHolder, provinces}

		files := []string{
			app.Path + "content/tech/tech_admin_institution.html",
			app.Path + "content/tech/template/sidebar.template.html",
			app.Path + "base/template/form/location-form.template.html",
			app.Path + "base/template/form/institution-form.template.html",
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

func getInstitutionTypeName(institution institutionDomain.Institution, institutionTypes []institutionDomain.InstitutionTypes) string {
	var institutionTypeName string
	for _, institutionType := range institutionTypes {
		if institution.InstitutionTypeId == institutionType.Id {
			institutionTypeName = institutionType.Name
		}
	}
	return institutionTypeName
}
