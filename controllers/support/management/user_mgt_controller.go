package management

import (
	"fmt"
	"github.com/go-chi/chi"
	"html/template"
	"net/http"
	"obas/config"
	userDomain "obas/domain/users"
	"obas/io/demographics"
	userIO "obas/io/users"
)

func UserManagement(app *config.Env) http.Handler {
	r := chi.NewRouter()

	r.Get("/", UserManagementHandler(app))
	r.Get("/role/delete/{roleId}", DeleteRoleManagementHandler(app))
	r.Post("/role/create", RoleCreateManagementHandler(app))
	r.Post("/role/update", RoleUpdateManagementHandler(app))

	return r
}

func DeleteRoleManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleId := chi.URLParam(r, "roleId")
		_ = app.Session.Destroy(r.Context())
		roleeObject, err := demographics.GetRole(roleId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
		}
		if roleeObject.Id != "" {
			_, err := demographics.DeleteRole(roleeObject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
			}
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/user", 301)
	}
}

func RoleCreateManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())
		r.ParseForm()
		role := r.PostFormValue("role")

		if role != "" {
			//newRole:=domain.Role{"",role}
			newRole := demographics.Role{"", role}
			_, err := demographics.CreateRole(newRole)
			if err != nil {
				fmt.Println("error creating Role")
			}
		}
		app.Session.Put(r.Context(), "tab", "tab1")
		http.Redirect(w, r, "/support/management/user", 301)
	}
}

func RoleUpdateManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = app.Session.Destroy(r.Context())
		r.ParseForm()
		actualUser := r.PostFormValue("userId")
		actualRole := r.PostFormValue("newRole")
		fmt.Println(actualUser, "<<<<<<actualUser||actualRole>>>>>", actualRole)
		if actualRole != "" || actualUser != "" {
			userRole := userDomain.UserRole{actualUser, actualRole}
			//_,err:=userIO.UpdateUserRole(userRole)
			_, err := userIO.DeleteUserRole(userRole)
			if err != nil {
				fmt.Println("error delete UserRole")
			} else {
				userIO.CreateUserRole(userRole)
				fmt.Println("about to update roles")
			}
		}
		app.Session.Put(r.Context(), "tab", "tab2")
		http.Redirect(w, r, "/support/management/user", 301)
	}

}

type myUserRole struct {
	UserId       string
	RoleId       string
	RoleName     string
	UserName     string
	UserSurname  string
	UserIdNumber string
	Roles        []demographics.Role
}

func getUserRole(userId string) userDomain.UserRole {
	entity := userDomain.UserRole{}
	user, err := userIO.GetUserRole(userId)
	if err != nil {
		return entity
	}
	return user
}
func getRole(roleId string) demographics.Role {
	entity := demographics.Role{}
	role, err := demographics.GetRole(roleId)
	if err != nil {
		return entity
	}
	return role
}
func UserManagementHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userRoleList []myUserRole
		users, err := userIO.GetUsers()
		if err != nil {
			fmt.Println("error reading Users")
		}
		roles, err := demographics.GetRoles()
		if err != nil {
			fmt.Println("error reading Roles")
		}
		userRoles, err := userIO.GetUserRoles()
		if err != nil {
			fmt.Println("error reading userRole")
		} else {
			for _, user := range users {
				myUser := getUserRole(user.Email)
				role := getRole(myUser.RoleId)
				userObject := myUserRole{user.Email, role.Id, role.RoleName, user.FirstName, user.LastName, user.IdNumber, roles}
				userRoleList = append(userRoleList, userObject)
			}
		}
		tab := app.Session.GetString(r.Context(), "tab")
		activeTab := myTabs(tab)
		type PageData struct {
			Users        []userDomain.User
			Roles        []demographics.Role
			UserRole     []userDomain.UserRole
			UserRoleList []myUserRole
			MyActiveTab  Roletabs
		}
		Data := PageData{users, roles, userRoles, userRoleList, activeTab}
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

type Roletabs struct {
	Tab1 string
	Tab2 string
}

func myTabs(tab string) Roletabs {
	switch tab {
	case "tab1":
		return Roletabs{"active show", ""}
	case "tab2":
		return Roletabs{"", "active show"}
	default:
		return Roletabs{"active show", ""}
	}
}
