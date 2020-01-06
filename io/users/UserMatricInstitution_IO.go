package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userMatricInstitutionURL = api.BASE_URL + "/users/institution/matric/"

func CreateUserMatricInstitution(entity domain.UserMatricInstitution) (domain.UserMatricInstitution, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userMatricInstitutionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserMatricInstitution(userId string) (domain.UserMatricInstitution, error) {
	entity := domain.UserMatricInstitution{}
	//entity = domain.UserMatricInstitution{userId, "1"}
	resp, _ := api.Rest().Get(userMatricInstitutionURL + "get/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
