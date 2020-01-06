package users

import (
	"errors"
	"obas/api"
	domain "obas/domain/users"
)

const userTertiarySubjectURL = api.BASE_URL + "/users/institution/tertiary/subject/"

func CreateUserTertiarySubject(entity domain.UserTertiarySubject) (domain.UserTertiarySubject, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(userTertiarySubjectURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetUserTertiarySubjects(userId, applicationId string) ([]domain.UserTertiarySubject, error) {
	entities := []domain.UserTertiarySubject{}
	//entities = append(entities, domain.UserTertiarySubject{userId, applicationId, "2", 58.52})
	resp, _ := api.Rest().Get(userTertiarySubjectURL + "allforapplication/" + userId + "/" + applicationId)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}
