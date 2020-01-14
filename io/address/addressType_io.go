package address

import (
	"errors"
	"obas/api"
	domain "obas/domain/address"
)

const addressTypeUrl = api.BASE_URL + "/address"

type AddressType domain.AddressType

func GetAddressTypes() ([]AddressType, error) {
	entities := []AddressType{}
	addT1 := AddressType{"123", "Physical Address"}
	addT2 := AddressType{"246", "Postal Address"}

	allAdd := []AddressType{addT1, addT2}

	entities = allAdd
	//resp, _ := api.Rest().Get(addressTypeUrl + "/all")
	//
	//if resp.IsError() {
	//	return entities, errors.New(resp.Status())
	//}
	//err := api.JSON.Unmarshal(resp.Body(), &entities)
	//if err != nil {
	//	return entities, errors.New(resp.Status())
	//}
	return entities, nil
}

func GetAddressType(id string) (domain.AddressType, error) {
	entity := domain.AddressType{}
	resp, _ := api.Rest().Get(addressTypeUrl + "/get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func CreateAddressType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressTypeUrl + "/create")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func UpdateAddressType(entity interface{}) (bool, error) {
	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressTypeUrl + "/update")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}

	return true, nil
}

func DeleteAddressType(entity interface{}) (bool, error) {

	resp, _ := api.Rest().
		SetBody(entity).
		Post(addressTypeUrl + "/delete")
	if resp.IsError() {
		return false, errors.New(resp.Status())
	}
	return true, nil
}
