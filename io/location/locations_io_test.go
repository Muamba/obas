package location

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/location"
	"testing"
)

func TestGetLocations(t *testing.T) {
	value, err := GetLocations()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}

func TestGetLocation(t *testing.T) {
	expected := "WC"
	value, err := GetLocation("53")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.Name, expected)
}

func TestCreateSchool(t *testing.T) {
	loc := domain.Location{}
	value, err := CreateLocation(loc)
	assert.Nil(t, err)
	assert.NotNil(t, value)
}

func TestUpdateDocument(t *testing.T) {
	loc := domain.Location{}
	value, err := UpdateLocation(loc)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteDocument(t *testing.T) {
	loc := domain.Location{}
	value, err := DeleteLocation(loc)
	assert.Nil(t, err)
	assert.True(t, value)
}
func TestGetTowns(t *testing.T) {
	value, err := GetTowns("loc")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}
