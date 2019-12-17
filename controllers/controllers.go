package controllers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"obas/config"
	academicsController "obas/controllers/academics"
	addressController "obas/controllers/address"
	administrationController "obas/controllers/administration"
	applicationController "obas/controllers/application"
	demographicsController "obas/controllers/demographics"
	documentsController "obas/controllers/documents"
	homeController "obas/controllers/home"
	institutionsController "obas/controllers/institutions"
	locationController "obas/controllers/location"
	logController "obas/controllers/log"
	loginController "obas/controllers/login"
	logoutController "obas/controllers/logout"
	registerController "obas/controllers/register"
	supportControllers "obas/controllers/support"
	usersController "obas/controllers/users"
)

func Controllers(env *config.Env) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	mux.Use(env.Session.LoadAndSave)

	mux.Handle("/", homeController.Home(env))
	mux.Mount("/login", loginController.Login(env))
	mux.Mount("/logout", logoutController.Logout(env))
	mux.Mount("/register", registerController.Register(env))
	mux.Mount("/users", usersController.Users(env))
	mux.Mount("/academics", academicsController.Academics(env))
	mux.Mount("/address", addressController.Addresses(env))
	mux.Mount("/demographics", demographicsController.Demographics(env))
	mux.Mount("/application", applicationController.Applications(env))
	mux.Mount("/documents", documentsController.Documents(env))
	mux.Mount("/institution", institutionsController.Institutions(env))
	mux.Mount("/log", logController.Logs(env))
	mux.Mount("/support", supportControllers.Support(env))
	mux.Mount("/admin", administrationController.Admin(env))
	mux.Mount("/location", locationController.Location(env))

	fileServer := http.FileServer(http.Dir("./views/assets/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/assets/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Mount("/assets/", http.StripPrefix("/assets", fileServer))
	return mux

}
