package users

import (
	"github.com/go-chi/chi"
	"html/template"
	"io"
	"net/http"
	"obas/config"
	institutionHelper "obas/controllers/institutions"
	locationHelper "obas/controllers/location"
	genericHelper "obas/controllers/misc"
	academicsDomain "obas/domain/academics"
	applicationDomain "obas/domain/application"
	documentDomain "obas/domain/documents"
	institutionDomain "obas/domain/institutions"
	locationDomain "obas/domain/location"
	userDomain "obas/domain/users"
	addressIO "obas/io/address"
	applicationIO "obas/io/applications"
	demographyIO "obas/io/demographics"
	documentIO "obas/io/documents"
	locationIO "obas/io/location"
	loginIO "obas/io/login"
	storageIO "obas/io/storage"
	usersIO "obas/io/users"
	utilIO "obas/io/util"
	"os"
	"strconv"
	"strings"
	"time"
)

type AddressPlaceHolder struct {
	AddressName string
	Address     string
	PostalCode  string
}

type ContactPlaceHolder struct {
	ContactName   string
	ContactDetail string
}

type DistrictData struct {
	Student   userDomain.User
	Provinces []locationDomain.Location
	TownName  string
	Alert     genericHelper.PageToast
	Menu      string
	SubMenu   string
}

func Users(app *config.Env) http.Handler {
	r := chi.NewRouter()
	r.Get("/", UsersHandler(app))
	r.Get("/admin", AdminHandler(app))
	r.Get("/student", StudentHandler(app))

	r.Get("/processingStatus", ProcessingStatusTypeHandler(app))
	r.Get("/student/bursary", StudentApplicationStatusHandler(app))
	r.Get("/studentContact", StudentContactsHandler(app))
	r.Get("/studentResults", StudentResultsHandler(app))

	r.Get("/student/profile/personal", StudentProfilePersonalHandler(app))
	r.Get("/student/profile/demography", StudentProfileDemographyHandler(app))
	r.Get("/student/profile/address", StudentProfileAddressHandler(app))
	r.Get("/student/profile/relative", StudentProfileRelativeHandler(app))
	r.Get("/student/profile/settings", StudentProfileSettingsHandler(app))
	r.Get("/student/profile/courses", StudentProfileCourseHandler(app))
	r.Get("/student/profile/academics", StudentProfileSubjectHandler(app))
	r.Get("/student/profile/districts", StudentProfileDistrictHandler(app))
	r.Get("/student/profile/contacts", StudentProfileContactsHandler(app))

	r.Post("/student/profile/personal/update", UpdateStudentProfilePersonalHandler(app))
	r.Post("/student/profile/address/addresstype", StudentProfileAddressTypeHandler(app))
	r.Post("/student/profile/address/update", StudentProfileAddressUpdateHandler(app))
	r.Post("/student-profile-relative-upate", StudentProfileRelativeUpdateHandler(app))
	r.Post("/student-profile-demography-update", StudentProfileDemographyUpdateHandler(app))
	r.Post("/student-profile-password-update", StudentProfilePasswordUpdate(app))
	r.Post("/student/profile/contact/contacttype", StudentProfileContactTypeHandler(app))
	r.Post("/student-profile-contact-update", StudentProfileContactUpdateHandler(app))
	r.Post("/student-profile-town-update", StudentProfileTownUpdateHandler(app))
	r.Post("/student-document-file-upload", StudentDocumentsUploadHandler(app))

	r.Get("/student/bursary/application", StudentBursaryApplicationHandler(app))
	r.Post("/student/bursary/application/start", StudentBursaryApplicationStartHandler(app))
	r.Post("/student/bursary/application/institution/matric/update", StudentBursaryApplicationMatricHandler(app))
	r.Post("/student/bursary/application/type/update", StudentBursaryApplicationTypeHandler(app))
	r.Post("/student/bursary/application/matric/subject/update", StudentBursaryApplicationMatricSubjectHandler(app))
	r.Post("/student/bursary/application/institution/current/update", StudentBursaryApplicationCurrentInstitutionHandler(app))
	r.Post("/student/bursary/application/institution/current/course/update", StudentBursaryApplicationCurrentInstitutionCourseHandler(app))
	r.Post("/student/bursary/application/institution/current/subject/update", StudentBursaryApplicationCurrentInstitutionSubjectHandler(app))
	r.Post("/student/bursary/application/institution/prospective/update", StudentBursaryApplicationProspectiveInstitutionHandler(app))
	r.Post("/student/bursary/application/institution/prospective/course/update", StudentBursaryApplicationProspectiveInstitutionCourseHandler(app))

	return r
}

