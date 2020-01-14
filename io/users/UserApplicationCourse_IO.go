package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userApplicationCourseURL = api.BASE_URL + "/user/application/course/"

func GetAllUserApplicationCourses() (domain.UserApplicationCourse, error) {
	entity := domain.UserApplicationCourse{}
	resp, _ := api.Rest().Get(userApplicationCourseURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetUserApplicationCourses(userId string) ([]domain.UserApplicationCourse, error) {
	entity := []domain.UserApplicationCourse{}
	resp, _ := api.Rest().Get(userApplicationCourseURL + "allforuser/" + userId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserApplicationCourse(userId, applicationId string) (domain.UserApplicationCourse, error) {
	entity := domain.UserApplicationCourse{}
	//entity = domain.UserApplicationCourse{userId, applicationId, "4"}
	resp, _ := api.Rest().Get(userApplicationCourseURL + "getforapplication/" + userId + "/" + applicationId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateUserApplicationCourse(obj domain.UserApplicationCourse) (domain.UserApplicationCourse, error) {
	entity := domain.UserApplicationCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(userApplicationCourseURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateUserApplicationCourse(obj domain.UserApplicationCourse) (domain.UserApplicationCourse, error) {
	entity := domain.UserApplicationCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(userApplicationCourseURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserApplicationCourse(obj domain.UserApplicationCourse) (domain.UserApplicationCourse, error) {
	entity := domain.UserApplicationCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(userApplicationCourseURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
