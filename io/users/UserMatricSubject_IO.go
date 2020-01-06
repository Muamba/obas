package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userMatricSubjectURL = api.BASE_URL + "/users/institution/matric/subject"

func CreateUserMatricSubject(entity domain.UserMatricSubject) (domain.UserMatricSubject, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userMatricSubjectURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserMatricSubjects(userId string) ([]domain.UserMatricSubject, error) {
	entities := []domain.UserMatricSubject{}
	//entities = append(entities, domain.UserMatricSubject{userId, "1", 34.78})
	//entities = append(entities, domain.UserMatricSubject{userId, "2", 74.78})
	resp, _ := api.Rest().Get(userMatricSubjectURL + "allforuser/" + userId)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