func StudentBursaryApplicationProspectiveInstitutionCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Prospective Institution Course NOT saved!"
		successMessage := "Prospective Institution Course saved!"
		proceed := true

		r.ParseForm()
		applicationId := r.PostFormValue("applicationId")
		courseId := r.PostFormValue("course")

		if applicationId == "" || courseId == "" {
			proceed = false
			errorMsg := " Course can't be empty!"
			app.ErrorLog.Println(errorMsg)
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
		}
		if proceed {
			userApplicationCourse := userDomain.UserApplicationCourse{user.Email, applicationId, courseId}
			app.InfoLog.Println("User prospective course to save: ", userApplicationCourse)
			userApplicationCourse, err = usersIO.CreateUserApplicationCourse(userApplicationCourse)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			} else {
				app.InfoLog.Println("User Prospective Institution Course saved: ", userApplicationCourse)
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			}
		}
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationProspectiveInstitutionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Prospective Institution NOT saved!"
		successMessage := "Prospective Institution saved!"
		proceed := true

		r.ParseForm()
		applicationId := r.PostFormValue("applicationId")
		institutionId := r.PostFormValue("institution")

		if applicationId == "" || institutionId == "" {
			proceed = false
			errorMsg := " Institution can't be empty!"
			app.ErrorLog.Println(errorMsg)
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
		}
		if proceed {
			userApplicationInstitution := userDomain.UserApplicationInstitution{user.Email, applicationId, institutionId}
			app.InfoLog.Println("User prospective institution to save: ", userApplicationInstitution)
			userApplicationInstitution, err = usersIO.CreateUserApplicationInstitution(userApplicationInstitution)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			} else {
				app.InfoLog.Println("User Prospective Institution saved: ", userApplicationInstitution)
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			}
		}
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationCurrentInstitutionSubjectHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Current Institution Subject NOT saved!"
		successMessage := "Current Institution Subject saved!"
		proceed := true
		score := 0.00

		r.ParseForm()
		applicationId := r.PostFormValue("applicationId")
		subjectId := r.PostFormValue("subject")
		scoreStr := r.PostFormValue("score")

		if applicationId == "" || subjectId == "" {
			proceed = false
			errorMsg := " Subject can't be empty!"
			app.ErrorLog.Println(errorMsg)
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
		} else {
			score, err = strconv.ParseFloat(scoreStr, 64)
			if err != nil {
				proceed = false
				errorMsg := " ~ Possible incorrect subject score value."
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
			}
		}
		if proceed {
			userTertiarySubject := userDomain.UserTertiarySubject{user.Email, applicationId, subjectId, score}
			app.InfoLog.Println("User tertiary subject to create: ", userTertiarySubject)
			userTertiarySubject, err := usersIO.CreateUserTertiarySubject(userTertiarySubject)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			} else {
				app.InfoLog.Println("User Tertiary Subject saved: ", userTertiarySubject)
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			}
		}
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationCurrentInstitutionCourseHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Current Institution Course NOT saved!"
		successMessage := "Current Institution Course saved!"
		proceed := true

		r.ParseForm()
		applicationId := r.PostFormValue("applicationId")
		courseId := r.PostFormValue("course")

		if applicationId == "" || courseId == "" {
			proceed = false
			errorMsg := " Course can't be empty!"
			app.ErrorLog.Println(errorMsg)
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
		}
		if proceed {
			userTertiaryCourse := userDomain.UserTertiaryCourse{user.Email, applicationId, courseId}
			app.InfoLog.Println("User current course to save: ", userTertiaryCourse)
			userTertiaryCourse, err = usersIO.CreateUserTertiaryCourse(userTertiaryCourse)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			} else {
				app.InfoLog.Println("User Current Institution Course saved: ", userTertiaryCourse)
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			}
		}
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationCurrentInstitutionHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Current Institution NOT saved!"
		successMessage := "Current Institution saved!"
		debtAmount := 0.00
		proceed := true
		var userTertiaryInstitution userDomain.UserTertiaryInstitution

		r.ParseForm()
		applicationId := r.PostFormValue("applicationId")
		institutionId := r.PostFormValue("institution")
		debtAmountStr := r.PostFormValue("debt")
		if applicationId == "" || institutionId == "" {
			proceed = false
			errorMsg := " Institution can't be empty!"
			app.ErrorLog.Println(errorMsg)
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
		} else {
			if debtAmountStr != "" {
				debtAmount, err = strconv.ParseFloat(debtAmountStr, 64)
				if err != nil {
					proceed = false
					errorMsg := " ~ Possible incorrect debt amount value."
					app.ErrorLog.Println(err.Error())
					genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
				}
			}
		}
		if proceed {
			userTertiaryInstitution = userDomain.UserTertiaryInstitution{user.Email, applicationId, institutionId, debtAmount}
			app.InfoLog.Println("User Current Institution to save: ", userTertiaryInstitution)
			userTertiaryInstitution, err = usersIO.CreateUserTertiaryInstitution(userTertiaryInstitution)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			} else {
				app.InfoLog.Println("User Current Institution saved: ", userTertiaryInstitution)
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			}
		}
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationMatricSubjectHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Subject score NOT saved!"
		successMessage := "Subject scored saved!"
		var userMatricSubject userDomain.UserMatricSubject

		r.ParseForm()
		subjectId := r.PostFormValue("subject")
		scoreStr := r.PostFormValue("score")
		if subjectId == "" || scoreStr == "" {
			errorMsg := " Subject and/or score can't be empty!"
			app.ErrorLog.Println(errorMsg)
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
		} else {
			score, err := strconv.ParseFloat(scoreStr, 64)
			if err != nil {
				errorMsg := " ~ Possible incorrect subject score value."
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+errorMsg)
			} else {
				userMatricSubject = userDomain.UserMatricSubject{user.Email, subjectId, score}
				app.InfoLog.Println("UserMatricSubject to save: ", userMatricSubject)
				userMatricSubject, err = usersIO.CreateUserMatricSubject(userMatricSubject)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
				} else {
					genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
				}
			}
		}
		app.InfoLog.Println("Matic Subject saved: ", userMatricSubject)
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationTypeHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		_, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Applicant Type NOT updated!"
		successMessage := "Applicant Type updated!"

		r.ParseForm()
		applicantTypeId := r.PostFormValue("applicantType")
		applicationId := r.PostFormValue("applicationId")
		application, err := applicationIO.GetApplication(applicationId)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
		} else {
			updatedApplication := application
			updatedApplication.ApplicantTypeId = applicantTypeId
			saved, err := applicationIO.UpdateApplication(updatedApplication, token)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			} else {
				app.InfoLog.Println("update application response is ", saved)
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			}
		}
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationMatricHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		failureMessage := "Matric Institution NOT saved!"
		successMessage := "Matric Institution saved!"

		r.ParseForm()
		institution := r.PostFormValue("institution")
		userMatricInstitution := userDomain.UserMatricInstitution{user.Email, institution}
		app.InfoLog.Println("User (matric) institution to save: ", userMatricInstitution)

		saved, err := usersIO.CreateUserMatricInstitution(userMatricInstitution)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
		} else {
			genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
		}
		app.InfoLog.Println("Matric Institution saved: ", saved)
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentBursaryApplicationHandler(app *config.Env) http.HandlerFunc {
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

		var alert genericHelper.PageToast
		var applicationTypes []applicationDomain.ApplicationType
		var applicantTypes []applicationDomain.ApplicantType
		var application applicationDomain.Application
		var provinces []locationDomain.Location
		var institutionTypes []institutionDomain.InstitutionTypes
		var userMatricInstitution userDomain.UserMatricInstitution
		var userMatricInstitutionName, userTertiaryInstitutionName, userApplicationInstitutionName string
		var matricSubjects []academicsDomain.Subject
		var eUserMatricSubjects []ExtendedUserMatricSubject
		var userTertiaryInstitution userDomain.UserTertiaryInstitution
		var userTertiaryCourse userDomain.UserTertiaryCourse
		var currentTertiaryCourses, prospectiveTertiaryCourses []academicsDomain.Course
		var currentCourseSubjects []academicsDomain.Subject
		var eUserCurrentTertiarySubjects []ExtendedUserTertiarySubject
		var userApplicationInstitution userDomain.UserApplicationInstitution
		var userApplicationCourse userDomain.UserApplicationCourse
		var documentTypes []documentDomain.DocumentType
		var userDocuments []ExtendedDocument
		var applicationProgress float64
		var matricApplicant applicationDomain.ApplicantType
		var progressBarValue float64
		isComplete := true
		isMatricApplicant := false

		latestUserApplication, err := usersIO.GetLatestUserApplication(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve your latest application!"}
		} else {
			if latestUserApplication.ApplicationId != "" {
				isComplete, err = applicationIO.IsApplicationCompleted(latestUserApplication.ApplicationId)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve status of latest application"}
				} else {
					application, err = applicationIO.GetApplication(latestUserApplication.ApplicationId)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve application!"}
					} else {
						matricApplicant, alert = getMatricApplicantType(app)
						if alert.AlertInfo == "" {
							isMatricApplicant = matricApplicant.Id == application.ApplicantTypeId
							if isMatricApplicant {
								progressBarValue = genericHelper.MatricProgressBarIncrementVale
							} else {
								progressBarValue = genericHelper.NonMatricProgressBarIncrementValue
							}
							if application.ApplicantTypeId != "" {
								applicationProgress += progressBarValue
							}
						}
					}
				}
			}
			applicantTypes, alert = getApplicantTypes(app)
			proceed := alert.AlertInfo == ""
			if proceed {
				if isComplete {
					applicationTypes, alert = getApplicationTypes(app)
					proceed = alert.AlertInfo == ""
					if proceed {
						alert = genericHelper.CheckForSessionAlert(app, r)
					}
				} else {
					provinces, alert = locationHelper.GetProvinces(app)
					proceed = alert.AlertInfo == ""
					if proceed {
						institutionTypes, alert = getInstitutionTypes(app)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						documentTypes, alert = getDocumentTypes(app)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						userMatricInstitution, alert = getUserMatricInstitution(app, user.Email)
						proceed = alert.AlertInfo == ""
					}
					if proceed && userMatricInstitution.InstitutionId != "" {
						applicationProgress += progressBarValue
						userMatricInstitutionName, alert = institutionHelper.GetInstitutionName(app, userMatricInstitution.InstitutionId)
						proceed = alert.AlertInfo == ""
					}
					if proceed && userMatricInstitution.InstitutionId != "" {
						matricSubjects, alert = getMatricSubjects(app, userMatricInstitution.InstitutionId)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						eUserMatricSubjects, alert = getTransformedUserMatricSubjects(app, user.Email)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						applicationProgress += progressBarValue * float64(len(eUserMatricSubjects))
						if !isMatricApplicant {
							userTertiaryInstitution, alert = getUserTertiaryInstitutionForApplication(app, user.Email, application.Id)
							proceed = alert.AlertInfo == ""
						}
					}
					if proceed && userTertiaryInstitution.InstitutionId != "" && !isMatricApplicant {
						applicationProgress += progressBarValue
						userTertiaryInstitutionName, alert = institutionHelper.GetInstitutionName(app, userTertiaryInstitution.InstitutionId)
						proceed = alert.AlertInfo == ""
					}
					if proceed && userTertiaryInstitution.InstitutionId != "" && !isMatricApplicant {
						currentTertiaryCourses, alert = getInstitutionCourses(app, userTertiaryInstitution.InstitutionId)
						proceed = alert.AlertInfo == ""
					}
					if proceed && !isMatricApplicant {
						userTertiaryCourse, alert = getUserTertiaryCourseForApplication(app, user.Email, application.Id)
						proceed = alert.AlertInfo == ""
					}
					if proceed && userTertiaryCourse.CourseId != "" && !isMatricApplicant {
						applicationProgress += progressBarValue
						currentCourseSubjects, alert = getCourseSubjects(app, userTertiaryCourse.CourseId)
						proceed = alert.AlertInfo == ""
					}
					if proceed && !isMatricApplicant {
						eUserCurrentTertiarySubjects, alert = getTransformedUserTertiarySubjects(app, user.Email, application.Id)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						applicationProgress += progressBarValue * float64(len(eUserCurrentTertiarySubjects))
						userApplicationInstitution, alert = getUserApplicationInstitution(app, user.Email, application.Id)
						proceed = alert.AlertInfo == ""
					}
					if proceed && userApplicationInstitution.InstitutionId != "" {
						applicationProgress += progressBarValue
						userApplicationInstitutionName, alert = institutionHelper.GetInstitutionName(app, userApplicationInstitution.InstitutionId)
						proceed = alert.AlertInfo == ""
					}
					if proceed && userApplicationInstitution.InstitutionId != "" {
						prospectiveTertiaryCourses, alert = getInstitutionCourses(app, userApplicationInstitution.InstitutionId)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						userApplicationCourse, alert = getUserApplicationCourse(app, user.Email, application.Id)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						if userApplicationCourse.CourseId != "" {
							applicationProgress += progressBarValue
						}
						userDocuments, alert = getUserDocumentData(app, user.Email)
						proceed = alert.AlertInfo == ""
					}
					if proceed {
						applicationProgress += progressBarValue * float64(len(userDocuments))
						alert = genericHelper.CheckForSessionAlert(app, r)
					}
				}
			}
		}

		type PageData struct {
			Student                        userDomain.User
			Menu                           string
			SubMenu                        string
			ApplicationTypes               []applicationDomain.ApplicationType
			ApplicantTypes                 []applicationDomain.ApplicantType
			DocumentTypes                  []documentDomain.DocumentType
			Application                    applicationDomain.Application
			IsComplete                     bool
			Provinces                      []locationDomain.Location
			InstitutionTypes               []institutionDomain.InstitutionTypes
			UserMatricInstitutionName      string
			MatricSubjects                 []academicsDomain.Subject
			UserMatricSubjects             []ExtendedUserMatricSubject
			UserTertiaryInstitutionName    string
			UserTertiaryInstitution        userDomain.UserTertiaryInstitution
			CurrentTertiaryCourses         []academicsDomain.Course
			UserTertiaryCourse             userDomain.UserTertiaryCourse
			CurrentCourseSubjects          []academicsDomain.Subject
			UserTertiarySubjects           []ExtendedUserTertiarySubject
			UserApplicationInstitutionName string
			ProspectiveTertiaryCourses     []academicsDomain.Course
			UserApplicationCourse          userDomain.UserApplicationCourse
			UserDocuments                  []ExtendedDocument
			ApplicationProgress            float64
			IsMatricApplicant bool
			Alert                          genericHelper.PageToast
		}

		data := PageData{
			user,
			"bursary",
			"application",
			applicationTypes,
			applicantTypes,
			documentTypes,
			application,
			isComplete,
			provinces,
			institutionTypes,
			userMatricInstitutionName,
			matricSubjects,
			eUserMatricSubjects,
			userTertiaryInstitutionName,
			userTertiaryInstitution,
			currentTertiaryCourses,
			userTertiaryCourse,
			currentCourseSubjects,
			eUserCurrentTertiarySubjects,
			userApplicationInstitutionName,
			prospectiveTertiaryCourses,
			userApplicationCourse,
			userDocuments,
			applicationProgress,
			isMatricApplicant,
			alert}

		files := []string{
			app.Path + "content/student/bursary/application.html",
			app.Path + "content/student/template/sidebar.template.html",
			app.Path + "content/student/template/application/matric-form.template.html",
			app.Path + "content/student/template/application/current-institution-form.template.html",
			app.Path + "base/template/form/institution-form.template.html",
			app.Path + "base/template/form/location-form.template.html",
			app.Path + "content/student/template/application/prospective-institution-form.template.html",
			app.Path + "content/student/template/application/document.template.html",
			app.Path + "content/student/template/application/applicant-type-form.template.html",
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

func StudentBursaryApplicationStartHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		user, err := usersIO.GetUser(email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Redirect(w, r, "/login", 301)
			return
		}

		var newApplication = applicationDomain.Application{}
		failureMessage := "Application NOT created!"
		successMessage := "Application created!"
		isSuccessful := false

		r.ParseForm()
		applicationTypeId := r.PostFormValue("applicationType")
		applicantTypeId := r.PostFormValue("applicantType")

		if applicationTypeId != "" && applicantTypeId != "" {
			application := applicationDomain.Application{"", applicationTypeId, applicantTypeId}
			app.InfoLog.Println("Application to create: ", application)
			newApplication, err = applicationIO.CreateApplication(application)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			} else {
				if newApplication.Id != "" {
					userApplication := userDomain.UserApplication{email, newApplication.Id, time.Now()}
					app.InfoLog.Println("User Application to create: ", userApplication)
					_, err := usersIO.CreateUserApplication(userApplication)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
					} else {
						incompleteStatus, err := utilIO.GetIncompleteStatus()
						if err != nil {
							app.ErrorLog.Println(err.Error())
							genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
						} else {
							if incompleteStatus.Id != "" {
								userApplicationStatus := applicationIO.ApplicationStatus{newApplication.Id, incompleteStatus.Id, user.Email, "Starting Application", time.Now()}
								app.InfoLog.Println("Application Status to create: ", userApplicationStatus)
								_, err = applicationIO.CreateApplicationStatus(userApplicationStatus)
								if err != nil {
									app.ErrorLog.Println(err.Error() + " ~ User Application Status NOT created!")
									genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
								} else {
									isSuccessful = true
									genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
								}
							} else {
								app.ErrorLog.Println("No status id found!")
								genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
							}
						}
					}
				} else {
					app.ErrorLog.Println("No application id!")
					genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
				}
			}
		} else {
			error := "Application type and/or applicant type is null!"
			app.ErrorLog.Println(error)
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+" Reason: "+error)
		}
		app.InfoLog.Println("application response is ", isSuccessful)
		http.Redirect(w, r, "/users/student/bursary/application", 301)
	}
}

func StudentDocumentsUploadHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseMultipartForm(10 << 20)
		documentTypeId := r.PostFormValue("documenttype")
		file, handler, err := r.FormFile("file")
		if err != nil {
			app.ErrorLog.Println(err.Error())
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, "Internal Server Error!")
		} else {
			defer file.Close()
			tempFilePath := "temp_file/" + handler.Filename
			tempFile, err := os.OpenFile(tempFilePath, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, "Internal Server Error!")
			} else {
				defer tempFile.Close()
				io.Copy(tempFile, file)
				fileName := handler.Filename
				successMessage := "File: " + fileName + " updated!"
				failureMessage := "File: " + fileName + " NOT updated!"
				app.InfoLog.Println("File Name: ", fileName)
				app.InfoLog.Println("File Size: ", handler.Size)
				app.InfoLog.Println("File Header: ", handler.Header)

				fileData, err := storageIO.UploadFile(tempFilePath, token)
				os.Remove(tempFilePath)
				if err != nil {
					app.ErrorLog.Println(err.Error())
					genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
				} else {
					if strings.Contains(fileData.Id, "error") {
						app.ErrorLog.Println(fileData.Id)
						genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, "Internal Server Error!")
					} else {
						document := documentDomain.Document{
							fileData.Id,
							documentTypeId,
							fileName,
							fileData.Url,
							handler.Header.Get("Content-Type"),
							time.Now(),
							"", ""}
						app.InfoLog.Println("Document to create: ", document)
						documentSaved, err := documentIO.CreateDocument(document, token)
						if err != nil {
							app.ErrorLog.Println(err.Error())
							genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, "Document NOT saved!")
						} else {
							if documentSaved {
								userDocument := userDomain.UserDocument{email, fileData.Id}
								app.InfoLog.Println("User document to save: ", userDocument)
								userDocumentSaved, err := usersIO.UpdateUserDocument(userDocument, token)
								if err != nil {
									app.ErrorLog.Println(err.Error())
									genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, "User document saved!")
								} else {
									if userDocumentSaved {
										genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
									} else {
										genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
									}
								}
							} else {
								genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, "User document NOT saved!")
							}
						}
					}
				}
			}
		}
		http.Redirect(w, r, "/users/student/documents", 301)
	}
}

func StudentProfileTownUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		townCode := r.PostFormValue("town")
		userTown := userDomain.UserTown{email, townCode}
		app.InfoLog.Println("UserTown to update: ", userTown)
		updated, err := usersIO.UpdateUserTown(userTown, token)
		successMessage := "User town updated!"
		failureMessage := "User town NOT updated!"
		if err != nil {
			app.ErrorLog.Println(err.Error())
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
		} else {
			if updated {
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			} else {
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/districts", 301)
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

		var alert genericHelper.PageToast
		var provinces []locationDomain.Location
		var townName string

		provinces, alert = locationHelper.GetProvinces(app)
		if alert.AlertInfo != "" {
			userTown, err := usersIO.GetUserTown(email)
			if err != nil {
				app.ErrorLog.Println(err.Error())
				alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user town!"}
			} else {
				if userTown.LocationId != "" {
					location, err := locationIO.GetLocation(userTown.LocationId)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve location!"}
					} else {
						townName = location.Name
						alert = genericHelper.CheckForSessionAlert(app, r)
					}
				} else {
					townName = "<<not set>>"
					alert = genericHelper.CheckForSessionAlert(app, r)
				}
			}
		}

		data := DistrictData{
			user,
			provinces,
			townName,
			alert, "profile", "districts"}

		files := []string{
			app.Path + "content/student/profile/district_and_municipality.html",
			app.Path + "content/student/template/sidebar.template.html",
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

func StudentProfileContactUpdateHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		contactTypeId := r.PostFormValue("contactTypeId")
		contact := r.PostFormValue("contact")
		userContact := usersIO.UserContact{email, contactTypeId, contact}
		app.InfoLog.Println("UserContact to update: ", userContact)
		updated, err := usersIO.UpdateUserContact(userContact, token)
		successMessage := "User contact updated!"
		failureMessage := "User contact NOT updated!"
		if err != nil {
			app.ErrorLog.Println(err.Error())
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
		} else {
			if updated {
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			} else {
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("Update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/contacts", 301)
	}
}

func StudentProfileContactTypeHandler(app *config.Env) http.HandlerFunc {
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
		var alert genericHelper.PageToast
		var userContact usersIO.UserContact
		var contactName string
		var contactTypes []addressIO.ContactType
		var contacts []ContactPlaceHolder
		r.ParseForm()
		contactTypeId := r.PostFormValue("contacttypes")
		if contactTypeId == "" {
			errMsg := "No contact type selected!"
			app.ErrorLog.Println(errMsg)
			alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, errMsg}
		} else {
			contactTypes, err := addressIO.GetContactTypes()
			if err != nil {
				errMsg := "Could not retrieve contact types!"
				app.ErrorLog.Println(err.Error() + " - " + errMsg)
				alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, errMsg}
			} else {
				for _, contactType := range contactTypes {
					if contactTypeId == contactType.ContactTypeId {
						contactName = contactType.Name
					}
					userContact, err := usersIO.GetUserContact(email, contactType.ContactTypeId)
					if err != nil {
						errMsg := "Could not retrieve user contact for " + contactType.Name
						app.ErrorLog.Println(err.Error() + " - " + errMsg)
					} else {
						contacts = append(contacts, ContactPlaceHolder{contactType.Name, userContact.Contact})
					}
				}
				userContact, err = usersIO.GetUserContact(email, contactTypeId)
				if err != nil {
					errMsg := "Could not retrieve user contact for " + contactName
					app.ErrorLog.Println(err.Error() + " - " + errMsg)
					alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, errMsg}
				}
			}
		}

		type PageData struct {
			Student       userDomain.User
			ContactTypes  []addressIO.ContactType
			Contacts      []ContactPlaceHolder
			Contact       usersIO.UserContact
			ContactTypeId string
			ContactName   string
			Alert         genericHelper.PageToast
		}
		data := PageData{user, contactTypes, contacts, userContact, contactTypeId, contactName, alert}
		files := []string{
			app.Path + "content/student/profile/contacts.html",
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

func StudentProfileContactsHandler(app *config.Env) http.HandlerFunc {
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
		var alert genericHelper.PageToast
		var contacts []ContactPlaceHolder
		contactTypes, err := addressIO.GetContactTypes()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve contact types!"}
		} else {
			for _, contactType := range contactTypes {
				userContact, err := usersIO.GetUserContact(email, contactType.ContactTypeId)
				if err != nil {
					errMsg := "Could not retrieve user contact for " + contactType.Name
					app.ErrorLog.Println(err.Error() + " - " + errMsg)
				} else {
					contacts = append(contacts, ContactPlaceHolder{contactType.Name, userContact.Contact})
				}
			}
			alert = genericHelper.CheckForSessionAlert(app, r)
		}

		type PageData struct {
			Student       userDomain.User
			ContactTypes  []addressIO.ContactType
			Contacts      []ContactPlaceHolder
			Contact       usersIO.UserContact
			ContactTypeId string
			ContactName   string
			Alert         genericHelper.PageToast
			Menu          string
			SubMenu       string
		}

		data := PageData{user, contactTypes, contacts, usersIO.UserContact{}, "", "", alert, "profile", "contacts"}
		files := []string{
			app.Path + "content/student/profile/contacts.html",
			app.Path + "content/student/template/sidebar.template.html",
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

func StudentProfilePasswordUpdate(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := app.Session.GetString(r.Context(), "userId")
		token := app.Session.GetString(r.Context(), "token")
		if email == "" || token == "" {
			http.Redirect(w, r, "/login", 301)
			return
		}
		r.ParseForm()
		currentPassword := r.PostFormValue("current_password")
		newPasswordOne := r.PostFormValue("new_password_one")
		newPasswordTwo := r.PostFormValue("new_password_two")

		successMessage := "User password updated!"
		failureMessage := "User password NOT Updated!"

		if newPasswordOne != newPasswordTwo {
			errMsg := "New password mismatch."
			app.ErrorLog.Println(errMsg)
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage+" - "+errMsg)
		} else {
			userChangePassword := loginIO.ChangePassword{email, currentPassword, newPasswordOne, time.Now()}
			app.InfoLog.Println("User password to update: ", userChangePassword)
			loginToken, err := loginIO.DoChangePassword(userChangePassword, token)

			if err != nil {
				app.ErrorLog.Println(err.Error())
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			}
			app.Session.Put(r.Context(), "userId", loginToken.Email)
			app.Session.Put(r.Context(), "token", loginToken.Token)
			genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
		}
		http.Redirect(w, r, "/users/student/profile/settings", 301)
	}
}

