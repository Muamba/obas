package users

import (
	"obas/config"
	academicsHelper "obas/controllers/academics"
	applicationHelper "obas/controllers/application"
	documentHelper "obas/controllers/documents"
	institutionHelper "obas/controllers/institutions"
	genericHelper "obas/controllers/misc"
	academicsDomain "obas/domain/academics"
	applicationDomain "obas/domain/application"
	documentDomain "obas/domain/documents"
	institutionDomain "obas/domain/institutions"
	userDomain "obas/domain/users"
	demographyIO "obas/io/demographics"
	usersIO "obas/io/users"
)

type ExtendedUserMatricSubject struct {
	userDomain.UserMatricSubject
	SubjectName string
}

type ExtendedUserTertiarySubject struct {
	userDomain.UserTertiarySubject
	SubjectName string
}

type ExtendedDocument struct {
	documentDomain.Document
	DocumentType        string
	DocumentDate        string
	DocumentStatusBadge string
}

/**
Get user matric institution given user id
*/
func getUserMatricInstitution(app *config.Env, userId string) (userDomain.UserMatricInstitution, genericHelper.PageToast) {
	var userMatricInstitution userDomain.UserMatricInstitution
	var alert genericHelper.PageToast
	userMatricInstitution, err := usersIO.GetUserMatricInstitution(userId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user matric institution!"}
	}
	return userMatricInstitution, alert
}

/**
Filter user race from list of race given user demography
*/
func getUserRace(demography usersIO.UserDemography, races []demographyIO.Race) demographyIO.Race {
	for _, race := range races {
		if demography.RaceId == race.RaceId {
			return race
		}
	}
	return demographyIO.Race{}
}

/**
Filter user title form list of titles given user demography
*/
func getUserTitle(demography usersIO.UserDemography, titles []demographyIO.Title) demographyIO.Title {
	for _, title := range titles {
		if demography.TitleId == title.TitleId {
			return title
		}
	}
	return demographyIO.Title{}
}

/**
Filter user gender from list of genders given user demography
*/
func getUserGender(demography usersIO.UserDemography, genders []demographyIO.Gender) demographyIO.Gender {
	for _, gender := range genders {
		if demography.GenderId == gender.GenderId {
			return gender
		}
	}
	return demographyIO.Gender{}
}

/**
Get Matric subjects given institution id
*/
func getMatricSubjects(app *config.Env, institutionId string) ([]academicsDomain.Subject, genericHelper.PageToast) {
	var subjects []academicsDomain.Subject
	var courseSubjects []academicsDomain.CourseSubject
	var alert genericHelper.PageToast
	institutionCourses, alert := institutionHelper.GetInstitutionCourses(app, institutionId)
	if alert.AlertInfo == "" && len(institutionCourses) > 0 {
		institutionCourse := institutionCourses[0]
		courseSubjects, alert = academicsHelper.GetSubjectsForCourse(app, institutionCourse.CourseId)
	}
	if alert.AlertInfo == "" {
		for _, courseSubject := range courseSubjects {
			subject, alert := academicsHelper.GetSubject(app, courseSubject.SubjectId)
			if alert.AlertInfo == "" {
				subjects = append(subjects, subject)
			}
		}
	}
	return subjects, alert
}

func getInstitutionCourses(app *config.Env, institutionId string) ([]academicsDomain.Course, genericHelper.PageToast) {
	var courses []academicsDomain.Course
	institutionCourses, alert := institutionHelper.GetInstitutionCourses(app, institutionId)
	if alert.AlertInfo == "" && len(institutionCourses) > 0 {
		for _, institutionCourse := range institutionCourses {
			course, alert := academicsHelper.GetCourse(app, institutionCourse.CourseId)
			if alert.AlertInfo == "" {
				courses = append(courses, course)
			}
		}
	}
	return courses, alert
}

