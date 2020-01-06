package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	users2 "obas/io/users"
)

func UserManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", UserManagementHandler(app))
	r.Post("/role/update", RoleUpdateManagementHandler(app))

	return r
}

func RoleUpdateManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("about to update roles")
	}
}

func UserManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := users2.GetUsers()
		if err != nil {
			fmt.Println("error reading Users")
		}
		type PageData struct {
			Users []users2.User
		}

		Data := PageData{users}
		files := []string{
			app.Path + "content/tech/tech_admin_users.html",
			app.Path + "content/tech/template/sidebar.template.html",
			app.Path + "base/template/footer.template.html",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			return
		}
		err = ts.Execute(w, Data)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
	}
}
