package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userapplicationURL = api.BASE_URL + "/users/application"

func CreateUserApplication(obj domain.UserApplication) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserApplication(userId, applicationId string) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	resp, _ := api.Rest().Get(userapplicationURL + "/get/" + userId + "/" + applicationId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserApplications(userId string) ([]domain.UserApplication, error) {
	entity := []domain.UserApplication{}
	resp, _ := api.Rest().Get(userapplicationURL + "/all/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetLatestUserApplication(userId string) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	//entity = domain.UserApplication{userId, "1", time.Now()}
	resp, _ := api.Rest().Get(userapplicationURL + "/latest/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateUserApplication(obj domain.UserApplication) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteUserApplication(obj domain.UserApplication) (domain.UserApplication, error) {
	entity := domain.UserApplication{}
	resp, _ := api.Rest().SetBody(obj).Post(userapplicationURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
