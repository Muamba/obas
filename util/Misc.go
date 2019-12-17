package util

import (
	locationDomain "obas/domain/location"
	locationIO "obas/io/location"
)

func GetProvinces() ([]locationDomain.Location, error) {
	southAfricaId, err := getCountryId()
	if err != nil {
		return nil, err
	} else {
		provinces, err := locationIO.GetLocationsForParent(southAfricaId)
		if err != nil {
			return nil, err
		} else {
			return provinces, nil
		}
	}
}

func getCountryId() (string, error) {
	var id string
	countries, err := locationIO.GetParentLocations()
	if err != nil {
		return id, err
	} else {
		if len(countries) > 0 {
			id = countries[0].LocationId
		}
		return id, nil
	}
}
