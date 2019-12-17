package location

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/location"
	"testing"
)

func TestGetLocationTypes(t *testing.T) {
	value, err := GetLocationTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}

func TestGetLocationType(t *testing.T) {
	expected := "Western"
	value, err := GetLocationType("26")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Name, expected)
}

func TestCreateLocationType(t *testing.T) {
	locType := domain.LocationType{"26", "Western", "789"}
	value, err := CreateLocationType(locType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}

func TestUpdateLocationType(t *testing.T) {
	locType := domain.LocationType{"27", "Western", "789"}
	value, err := UpdateLocationType(locType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteLocationType(t *testing.T) {
	locType := domain.LocationType{"26", "Western", "789"}
	value, err := DeleteLocationType(locType)
	assert.Nil(t, err)
	assert.True(t, value)
}
