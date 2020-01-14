package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userTertiaryCourseURL = api.BASE_URL + "/users/institution/tertiary/course/"

func CreateUserTertiaryCourse(entity domain.UserTertiaryCourse) (domain.UserTertiaryCourse, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userTertiaryCourseURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserTertiaryCourseForApplication(userId, applicationId string) (domain.UserTertiaryCourse, error) {
	entity := domain.UserTertiaryCourse{}
	//entity = domain.UserTertiaryCourse{userId, applicationId, "2"}
	resp, _ := api.Rest().Get(userTertiaryCourseURL + "getforapplication/" + userId + "/" + applicationId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
