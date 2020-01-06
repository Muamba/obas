package users

import (
	"errors"
	"obas/api"
)

const userInstitutionURL = api.BASE_URL + "/users/institution/"

//func GetUserInstitutions() ([]domain.UserInstitution, error) {
//	entites := []domain.UserInstitution{}
//	resp, _ := api.Rest().Get(userInstitutionURL + "/institution/all")
//
//	if resp.IsError() {
//		return entites, errors.New(resp.Status())
//	}
//	err := api.JSON.Unmarshal(resp.Body(), &entites)
//	if err != nil {
//		return entites, errors.New(resp.Status())
//	}
//	return entites, nil
//}

//func GetUserInstitution(id string) (domain.UserInstitution, error) {
//	entity := domain.UserInstitution{}
//	resp, _ := api.Rest().Get(userInstitutionURL + "/institution/get/" + id)
//	if resp.IsError() {
//		return entity, errors.New(resp.Status())
//	}
//	err := api.JSON.Unmarshal(resp.Body(), &entity)
//	if err != nil {
//		return entity, errors.New(resp.Status())
//	}
//	return entity, nil
//}

func UpdateUserInstitution(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userInstitutionURL + "/institution/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteUserInstitution(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userInstitutionURL + "/institution/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}
