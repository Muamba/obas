package applications

import (
	"errors"
	"obas/api"
	domain "obas/domain/application"
)

const applicationStatusUrl = api.BASE_URL + "/application/status/"

type ApplicationStatus domain.ApplicationStatus

func GetAllStatusesForApplication(applicationId string) ([]ApplicationStatus, error) {
	entites := []ApplicationStatus{}
	resp, _ := api.Rest().Get(applicationStatusUrl + "all/" + applicationId)
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetApplicationStatus(applicationId string) (ApplicationStatus, error) {
	entity := ApplicationStatus{}
	resp, _ := api.Rest().Get(applicationStatusUrl + "getforapplication/" + applicationId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateApplicationStatus(entity ApplicationStatus) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationStatusUrl + "create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func GetLatestForStatus(applicationId, statusId string) (ApplicationStatus, error) {
	entity := ApplicationStatus{}
	resp, _ := api.Rest().Get(applicationStatusUrl + "getforstatus/" + applicationId + "/" + statusId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func IsApplicationCompleted(applicationId string) (bool, error) {
	isComplete := false
	resp, _ := api.Rest().
		Get(applicationStatusUrl + "iscompleted/"+applicationId)
	if resp.IsError() {
		return isComplete, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &isComplete)
	if err != nil {
		return isComplete, errors.New(resp.Status())
	}
	return isComplete, nil
}
