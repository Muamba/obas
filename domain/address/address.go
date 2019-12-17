package domain

type AddressType struct {
	AddressTypeID string `json:"addressTypeID"`
	AddressName   string `json:"addressName"`
}

type ContactType struct {
	ContactTypeId string `json:"contactTypeId"`
	Name          string `json:"name"`
}
