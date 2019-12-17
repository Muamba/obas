package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userTownURL = api.BASE_URL + "/users/town"

func CreateUserTown(entity domain.UserTown) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userTownURL + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func GetUserTown(userId string) (domain.UserTown, error) {
	entity := domain.UserTown{}
	//entity = domain.UserTown{userId, "6"}
	resp, _ := api.Rest().Get(userTownURL + "/get/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateUserTown(entity domain.UserTown, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(entity).
		Post(userTownURL + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
