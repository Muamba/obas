package location

import (
	"obas/config"
	genericHelper "obas/controllers/misc"
	locationDomain "obas/domain/location"
	locationIO "obas/io/location"
)

/**
Get provinces
*/
func GetProvinces(app *config.Env) ([]locationDomain.Location, genericHelper.PageToast) {
	var provinces []locationDomain.Location
	var alert genericHelper.PageToast
	provinces, err := getProvinces()
	if err != nil {
		app.ErrorLog.Println(err.Error())
		alert = genericHelper.PageToast{genericHelper.DangerAlertStyle, "Could not retrieve provinces!"}
	}
	return provinces, alert
}

func getProvinces() ([]locationDomain.Location, error) {
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
