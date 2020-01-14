package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userApplicationInstitutionURL = api.BASE_URL + "user/application/institution/"

func GetAllUserApplicationInstitutions() ([]domain.UserApplicationInstitution, error) {
	entity := []domain.UserApplicationInstitution{}
	resp, _ := api.Rest().Get(userApplicationInstitutionURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserApplicationInstitutions(userId string) ([]domain.UserApplicationInstitution, error) {
	entity := []domain.UserApplicationInstitution{}
	resp, _ := api.Rest().Get(userApplicationInstitutionURL + "allforuser/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserApplicationInstitution(userId, applicationId string) (domain.UserApplicationInstitution, error) {
	entity := domain.UserApplicationInstitution{}
	//entity = domain.UserApplicationInstitution{userId, applicationId, "2"}
	resp, _ := api.Rest().Get(userApplicationInstitutionURL + "getforapplication/" + userId + "/" + applicationId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserApplicationInstitution(obj domain.UserApplicationInstitution) (domain.UserApplicationInstitution, error) {
	entity := domain.UserApplicationInstitution{}
	resp, _ := api.Rest().SetBody(obj).Post(userApplicationInstitutionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateUserApplicationInstitution(obj domain.UserApplicationInstitution) (domain.UserApplicationInstitution, error) {
	entity := domain.UserApplicationInstitution{}
	resp, _ := api.Rest().SetBody(obj).Post(userApplicationInstitutionURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserApplicationInstitution(obj domain.UserApplicationInstitution) (domain.UserApplicationInstitution, error) {
	entity := domain.UserApplicationInstitution{}
	resp, _ := api.Rest().SetBody(obj).Post(userApplicationInstitutionURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