/**
Get (ETL) transformed user matric subjects given userId
*/
func getTransformedUserMatricSubjects(app *config.Env, userId string) ([]ExtendedUserMatricSubject, genericHelper.PageToast) {
	var eUserMatricSubjects []ExtendedUserMatricSubject
	userMatricSubjects, alert := getUserMatricSubjects(app, userId)
	if alert.AlertInfo == "" {
		for _, userMatricSubject := range userMatricSubjects {
			subject, alert := academicsHelper.GetSubject(app, userMatricSubject.SubjectId)
			if alert.AlertInfo == "" {
				eUserMatricSubject := ExtendedUserMatricSubject{userMatricSubject, subject.Name}
				eUserMatricSubjects = append(eUserMatricSubjects, eUserMatricSubject)
			} else {
				break
			}
		}
	}
	return eUserMatricSubjects, alert
}

/**
Get (ETL) user tertiary subjects given userid and application id
*/
func getTransformedUserTertiarySubjects(app *config.Env, userId, applicationId string) ([]ExtendedUserTertiarySubject, genericHelper.PageToast) {
	var eUserTertiarySubjects []ExtendedUserTertiarySubject
	userTertiarySubjects, alert := getUserTertiarySubjectsForApplication(app, userId, applicationId)
	if alert.AlertInfo == "" {
		for _, userTertiarySubject := range userTertiarySubjects {
			subject, alert := academicsHelper.GetSubject(app, userTertiarySubject.SubjectId)
			if alert.AlertInfo == "" {
				eUserTertiarySubject := ExtendedUserTertiarySubject{userTertiarySubject, subject.Name}
				eUserTertiarySubjects = append(eUserTertiarySubjects, eUserTertiarySubject)
			} else {
				app.ErrorLog.Println(alert.AlertInfo)
			}
		}
	}
	return eUserTertiarySubjects, alert
}

/**
Get user tertiary subjects for application given user id and application id
 */
func getUserTertiarySubjectsForApplication(app *config.Env, userId, applicationId string) ([]userDomain.UserTertiarySubject, genericHelper.PageToast) {
	var userTertiarySubjects []userDomain.UserTertiarySubject
	var alert genericHelper.PageToast
	userTertiarySubjects, err := usersIO.GetUserTertiarySubjects(userId, applicationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user tertiary subjects!"}
	}
	return userTertiarySubjects, alert
}

/**
Get user matric subjects given user id.
*/
func getUserMatricSubjects(app *config.Env, userId string) ([]userDomain.UserMatricSubject, genericHelper.PageToast) {
	var userMatricSubjects []userDomain.UserMatricSubject
	var alert genericHelper.PageToast
	userMatricSubjects, err := usersIO.GetUserMatricSubjects(userId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user matric subjects!"}
	}
	return userMatricSubjects, alert
}

/**
Get user tertiary institution for application given user id
*/
func getUserTertiaryInstitutionForApplication(app *config.Env, userId string, applicationId string) (userDomain.UserTertiaryInstitution, genericHelper.PageToast) {
	var userTertiaryInstitution userDomain.UserTertiaryInstitution
	var alert genericHelper.PageToast
	userTertiaryInstitution, err := usersIO.GetUserTertiaryInstitutionForApplication(userId, applicationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user current institution!"}
	}
	return userTertiaryInstitution, alert
}

/**
Get user tertiary course for application given user id
 */
func getUserTertiaryCourseForApplication(app *config.Env, userId string, applicationId string) (userDomain.UserTertiaryCourse, genericHelper.PageToast) {
	var userTertiaryCourse userDomain.UserTertiaryCourse
	var alert genericHelper.PageToast
	userTertiaryCourse, err := usersIO.GetUserTertiaryCourseForApplication(userId, applicationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user current course!"}
	}
	return userTertiaryCourse, alert
}

/**
Get user (prospective) application course given user id and application id
 */
func getUserApplicationCourse(app *config.Env, userId, applicationId string) (userDomain.UserApplicationCourse, genericHelper.PageToast) {
	var userApplicationCourse userDomain.UserApplicationCourse
	var alert genericHelper.PageToast
	userApplicationCourse, err := usersIO.GetUserApplicationCourse(userId, applicationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user application course!"}
	}
	return userApplicationCourse, alert
}

