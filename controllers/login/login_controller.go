package login

import (
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	"obas/io/login"
)

// Route Path
func Login(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", loginHome(app))
	r.Get("/error", loginError(app))
	r.Get("/resetError", resetError(app))
	r.Post("/login", loginHandler(app))
	r.Post("/reset", passwordResetHandler(app))
	r.Post("/forgotPassword", forgotpasswordHandler(app))
	r.Post("/accounts", getAccountsHandler(app))
	r.Get("/password", passwordHandler(app))
	r.Get("/verify", passwordHandler(app))
	return r
}

func loginHome(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "base/login/login.page.html",
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

func loginError(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "base/login/login.page_Error.html",
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

func resetError(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files := []string{
			app.Path + "base/login/forgotpassword.page.html",
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

func passwordResetHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("password")
		result, err := login.DoRest(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login/resetError", 301)
		}
		app.InfoLog.Println("Login is successful. Result is ", result)
		http.Redirect(w, r, "/login", 301)
	}
}
func forgotpasswordHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		result, err := login.DoForgetPassword(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login/resetError", 301)
		}
		app.InfoLog.Println("Login is successful. Result is ", result)
		http.Redirect(w, r, "/login", 301)
	}
}
func loginHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		loginToken, err := login.DoLogin(email, password)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login/error", 301)
		}
		app.InfoLog.Println("Login is successful. Result is ", loginToken)
		http.Redirect(w, r, "/users/student", 301)
	}
}

func logout(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))

	}
}

func forgotPassword(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))

	}
}

func passwordHandler(app *config.Env) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		files := []string{
			app.Path + "base/login/passwordReset.page.html",
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

func getAccountsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
