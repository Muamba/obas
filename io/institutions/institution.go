package institutions

import (
	"errors"
	"obas/api"
	domain "obas/domain/institutions"
)

const institutionURL = api.BASE_URL + "/institutions"

func CreateInstitution(obj domain.Institution) (domain.Institution, error) {
	entity := domain.Institution{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionURL + "/create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteInstitution(obj domain.Institution) (domain.Institution, error) {
	entity := domain.Institution{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionURL + "/delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitution(id string) (domain.Institution, error) {
	entity := domain.Institution{}
	//if id == "1" {
	//	entity = domain.Institution{id, "3", "Sea Point High School"}
	//} else if id == "2" {
	//	entity = domain.Institution{id, "1", "University of Western Cape"}
	//} else if id == "3" {
	//	entity = domain.Institution{id, "1", "Stellenbosch"}
	//}
	resp, _ := api.Rest().Get(institutionURL + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func GetInstitutionsByType(institutionTypeId string) ([]domain.Institution, error) {
	entities := []domain.Institution{}
	//if institutionTypeId == "1" {
	//	entities = append(entities, domain.Institution{"1", institutionTypeId, "CPUT"})
	//	entities = append(entities, domain.Institution{"2", institutionTypeId, "UWC"})
	//	entities = append(entities, domain.Institution{"3", institutionTypeId, "Stellenbosch"})
	//}
	resp, _ := api.Rest().Get(institutionURL + "/getforinstitutiontype/" + institutionTypeId)
	if resp.IsError() {
		return entities, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entities)
	if err != nil {
		return entities, errors.New(resp.Status())
	}
	return entities, nil
}

func GetInstitutions() ([]domain.Institution, error) {
	entity := []domain.Institution{}
	//entity = append(entity, domain.Institution{"1", "3", "Sea Point High School"})
	//entity = append(entity, domain.Institution{"2", "1", "UWC"})
	//entity = append(entity, domain.Institution{"3", "1", "Stellenbosch"})
	resp, _ := api.Rest().Get(institutionURL + "/all")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateInstitution(obj domain.Institution) (domain.Institution, error) {
	entity := domain.Institution{}
	resp, _ := api.Rest().SetBody(obj).Post(institutionURL + "/update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
