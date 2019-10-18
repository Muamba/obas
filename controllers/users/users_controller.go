package controllers

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	addressIO "obas/io/address"
	demograhpyIO "obas/io/demographics"
	usersIO "obas/io/users"
	"strings"
	"time"
)

const (
	layoutOBAS        = "2006-01-02"
	dangerAlertStyle  = "alert-danger"
	successAlertStyle = "alert-success"
)

type AddressPlaceHolder struct {
	AddressName string
	Address     string
	PostalCode  string
}

type PageToast struct {
	AlertType string
	AlertInfo string
}

func Users(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", UsersHandler(app))
	r.Get("/admin", AdminHandler(app))
	r.Get("/student", StudentHandler(app))

	r.Get("/processingStatus", ProcessingStatusTypeHandler(app))
	r.Get("/student/application", StudentApplicationStatusHandler(app))
	r.Get("/studentContact", StudentContactsHandler(app))
	r.Get("/student/documents", StudentDocumentsHandler(app))
	r.Get("/studentResults", StudentResultsHandler(app))

	r.Get("/student/profile/personal", StudentProfilePersonalHandler(app))
	r.Get("/student/profile/demography", StudentProfileDemographyHandler(app))
	r.Get("/student/profile/address", StudentProfileAddressHandler(app))
	r.Get("/student/profile/relative", StudentProfileRelativeHandler(app))
	r.Get("/student/profile/settings", StudentProfileRegistrationHandler(app))
	r.Get("/student/profile/courses", StudentProfileCourseHandler(app))
	r.Get("/student/profile/subjects", StudentProfileSubjectHandler(app))
	r.Get("/student/profile/districts", StudentProfileDistrictHandler(app))

	r.Post("/student/profile/personal/update", UpdateStudentProfilePersonalHandler(app))
	r.Post("/student/profile/address/addresstype", StudentProfileAddressTypeHandler(app))
	r.Post("/student/profile/address/update", StudentProfileAddressUpdateHandler(app))
	r.Post("/student-profile-relative-upate", StudentProfileRelativeUpdateHandler(app))
	r.Post("/student-profile-demography-update", StudentProfileDemographyUpdateHandler(app))

	return r
}

func StudentProfileDemographyUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		title := r.PostFormValue("title")
		gender := r.PostFormValue("gender")
		race := r.PostFormValue("race")
		userDemograpgy := usersIO.UserDemography{email, title, gender, race}
		app.InfoLog.Println("userDemography to update: ", userDemograpgy)

		updated, err := usersIO.UpdateUserDemographics(userDemograpgy)
		successMessage := "User demography updated!"
		failureMessage := "User demography NOT Updated!"

		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, failureMessage)
		} else {
			if updated {
				setSessionMessage(app, r, successAlertStyle, successMessage)
			} else {
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("UserDemography update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/demography", 301)
	}
}
func StudentProfileDemographyHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		var alert PageToast
		var genders []demograhpyIO.Gender
		var races []demograhpyIO.Race
		var gender demograhpyIO.Gender
		var race demograhpyIO.Race
		var title demograhpyIO.Title
		titles, err := demograhpyIO.GetTitles()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve titles!"}
		} else {
			genders, err = demograhpyIO.GetGenders()
			if err != nil {
				app.ErrorLog.Println(err.Error())
				alert = PageToast{dangerAlertStyle, "Could not retrieve genders!"}
			} else {
				races, err = demograhpyIO.GetRaces()
				if err != nil {
					app.ErrorLog.Println(err.Error())
					alert = PageToast{dangerAlertStyle, "Could not retrieve races!"}
				} else {
					userDemography, err := usersIO.GetUserDemographic(email)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						alert = PageToast{dangerAlertStyle, "Could not retrieve student demography!"}
					} else {
						message := app.Session.GetString(r.Context(), "message")
						messageType := app.Session.GetString(r.Context(), "message-type")
						if message != "" && messageType != "" {
							alert = PageToast{messageType, message}
							app.Session.Remove(r.Context(), "message")
							app.Session.Remove(r.Context(), "message-type")
						}
						title = getUserTitle(userDemography, titles)
						gender = getUserGender(userDemography, genders)
						race = getUserRace(userDemography, races)
					}
				}
			}
		}

		type PageData struct {
			Student       usersIO.User
			Titles        []demograhpyIO.Title
			Genders       []demograhpyIO.Gender
			Races         []demograhpyIO.Race
			Alert         PageToast
			StudentTitle  demograhpyIO.Title
			StudentGender demograhpyIO.Gender
			StudentRace   demograhpyIO.Race
		}

		data := PageData{user, titles, genders, races, alert, title, gender, race}
		app.InfoLog.Println("PageData: ", data)
		files := []string{
			app.Path + "content/student/profile/demography.html",
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

func getUserRace(demography usersIO.UserDemography, races []demograhpyIO.Race) demograhpyIO.Race {
	for _, race := range races {
		if demography.RaceId == race.RaceId {
			return race
		}
	}
	return demograhpyIO.Race{}
}

func getUserTitle(demography usersIO.UserDemography, titles []demograhpyIO.Title) demograhpyIO.Title {
	for _, title := range titles {
		if demography.TitleId == title.TitleId {
			return title
		}
	}
	return demograhpyIO.Title{}
}

func getUserGender(demography usersIO.UserDemography, genders []demograhpyIO.Gender) demograhpyIO.Gender {
	for _, gender := range genders {
		if demography.GenderId == gender.GenderId {
			return gender
		}
	}
	return demograhpyIO.Gender{}
}

func StudentProfileRelativeUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		relativeName := r.PostFormValue("relative_name")
		relationship := r.PostFormValue("relationship")
		cellphone := r.PostFormValue("relative_cellphone")
		relativeEmail := r.PostFormValue("relative_email")
		userRelative := usersIO.UserRelative{email, relativeName, cellphone, relativeEmail, relationship}
		app.InfoLog.Println("UserRelative to update: ", userRelative)
		updated, err := usersIO.UpdateUserRelative(userRelative, token)

		successMessage := "User relative updated!"
		failureMessage := "User relative NOT Updated!"

		if err != nil {
			app.ErrorLog.Println(err.Error())
			setSessionMessage(app, r, dangerAlertStyle, failureMessage)
		} else {
			if updated {
				setSessionMessage(app, r, successAlertStyle, successMessage)
			} else {
				setSessionMessage(app, r, dangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("UserRelative update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/relative", 301)

	}
}

func setSessionMessage(app *config.Env, r *http.Request, messageType string, message string) {
	app.Session.Put(r.Context(), "message-type", messageType)
	app.Session.Put(r.Context(), "message", message)
}

func StudentProfileSubjectHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/subjects.html",
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

func StudentProfileRegistrationHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/settings.html",
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

func StudentProfileCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/courses.html",
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

func StudentProfileDistrictHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		type PageData struct {
			Student usersIO.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/district_and_municipality.html",
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

func StudentProfileRelativeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		var alert PageToast

		userRelative, err := usersIO.GetUserRelative(user.Email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = PageToast{dangerAlertStyle, "Could not retrieve student relative!"}
		} else {
			message := app.Session.GetString(r.Context(), "message")
			messageType := app.Session.GetString(r.Context(), "message-type")
			if message != "" && messageType != "" {
				alert = PageToast{messageType, message}
				app.Session.Remove(r.Context(), "message")
				app.Session.Remove(r.Context(), "message-type")
			}
		}

		type PageData struct {
			Student         usersIO.User
			StudentRelative usersIO.UserRelative
			Alert           PageToast
		}

		data := PageData{user, userRelative, alert}
		files := []string{
			app.Path + "content/student/profile/relative.html",
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

func StudentProfileAddressUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		addressTypeId := r.PostFormValue("addressTypeId")
		address := r.PostFormValue("address")
		postalCode := r.PostFormValue("postalCode")
		userAddress := usersIO.UserAddress{email, addressTypeId, address, postalCode}
		app.InfoLog.Println("UserAddress to update: ", userAddress)
		updated, err := usersIO.UpdateUserAddress(userAddress, token)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/address", 301)
	}
}

func StudentProfileAddressHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		addressTypes, err := addressIO.GetAddressTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), addressTypes)
		}

		addresses := []AddressPlaceHolder{}

		for _, addressType := range addressTypes {
			userAddress, err := usersIO.GetUserAddress(email, addressType.AddressTypeID)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				addresses = append(addresses, AddressPlaceHolder{addressType.AddressName, userAddress.Address, userAddress.PostalCode})
			}
		}

		type PageData struct {
			Student       usersIO.User
			AddressTypes  []addressIO.AddressType
			Addresses     []AddressPlaceHolder
			Address       usersIO.UserAddress
			AddressTypeID string
			AddressName   string
		}

		data := PageData{user, addressTypes, addresses, usersIO.UserAddress{}, "", ""}
		files := []string{
			app.Path + "content/student/profile/address.html",
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

func StudentProfileAddressTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		addressTypeId := r.PostFormValue("addresstypes")
		userAddress, err := usersIO.GetUserAddress(email, addressTypeId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

		addressTypes, err := addressIO.GetAddressTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error(), addressTypes)
		}

		addresses := []AddressPlaceHolder{}
		var addressName string

		for _, addressType := range addressTypes {
			if addressTypeId == addressType.AddressTypeID {
				addressName = addressType.AddressName
			}
			userAddress, err := usersIO.GetUserAddress(email, addressType.AddressTypeID)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			} else {
				addresses = append(addresses, AddressPlaceHolder{addressType.AddressName, userAddress.Address, userAddress.PostalCode})
			}
		}

		type PageData struct {
			Student       usersIO.User
			AddressTypes  []addressIO.AddressType
			Addresses     []AddressPlaceHolder
			Address       usersIO.UserAddress
			AddressTypeID string
			AddressName   string
		}

		data := PageData{user, addressTypes, addresses, userAddress, addressTypeId, addressName}
		files := []string{
			app.Path + "content/student/profile/address.html",
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

func StudentHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" || len(email) <= 0 {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		type PageData struct {
			Student usersIO.User
		}
		data := PageData{user}
		files := []string{
			app.Path + "content/student/student_dashboard.page.html",
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

func StudentProfilePersonalHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		if email == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}
		dobString := strings.Split(user.DateOfBirth.String(), " ")[0] // split date and get in format: yyy-mm-dd

		type PageData struct {
			Student     usersIO.User
			DateOfBirth string
		}

		data := PageData{user, dobString}
		files := []string{
			app.Path + "content/student/profile/personal.html",
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

func UpdateStudentProfilePersonalHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		idNumber := r.PostFormValue("id_number")
		firstName := r.PostFormValue("first_name")
		lastName := r.PostFormValue("last_name")
		dateOfBirthStr := r.PostFormValue("dateOfBirth")
		dateOfBirth, _ := time.Parse(layoutOBAS, dateOfBirthStr)
		user := usersIO.User{email, idNumber, firstName, "", lastName, dateOfBirth}
		app.InfoLog.Println("User to update: ", user)
		updated, err := usersIO.UpdateUser(user, token)

		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/personal", 301)
	}
}

func UsersHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		type PageData struct {
			//courses []io.Users
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "base/register/register.page.html",
			/**app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.page.html",*/
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func AdminHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allAdmin, err := io.GetAdmins()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//courses []io.Admin
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func ProcessingStatusTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allProcess, err := io.GetProcessingStatusTypes()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//subjects []io.ProcessingStatusType
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func StudentApplicationStatusHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "content/student/Student_Application.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())

		}
	}
}

func StudentContactsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allStudentContacts, err := io.GetStudentContacts()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//subjects []io.StudentContacts
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}

func StudentDocumentsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{

			app.Path + "content/student/Student_Documents.html",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())

		}

	}
}

func StudentResultsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allStudentResults, err := io.GetStudentResults()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//subjects []io.StudentResults
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.page.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.ExecuteTemplate(w, "base", data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}

	}
}
