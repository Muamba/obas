package domain

import "time"

type ApplicantType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ApplicationResult struct {
	ApplicationResultId string `json:"applicationResultId"`
	Description         string `json:"description"`
	Date                string `json:"date"`
}

type ApplicationStatus struct {
	ApplicationId string    `json:"applicationId"`
	StatusId      string    `json:"statusId"`
	ModifiedBy    string    `json:"modifiedBy"`
	Comment       string    `json:"comment"`
	DateTime      time.Time `json:"dateTime"`
}

type Application struct {
	Id                string `json:"id"`
	ApplicationTypeId string `json:"applicationTypeId"`
	ApplicantTypeId   string `json:"applicantTypeId"`
}

type ApplicationType struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