/**
Get subjects for course given course id
 */
func getCourseSubjects(app *config.Env, courseId string) ([]academicsDomain.Subject, genericHelper.PageToast) {
	var subjects []academicsDomain.Subject
	var courseSubjects []academicsDomain.CourseSubject
	var alert genericHelper.PageToast
	courseSubjects, alert = academicsHelper.GetSubjectsForCourse(app, courseId)
	if alert.AlertInfo == "" {
		for _, courseSubject := range courseSubjects {
			subject, alert := academicsHelper.GetSubject(app, courseSubject.SubjectId)
			if alert.AlertInfo == "" {
				subjects = append(subjects, subject)
			}
		}
	}
	return subjects, alert
}

/**
Get user application (prospective) institution given user id and application id
 */
func getUserApplicationInstitution(app *config.Env, userId, applicationId string) (userDomain.UserApplicationInstitution, genericHelper.PageToast) {
	var userApplicationInstitution userDomain.UserApplicationInstitution
	var alert genericHelper.PageToast
	userApplicationInstitution, err := usersIO.GetUserApplicationInstitution(userId, applicationId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user application institution!"}
	}
	return userApplicationInstitution, alert
}

/**
Get document types
 */
func getDocumentTypes(app *config.Env) ([]documentDomain.DocumentType, genericHelper.PageToast) {
	return documentHelper.GetDocumentTypes(app)
}

/**
Get institution types
 */
func getInstitutionTypes(app *config.Env) ([]institutionDomain.InstitutionTypes, genericHelper.PageToast) {
	return institutionHelper.GetInstitutionTypes(app)
}

/**
Get transformed/extended (ETL) user documents data given user id
 */
func getUserDocumentData(app *config.Env, userId string) ([]ExtendedDocument, genericHelper.PageToast) {
	var eUserDocuments []ExtendedDocument
	userDocuments, alert := getUserDocuments(app, userId)
	proceed := alert.AlertInfo == ""
	if proceed {
		for _, userDocument := range userDocuments {
			document, alert := documentHelper.GetDocument(app, userDocument.DocumentId)
			proceed = alert.AlertInfo == ""
			var documentType documentDomain.DocumentType
			if proceed {
				documentType, alert = documentHelper.GetDocumentType(app, document.DocumentTypeId)
				proceed = alert.AlertInfo == ""
			}
			if proceed {
				date := genericHelper.GetDate_YYYYMMDD(document.Date.String())
				var progressBadge string
				documentStatus := document.DocumentStatus
				if documentStatus == "Approved" {
					progressBadge = "badge-success"
				} else if documentStatus == "Not Approved" {
					progressBadge = "badge-danger"
				} else {
					progressBadge = "badge-warning"
				}
				eUserDocument := ExtendedDocument{document, documentType.DocumentTypename, date, progressBadge}
				eUserDocuments = append(eUserDocuments, eUserDocument)
			} else {
				app.ErrorLog.Println(alert.AlertInfo)
			}
		}
	}
	return eUserDocuments, alert
}

/**
Get user documents given user id
 */
func getUserDocuments(app *config.Env, userId string) ([]userDomain.UserDocument, genericHelper.PageToast) {
	var userDocuments []userDomain.UserDocument
	var alert genericHelper.PageToast
	userDocuments, err := usersIO.GetUserDocuments(userId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve user documents!"}
	}
	return userDocuments, alert
}

/**
Get applicant types
 */
func getApplicantTypes(app *config.Env) ([]applicationDomain.ApplicantType, genericHelper.PageToast) {
	return applicationHelper.GetApplicantTypes(app)
}

/**
Get application types
 */
func getApplicationTypes(app *config.Env) ([]applicationDomain.ApplicationType, genericHelper.PageToast) {
	return applicationHelper.GetApplicationTypes(app)
}

/**
Get matric appliant
 */
func getMatricApplicantType(app *config.Env) (applicationDomain.ApplicantType, genericHelper.PageToast) {
	return applicationHelper.GetMatricApplicantType(app)
}
