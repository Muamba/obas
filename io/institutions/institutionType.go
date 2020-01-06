package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionTypeURL = api.BASE_URL + "/institutions/type/"

func CreateInstitutionType(obj domain.InstitutionTypes) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionTypeURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitutionType(id string) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
	resp, _ := api.Rest().Get(institutionTypeURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitutionTypes() ([]domain.InstitutionTypes, error) {
	entity := []domain.InstitutionTypes{}
	//entity = append(entity, domain.InstitutionTypes{"1", "University", ""})
	//entity = append(entity, domain.InstitutionTypes{"2", "College", ""})
	//entity = append(entity, domain.InstitutionTypes{"3", "High School", ""})
	resp, _ := api.Rest().Get(institutionTypeURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteInstitutionType(obj domain.InstitutionTypes) (bool, error) {
	resp, _ := api.Rest().SetBody(obj).Post(institutionTypeURL + "delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}

func UpdateInstitutionType(obj domain.InstitutionTypes) (domain.InstitutionTypes, error) {
	entity := domain.InstitutionTypes{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionTypeURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
