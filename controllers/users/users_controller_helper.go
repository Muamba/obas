package users

import (
	"obas/config"
	academicsHelper "obas/controllers/academics"
	institutionHelper "obas/controllers/institutions"
	genericHelper "obas/controllers/misc"
	academicsDomain "obas/domain/academics"
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
