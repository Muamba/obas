package controllers

import (
	"obas/config"
	genericHelper "obas/controllers/misc"
	domain "obas/domain/application"
	applicationIO "obas/io/applications"
)

func GetApplicantTypes(app *config.Env)([]domain.ApplicantType, genericHelper.PageToast){
	var applicantTypes []domain.ApplicantType
	var alert genericHelper.PageToast
	applicantTypes, err := applicationIO.GetApplicantTypes()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve applicant types!"}
	}
	return applicantTypes, alert
}

func GetApplicationTypes(app *config.Env)([]domain.ApplicationType, genericHelper.PageToast) {
	var applicationTypes []domain.ApplicationType
	var alert genericHelper.PageToast
	applicationTypes, err := applicationIO.GetApplicationTypes()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve application types!"}
	}
	return applicationTypes, alert
}

func GetMatricApplicantType(app *config.Env) (domain.ApplicantType, genericHelper.PageToast) {
	var matricApplicantType domain.ApplicantType
	var alert genericHelper.PageToast
	matricApplicantType, err := applicationIO.GetMatricApplicantType()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve matric applicant type!"}
	}
	return matricApplicantType, alert

}
