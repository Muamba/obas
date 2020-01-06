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
	value, err := GetLocation("CTTE-4OILY")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	fmt.Println(" The Results", value.Name)
	assert.NotNil(t, value)
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

func TestGetLocationsForParent(t *testing.T) {
	value, err := GetLocationsForParent("AAIL-YLRUW")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.NotNil(t, value)
}
