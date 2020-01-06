package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	locationHelper "obas/controllers/location"
	academicsDomain "obas/domain/academics"
	institutionDomain "obas/domain/institutions"
	locationDomain "obas/domain/location"
	"obas/io/academics"
	"obas/io/address"
	institutionIO "obas/io/institutions"
	location "obas/io/location"
)

func InstitutionManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", InstitutionManagementHandler(app))
	r.Get("/type/delete/{formData}", DeleteTypeHandler(app))
	r.Get("/delete/institution/{institutionId}", DeleteInstitutionHandler(app))
	r.Get("/delete-institutionLocation/{institutionLocationId}", DeleteInstitutionLocationHandler(app))
	r.Get("/delete-institutionCourse/{institutionId}/{CourseId}", DeleteInstitutionCourseHandler(app))
	r.Get("/delete-institutionAddress/{InstitutionAddressId}/{AddressTypeId}", DeleteInstitutionAddressHandler(app))
	r.Post("/type/add", AddInstitutiontypeHandler(app))
	r.Post("/type/edit", EditTypeHandler(app))
	r.Post("/add", AddInstitutionHandler(app))
	r.Post("/save/institution-location", SaveInstitutionLocationHandler(app))
	r.Post("/save/institution-course", SaveInstitutionCourseHandler(app))
	r.Post("/save/institution-address", SaveInstitutionAddressHandler(app))
	return r
}

func DeleteInstitutionAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		InstitutionAddressId := chi.URLParam(r, "InstitutionAddressId")
		AddressTypeId := chi.URLParam(r, "AddressTypeId")
		_ = app.Session.Destroy(r.Context())

		if InstitutionAddressId != "" {
			institutionAddressObject, err := institutionIO.ReadInstitutionAddress(InstitutionAddressId, AddressTypeId)
			if err != nil {
				fmt.Println("an error in DeleteInstitutionAddressHandler")
				app.ErrorLog.Println(err.Error())
			}
			if institutionAddressObject.InstitutionId != "" {
				_, err := institutionIO.DeleteInstitutionAddress(institutionAddressObject)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				}
			}
		}
		app.Session.Put(r.Context(), "tab", "tab4")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func SaveInstitutionAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		addressType := r.PostFormValue("addressType")
		institutionId := r.PostFormValue("institutionAddress")
		postalCode := r.PostFormValue("postalCode")
		address := r.PostFormValue("address")

		_ = app.Session.Destroy(r.Context())

		if addressType != "" || institutionId != "" || postalCode != "" || address != "" {
			institutionAddress := institutionDomain.InstitutionAddress{institutionId, addressType, address, postalCode}
			_, err := institutionIO.CreateInstitutionAddress(institutionAddress)
			if err != nil {
				fmt.Println("error in creating institutionAddress")
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab4")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func DeleteInstitutionCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		institutionId := chi.URLParam(r, "institutionId")
		CourseId := chi.URLParam(r, "CourseId")
		_ = app.Session.Destroy(r.Context())

		if CourseId != "" || institutionId != "" {
			institutionCourserObject, err := institutionIO.ReadInstitutionCourse(institutionId, CourseId)
			if err != nil {
				fmt.Println("an error in DeleteInstitutionCourseHandler")
				app.ErrorLog.Println(err.Error())
			}
			if institutionCourserObject.InstitutionId != "" {
				_, err := institutionIO.DeleteInstitutionCourse(institutionCourserObject)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				}
			}
		}
		app.Session.Put(r.Context(), "tab", "tab5")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func SaveInstitutionCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		institutionId := r.PostFormValue("institutionCourseDrop")
		courseId := r.PostFormValue("courseId")

		_ = app.Session.Destroy(r.Context())

		if courseId != "" || institutionId != "" {
			institutionCourse := institutionDomain.InstitutionCourse{institutionId, courseId}
			_, err := institutionIO.CreateInstitutionCourse(institutionCourse)
			if err != nil {
				fmt.Println("error in creating institutionCourse")
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab5")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

type Tabs struct {
	Tab1 string
	Tab2 string
	Tab3 string
	Tab4 string
	Tab5 string
}

func DeleteInstitutionLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "institutionLocationId")
		_ = app.Session.Destroy(r.Context())

		if resetKey != "" {
			institutionLocationObject, err := institutionIO.ReadInstitutionLocation(resetKey)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
			if institutionLocationObject.InstitutionId != "" {
				_, err := institutionIO.DeleteInstitutionLocation(institutionLocationObject)
				if err != nil {
					app.ErrorLog.Println(err.Error())
				}
			}
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func SaveInstitutionLocationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		locationId := r.PostFormValue("townId")
		institutionId := r.PostFormValue("institution")
		longitude := r.PostFormValue("longitude")
		latitude := r.PostFormValue("latitude")

		_ = app.Session.Destroy(r.Context())

		if locationId != "" || institutionId != "" || longitude != "" || latitude != "" {
			institutionLocation := institutionDomain.InstitutionLocation{institutionId, locationId, longitude, latitude}
			_, err := institutionIO.CreateInstitutionLocation(institutionLocation)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab3")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func DeleteInstitutionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "institutionId")
		_ = app.Session.Destroy(r.Context())
		institutionObject, err := institutionIO.GetInstitution(resetKey)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if institutionObject.Id != "" {
			_, err := institutionIO.DeleteInstitution(institutionObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func AddInstitutionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		institutionName := r.PostFormValue("institutionName")
		institutionType := r.PostFormValue("institutionType")
		_ = app.Session.Destroy(r.Context())
		fmt.Println(institutionName, "  <<<<institutionName    institutionType", institutionType)
		if institutionName != "" || institutionType != "" {
			institu := institutionDomain.Institution{"", institutionType, institutionName}
			_, err := institutionIO.CreateInstitution(institu)
			if err != nil {
				fmt.Println("error creating institution")
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
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

		_ = app.Session.Destroy(r.Context())

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
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}

func DeleteTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resetKey := chi.URLParam(r, "formData")

		_ = app.Session.Destroy(r.Context())

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
		app.Session.Put(r.Context(), "tab", "tab1")
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

		_ = app.Session.Destroy(r.Context())

		institution := institutionDomain.InstitutionTypes{"", institutionName, institutionDescription}
		app.InfoLog.Println("Institution to save: ", institution)
		_, err := institutionIO.CreateInstitutionType(institution)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/institution", 301)
	}
}
func GetTabs(tab string) Tabs {

	switch tab {
	case "tab1":
		return Tabs{"active show", "", "", "", ""}
	case "tab2":
		return Tabs{"", "active show", "", "", ""}
	case "tab3":
		return Tabs{"", "", "active show", "", ""}
	case "tab4":
		return Tabs{"", "", "", "active show", ""}
	case "tab5":
		return Tabs{"", "", "", "", "active show"}
	default:
		return Tabs{"active show", "", "", "", ""}
	}
}

type InstitutionCourseHolder struct {
	InstitutionId     string
	CourseId          string
	InstitutionNane   string
	CourseName        string
	CourseDescription string
}

type InstitutionAddressHolder struct {
	InstitutionAddressId string
	AddressTypeId        string
	Institution          string
	Address              string
	Postal               string
}

func InstitutionManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tab := app.Session.GetString(r.Context(), "tab")

		activeTab := GetTabs(tab)
		type InstitutionHolder struct {
			Id              string
			InstitutionName string
			InstitutionType string
		}
		type InstitutionLocHolder struct {
			Id          string
			Institution string
			Town        string
			Longitude   string
			Latitude    string
		}

		var institutionLocation []InstitutionLocHolder
		var institutions []institutionDomain.Institution
		var institutionsHolder []InstitutionHolder
		var institutionsCourseHolder []InstitutionCourseHolder
		var institutionsAddressHolder []InstitutionAddressHolder

		institutionAddresses, err := institutionIO.GetInstitutionAddresses()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else if institutionAddresses != nil {
			for _, institutionAddrres := range institutionAddresses {
				myInstitution, err := institutionIO.GetInstitution(institutionAddrres.InstitutionId)
				if err != nil {
					fmt.Println("An error in InstitutionManagementHandler when reading myInstitution")
					app.ErrorLog.Println(err.Error())
				} else if myInstitution.Name != "" {
					institutionAddress := InstitutionAddressHolder{institutionAddrres.AddressTypeId, institutionAddrres.InstitutionId, myInstitution.Name, institutionAddrres.Address, institutionAddrres.PostalCode}
					institutionsAddressHolder = append(institutionsAddressHolder, institutionAddress)
				}

			}
		}

		allInstitutionCourse, err := institutionIO.GetAllInstitutionCourse()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			for _, institutionCourse := range allInstitutionCourse {
				institution, err := institutionIO.GetInstitution(institutionCourse.InstitutionId)
				if err != nil {
					fmt.Println("error reading institution in InstitutionManagementHandler method")
					app.ErrorLog.Println(err.Error())
				}
				couse, err := academics.GetCourse(institutionCourse.CourseId)
				if err != nil {
					fmt.Println("error reading institution in InstitutionManagementHandler method")
					app.ErrorLog.Println(err.Error())
				}
				if institution.Name != "" || couse.CourseName != "" {
					myInstitutionCours := InstitutionCourseHolder{institutionCourse.InstitutionId, institutionCourse.CourseId, institution.Name, couse.CourseName, couse.CourseDesc}
					institutionsCourseHolder = append(institutionsCourseHolder, myInstitutionCours)
				}
			}
		}

		institutsLocation, err := institutionIO.ReadInstitutionLocations()
		if err != nil {
			app.ErrorLog.Println(err.Error())
		} else {
			for _, institutionLoca := range institutsLocation {
				fmt.Println("error in reading townNamw in InstitutionManagementHandler method", institutionLoca.LocationId)

				institutionName, errr := institutionIO.GetInstitution(institutionLoca.InstitutionId)
				if errr != nil {
					fmt.Println("error in reading institutionName in InstitutionManagementHandler method")
					app.ErrorLog.Println(errr.Error())
				}
				townNamw, err := location.GetLocation(institutionLoca.LocationId)
				fmt.Println("error in reading townNamw in InstitutionManagementHandler method", townNamw)
				if err != nil {
					fmt.Println("error in reading townNamw in InstitutionManagementHandler method")
					app.ErrorLog.Println(err.Error())
				}
				institutionLocation = append(institutionLocation, InstitutionLocHolder{institutionLoca.InstitutionId, institutionName.Name, townNamw.Name, institutionLoca.Longitude, institutionLoca.Latitude})
			}
		}

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
					institutionsHolder = append(institutionsHolder, InstitutionHolder{institution.Id, institution.Name, institutionTypeName})
				}
			}
		}
		provinces, _ := locationHelper.GetProvinces(app)
		courses, err := academics.GetAllCourses()
		if err != nil {
			fmt.Println("error in InstitutionManagementHandler when reading courses")
			app.ErrorLog.Println(err.Error())
		}
		addressTypes, err := address.GetAddressTypes()
		if err != nil {
			fmt.Println("error in InstitutionManagementHandler when reading addressTypes")
			app.ErrorLog.Println(err.Error())
		}

		type PageData struct {
			InstitutionTypes    []institutionDomain.InstitutionTypes
			InstitutionsHolder  []InstitutionHolder
			Provinces           []locationDomain.Location
			InstitutionLocation []InstitutionLocHolder
			MyActiveTab         Tabs
			Courses             []academicsDomain.Course
			InstitutionCourse   []InstitutionCourseHolder
			AddressTypes        []address.AddressType
			InstitutionAddress  []InstitutionAddressHolder
		}
		data := PageData{institutionTypes, institutionsHolder, provinces, institutionLocation, activeTab, courses, institutionsCourseHolder, addressTypes, institutionsAddressHolder}

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
