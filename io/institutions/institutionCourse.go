package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutioncourseURL = api.BASE_URL + "/institutions"

func ReadInstitutionCourse(institutionId, courseId string) (domain.InstitutionCourse, error) {
	entity := domain.InstitutionCourse{}
	resp, _ := api.Rest().Get(institutioncourseURL + "/course/get/" + institutionId + "/" + courseId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetAllInstitutionCourse() ([]domain.InstitutionCourse, error) {
	entity := []domain.InstitutionCourse{}
	resp, _ := api.Rest().Get(institutioncourseURL + "/course/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteInstitutionCourse(obj domain.InstitutionCourse) (bool, error) {
	resp, _ := api.Rest().SetBody(obj).Post(institutioncourseURL + "/course/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func CreateInstitutionCourse(obj domain.InstitutionCourse) (domain.InstitutionCourse, error) {
	entity := domain.InstitutionCourse{}
	resp, _ := api.Rest().SetBody(obj).Post(institutioncourseURL + "/course/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitutionCourses(institutionId string) ([]domain.InstitutionCourse, error) {
	entities := []domain.InstitutionCourse{}
	//if institutionId == "1" {
	//	entities = append(entities, domain.InstitutionCourse{institutionId, "1"})
	//} else {
	//	entities = append(entities, domain.InstitutionCourse{institutionId, "2"})
	//	entities = append(entities, domain.InstitutionCourse{institutionId, "3"})
	//	entities = append(entities, domain.InstitutionCourse{institutionId, "4"})
	//}
	resp, _ := api.Rest().Get(institutioncourseURL + "/course/getcourses/" + institutionId)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
