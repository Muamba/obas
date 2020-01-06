package domain

type Institution struct {
	Id                string `json:"id"`
	InstitutionTypeId string `json:"institutionTypeId"`
	Name              string `json:"name"`
}
type InstitutionTypes struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type InstitutionCourse struct {
	InstitutionId string `json:"institutionId"`
	CourseId      string `json:"courseId"`
}

type InstitutionLocation struct {
	InstitutionId string `json:"institutionId"`
	LocationId    string `json:"locationId"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
}

type InstitutionAddress struct {
	InstitutionId string `json:"institutionId"`
	AddressTypeId string `json:"addressTypeId"`
	Address       string `json:"address"`
	PostalCode    string `json:"postalCode"`
}
