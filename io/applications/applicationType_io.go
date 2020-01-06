package applications

import (
	"errors"
	"obas/api"
	domain "obas/domain/application"
)

const applicationTypeUrl = api.BASE_URL + "/application"

func GetApplicationTypes() ([]domain.ApplicationType, error) {
	entites := []domain.ApplicationType{}
	//entites = append(entites, domain.ApplicationType{"1", "Motsepe Bursary", ""})
	resp, _ := api.Rest().Get(applicationTypeUrl + "/type/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetApplicationType(id string) (domain.ApplicationType, error) {
	entity := domain.ApplicationType{}
	resp, _ := api.Rest().Get(applicationTypeUrl + "/type/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateApplicationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTypeUrl + "/type/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateApplicationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTypeUrl + "/type/update")
	if resp.IsError() {
		return true, nil
	}
	return true, nil
}

func DeleteApplicationType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(applicationTypeUrl + "/type/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
