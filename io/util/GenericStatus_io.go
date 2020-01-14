package util

import (
	"errors"
	"obas/api"
	domain "obas/domain/util"
)

const utilStatusURL = api.BASE_URL + "/generics/util/status/"

func GetStatuses()([]domain.GenericStatus, error) {
	entites := []domain.GenericStatus{}
	resp, _ := api.Rest().Get(utilStatusURL + "all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func CreateStatus(entity domain.GenericStatus) (domain.GenericStatus, error) {
	saved := domain.GenericStatus{}
	resp, _ := api.Rest().
		SetBody(entity).
		Post(utilStatusURL + "create")
	if resp.IsError() {
		return saved, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &saved)
	if err != nil {
		return saved, errors.New(resp.Status())
	}
	return saved, nil
}
func GetStatus(id string) (domain.GenericStatus, error) {
	entites := domain.GenericStatus{}
	resp, _ := api.Rest().Get(utilStatusURL + "get/" + id)

	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetIncompleteStatus()(domain.GenericStatus, error) {
	entity := domain.GenericStatus{}
	//entity = domain.GenericStatus{"1", "Incomplete", ""}
	resp, _ := api.Rest().Get(utilStatusURL + "incomplete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(err.Error())
	}
	return entity, nil
}