func StudentProfileSettingsHandler(app *config.Env) http.HandlerFunc {
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

		alert := genericHelper.CheckForSessionAlert(app, r)

		type PageData struct {
			Student userDomain.User
			Alert   genericHelper.PageToast
			Menu    string
			SubMenu string
		}

		data := PageData{user, alert, "profile", "settings"}
		files := []string{
			app.Path + "content/student/profile/settings.html",
			app.Path + "content/student/template/sidebar.template.html",
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

		updated, err := usersIO.UpdateUserDemographics(userDemograpgy, token)
		successMessage := "User demography updated!"
		failureMessage := "User demography NOT Updated!"

		if err != nil {
			app.ErrorLog.Println(err.Error())
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
		} else {
			if updated {
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			} else {
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
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
		var alert genericHelper.PageToast
		var genders []demographyIO.Gender
		var races []demographyIO.Race
		var gender demographyIO.Gender
		var race demographyIO.Race
		var title demographyIO.Title
		titles, err := demographyIO.GetTitles()
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve titles!"}
		} else {
			genders, err = demographyIO.GetGenders()
			if err != nil {
				app.ErrorLog.Println(err.Error())
				alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve genders!"}
			} else {
				races, err = demographyIO.GetRaces()
				if err != nil {
					app.ErrorLog.Println(err.Error())
					alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve races!"}
				} else {
					userDemography, err := usersIO.GetUserDemographic(email)
					if err != nil {
						app.ErrorLog.Println(err.Error())
						alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve student demography!"}
					} else {
						title = getUserTitle(userDemography, titles)
						gender = getUserGender(userDemography, genders)
						race = getUserRace(userDemography, races)
						alert = genericHelper.CheckForSessionAlert(app, r)
					}
				}
			}
		}

		type PageData struct {
			Student       userDomain.User
			Titles        []demographyIO.Title
			Genders       []demographyIO.Gender
			Races         []demographyIO.Race
			Alert         genericHelper.PageToast
			StudentTitle  demographyIO.Title
			StudentGender demographyIO.Gender
			StudentRace   demographyIO.Race
			Menu          string
			SubMenu       string
		}

		data := PageData{
			user,
			titles,
			genders,
			races,
			alert,
			title,
			gender,
			race,
			"profile",
			"demography"}
		app.InfoLog.Println("DistrictData: ", data.Alert)
		files := []string{
			app.Path + "content/student/profile/demography.html",
			app.Path + "content/student/template/sidebar.template.html",
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
			genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
		} else {
			if updated {
				genericHelper.SetSessionMessage(app, r, genericHelper.SuccessAlertStyle, successMessage)
			} else {
				genericHelper.SetSessionMessage(app, r, genericHelper.DangerAlertStyle, failureMessage)
			}
		}
		app.InfoLog.Println("UserRelative update response is ", updated)
		http.Redirect(w, r, "/users/student/profile/relative", 301)

	}
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
			Student userDomain.User
		}

		data := PageData{user}
		files := []string{
			app.Path + "content/student/profile/academics.html",
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
			Student userDomain.User
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

		var alert genericHelper.PageToast

		userRelative, err := usersIO.GetUserRelative(user.Email)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve student relative!"}
		} else {
			alert = genericHelper.CheckForSessionAlert(app, r)
		}

		type PageData struct {
			Student         userDomain.User
			StudentRelative usersIO.UserRelative
			Alert           genericHelper.PageToast
			Menu            string
			SubMenu         string
		}

		data := PageData{user, userRelative, alert, "profile", "relative"}
		files := []string{
			app.Path + "content/student/profile/relative.html",
			app.Path + "content/student/template/sidebar.template.html",
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
			Student       userDomain.User
			AddressTypes  []addressIO.AddressType
			Addresses     []AddressPlaceHolder
			Address       usersIO.UserAddress
			AddressTypeID string
			AddressName   string
			Menu          string
			SubMenu       string
		}

		data := PageData{user, addressTypes, addresses, usersIO.UserAddress{}, "", "", "profile", "address"}
		files := []string{
			app.Path + "content/student/profile/address.html",
			app.Path + "content/student/template/sidebar.template.html",
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
			Student       userDomain.User
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
			Student userDomain.User
			Menu    string
			SubMenu string
		}
		data := PageData{user, "", ""}
		files := []string{
			app.Path + "content/student/student_dashboard.page.html",
			app.Path + "content/student/template/sidebar.template.html",
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
		dobString := genericHelper.GetDate_YYYYMMDD(user.DateOfBirth.String()) // split date and get in format: yyy-mm-dd

		type PageData struct {
			Student     userDomain.User
			DateOfBirth string
			Menu        string
			SubMenu     string
		}

		data := PageData{user, dobString, "profile", "personal"}
		files := []string{
			app.Path + "content/student/profile/personal.html",
			app.Path + "content/student/template/sidebar.template.html",
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
		dateOfBirth, _ := time.Parse(genericHelper.LayoutOBAS, dateOfBirthStr)
		user := userDomain.User{email, idNumber, firstName, "", lastName, dateOfBirth}
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
			app.Path + "/base/footer.template.html",*/
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
			app.Path + "/base/footer.template.html",
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
			//academics []io.ProcessingStatusType
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.template.html",
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
			app.Path + "content/student/bursary/bursary.html",
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
			//academics []io.StudentContacts
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.template.html",
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

func StudentResultsHandler(app *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//allStudentResults, err := io.GetStudentResults()
		//
		//if err != nil {
		//	app.ServerError(w, err)
		//}

		type PageData struct {
			//academics []io.StudentResults
			name string
		}
		data := PageData{""}

		files := []string{
			app.Path + "/users/users.page.html",
			app.Path + "/base/base.page.html",
			app.Path + "/base/navbar.page.html",
			app.Path + "/base/sidebarOld.page.html",
			app.Path + "/base/footer.template.html",
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
