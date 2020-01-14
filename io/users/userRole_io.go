package users

import (
	"errors"
	"obas/api"
	userDomain "obas/domain/users"
)

const userRoleUrl = api.BASE_URL + "/users"

func GetUserRoles() ([]userDomain.UserRole, error) {
	entites := []userDomain.UserRole{}
	resp, _ := api.Rest().Get(userRoleUrl + "/role/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetUserRole(id string) (userDomain.UserRole, error) {
	entity := userDomain.UserRole{}
	resp, _ := api.Rest().Get(userRoleUrl + "/role/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRoleUrl + "/role/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateUserRole(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRoleUrl + "/role/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func DeleteUserRole(entity userDomain.UserRole) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userRoleUrl + "/role/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func GetUserRoleWithUserId(userId string) (userDomain.UserRole, error) {
	entity := userDomain.UserRole{}
	resp, _ := api.Rest().Get(userRoleUrl + "/role/getforuser/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
