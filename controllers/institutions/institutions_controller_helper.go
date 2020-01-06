package institutions

import (
	"obas/config"
	genericHelper "obas/controllers/misc"
	institutionDomain "obas/domain/institutions"
	institutionIO "obas/io/institutions"
)

/**
Get institution types
*/
func GetInstitutionTypes(app *config.Env) ([]institutionDomain.InstitutionTypes, genericHelper.PageToast) {
	var alert genericHelper.PageToast
	var institutionTypes []institutionDomain.InstitutionTypes
	institutionTypes, err := institutionIO.GetInstitutionTypes()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve institution types!"}
	}
	return institutionTypes, alert
}

/**
Get courses for institution given institution id
*/
func GetInstitutionCourses(app *config.Env, institutionId string) ([]institutionDomain.InstitutionCourse, genericHelper.PageToast) {
	var institutionCourses []institutionDomain.InstitutionCourse
	var alert genericHelper.PageToast
	institutionCourses, err := institutionIO.GetInstitutionCourses(institutionId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve institution courses!"}
	}
	return institutionCourses, alert
}

/**
Get institution given institution id
*/
func GetInstitution(app *config.Env, institutionId string) (institutionDomain.Institution, genericHelper.PageToast) {
	var institution institutionDomain.Institution
	var alert genericHelper.PageToast
	institution, err := institutionIO.GetInstitution(institutionId)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve institution!"}
	}
	return institution, alert
}

/**
Get user institution name given institution id
*/
func GetInstitutionName(app *config.Env, institutionId string) (string, genericHelper.PageToast) {
	var institutionName string
	var institution institutionDomain.Institution
	var alert genericHelper.PageToast
	institution, alert = GetInstitution(app, institutionId)
	if alert.AlertInfo == "" {
		institutionName = institution.Name
	}
	return institutionName, alert
}
