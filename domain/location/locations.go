package domain

type LocationType struct {
	LocationTypeId string `json:"locationTypeId"`
	Name           string `json:"name"`
	Code           string `json:"code"`
}

type Location struct {
	LocationId       string `json:"locationId"`
	LocationTypeId   string `json:"locationTypeId"`
	Name             string `json:"name"`
	Latitude         string `json:"latitude"`
	Longitude        string `json:"longitude"`
	LocationParentId string `json:"locationParentId"`
}
