package location

import (
	"errors"
	"obas/api"
	domain "obas/domain/location"
)

const locationUrl = api.BASE_URL + "/location"

func GetLocations() ([]domain.Location, error) {
	entites := []domain.Location{}
	resp, _ := api.Rest().Get(locationUrl + "/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetParentLocations() ([]domain.Location, error) {
	entites := []domain.Location{}
	//entites = append(entites, domain.Location{"1", "1", "South Africa", "", "", ""})
	resp, _ := api.Rest().Get(locationUrl + "/parents/all")
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetLocationsForParent(locationParentId string) ([]domain.Location, error) {
	entites := []domain.Location{}
	//if locationParentId == "1" {
	//	entites = append(entites, domain.Location{"2", "2", "Western Cape", "", "", locationParentId})
	//	entites = append(entites, domain.Location{"3", "2", "Eastern Cape", "", "", locationParentId})
	//} else if locationParentId == "2" {
	//	entites = append(entites, domain.Location{"4", "3", "Cape Winelands", "", "", locationParentId})
	//	entites = append(entites, domain.Location{"5", "3", "Garden Route", "", "", locationParentId})
	//} else if locationParentId == "4" {
	//	entites = append(entites, domain.Location{"6", "4", "Breede Valley", "", "", locationParentId})
	//	entites = append(entites, domain.Location{"7", "4", "Drakenstein", "", "", locationParentId})
	//}
	resp, _ := api.Rest().Get(locationUrl + "/getforparents/" + locationParentId)
	if resp.IsError() {
		return entites, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entites)
	if err != nil {
		return entites, errors.New(resp.Status())
	}
	return entites, nil
}

func GetLocation(id string) (domain.Location, error) {
	entity := domain.Location{}
	//entity = domain.Location{"6", "4", "Breede Valley", "", "", "4"}
	resp, _ := api.Rest().Get(locationUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateLocation(entity domain.Location) (domain.Location, error) {
	location := domain.Location{}
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationUrl + "/create")
	if resp.IsError() {
		return location, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &location)
	if err != nil {
		return location, errors.New(resp.Status())
	}

	return location, nil
}

func UpdateLocation(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteLocation(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(locationUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
