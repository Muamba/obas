package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userTertiaryInstitutionURL = api.BASE_URL + "/users/institution/tertiary/"

func CreateUserTertiaryInstitution(entity domain.UserTertiaryInstitution) (domain.UserTertiaryInstitution, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userTertiaryInstitutionURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserTertiaryInstitutionForApplication(userId, applicationId string) (domain.UserTertiaryInstitution, error) {
	entity := domain.UserTertiaryInstitution{}
	//entity = domain.UserTertiaryInstitution{userId, applicationId, "2", 62889.09}
	resp, _ := api.Rest().Get(userTertiaryInstitutionURL + "getforapplication/" + userId + "/" + applicationId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
