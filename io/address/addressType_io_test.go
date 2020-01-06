package address

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/address"
	"testing"
)

//var entity = domain.AddressType{AddressTypeID: "EMAIL", AddressName: "EMAIL ADDRESS"}

func TestCreateAddressType(t *testing.T) {
	addrType := domain.AddressType{"UBER", "UBER CLIENT"}
	value, err := CreateAddressType(addrType)
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestGetAddressType(t *testing.T) {
	expected := "UBER CLIENT"
	value, err := GetAddressType("UBER")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.AddressName, expected)

}
func TestGetAddressTypes(t *testing.T) {
	value, err := GetAddressTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}

func TestUpdateAddressType(t *testing.T) {
	addrType := domain.AddressType{"UBER", "UBER DRIVER"}
	value, err := UpdateAddressType(addrType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteAddressType(t *testing.T) {
	addrType := domain.AddressType{"UBER", "UBER DRIVER"}
	value, err := DeleteAddressType(addrType)
	assert.Nil(t, err)
	assert.True(t, value)
}
