package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionLocationURL = api.BASE_URL + "/institutions/location/"

func CreateInstitutionLocation(obj domain.InstitutionLocation) (domain.InstitutionLocation, error) {
	entity := domain.InstitutionLocation{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionLocationURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteInstitutionLocation(obj domain.InstitutionLocation) (domain.InstitutionLocation, error) {
	entity := domain.InstitutionLocation{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionLocationURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadInstitutionLocation(id string) (domain.InstitutionLocation, error) {
	entity := domain.InstitutionLocation{}
	resp, _ := api.Rest().Get(institutionLocationURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadInstitutionLocations() ([]domain.InstitutionLocation, error) {
	entity := []domain.InstitutionLocation{}
	resp, _ := api.Rest().Get(institutionLocationURL + "all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitutionsInLocation(locationId string) ([]domain.InstitutionLocation, error) {
	entity := []domain.InstitutionLocation{}
	//entity = append(entity, domain.InstitutionLocation{"1", locationId, "", ""})
	//entity = append(entity, domain.InstitutionLocation{"2", locationId, "", ""})
	resp, _ := api.Rest().Get(institutionLocationURL + "getforlocation/" + locationId)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateInstitutionLocation(obj domain.InstitutionLocation) (domain.InstitutionLocation, error) {
	entity := domain.InstitutionLocation{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionLocationURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
