package academics

import (
	"errors"
	"obas/api"
	domain "obas/domain/academics"
)

const courseSubjectURL = api.BASE_URL + "/academics/coursesubject"

func CreateCourseSubject(obj domain.CourseSubject) (domain.CourseSubject, error) {
	entity := domain.CourseSubject{}

	resp, _ := api.Rest().SetBody(obj).Post(courseSubjectURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetCourseSubject(courseId, subjectId string) (domain.CourseSubject, error) {
	entity := domain.CourseSubject{}
	resp, _ := api.Rest().Get(courseSubjectURL + "/get/" + courseId + "/" + subjectId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetCourseSubjects(courseId string) ([]domain.CourseSubject, error) {
	entities := []domain.CourseSubject{}
	//entities = append(entities, domain.CourseSubject{courseId, "1"})
	//entities = append(entities, domain.CourseSubject{courseId, "2"})
	resp, _ := api.Rest().Get(courseSubjectURL + "/getsubjects/" + courseId)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetAllCourseSubject() ([]domain.CourseSubject, error) {
	entity := []domain.CourseSubject{}
	resp, _ := api.Rest().Get(courseSubjectURL + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteCourseSubject(obj domain.CourseSubject) (domain.CourseSubject, error) {
	entity := domain.CourseSubject{}
	resp, _ := api.Rest().SetBody(obj).Post(courseSubjectURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
