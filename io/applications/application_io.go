package applications

import (
	"errors"
	"fmt"
	"obas/api"
	domain "obas/domain/application"
)

const applicationURL = api.BASE_URL + "/application"

func CreateApplication(obj domain.Application) (domain.Application, error) {
	entity := domain.Application{}
	//entity.Id = "123"
	//entity.ApplicantTypeId = obj.ApplicantTypeId
	//entity.ApplicationTypeId = obj.ApplicationTypeId
	resp, _ := api.Rest().SetBody(obj).Post(applicationURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteApplication(obj domain.Application) (domain.Application, error) {
	entity := domain.Application{}

	resp, _ := api.Rest().SetBody(obj).Post(applicationURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetApplication(id string) (domain.Application, error) {
	entity := domain.Application{}
	//entity = domain.Application{"1", "1", "1"}
	resp, _ := api.Rest().Get(applicationURL + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetApplications() ([]domain.Application, error) {
	entity := []domain.Application{}

	resp, _ := api.Rest().Get(applicationURL + "/all")
	if resp.IsError() {
		fmt.Println("in resp.IsError()")
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		fmt.Println("in err!=ni")
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateApplication(obj domain.Application, token string) (bool, error) {
	resp, _ := api.Rest().SetAuthToken(token).SetBody(obj).Post(applicationURL + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
