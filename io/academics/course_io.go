package academics

import (
	"errors"
	"obas/api"
	domain "obas/domain/academics"
)

const academicsURL = api.BASE_URL + "/academics/course"

func SaveCourse(entity domain.Course) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(academicsURL + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func GetCourse(id string) (domain.Course, error) {
	entity := domain.Course{}
	//if id == "2" {
	//	entity = domain.Course{id, "Information Technology", ""}
	//} else if id == "3" {
	//	entity = domain.Course{id, "Accounting", ""}
	//} else if id == "4" {
	//	entity = domain.Course{id, "Chemistry", ""}
	//}
	resp, _ := api.Rest().Get(academicsURL + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateCourse(course domain.Course, token string) (bool, error) {
	resp, _ := api.Rest().
		SetAuthToken(token).
		SetBody(course).
		Post(academicsURL + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func DeleteCourse(course domain.Course) (bool, error) {
	resp, _ := api.Rest().
		SetBody(course).
		Post(academicsURL + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func GetAllCourses() ([]domain.Course, error) {
	entites := []domain.Course{}
	resp, _ := api.Rest().Get(academicsURL + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}
