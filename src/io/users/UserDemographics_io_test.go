package io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/src/domain/users"
	"testing"
)

func TestGetUserDemographics(t *testing.T) {
	result, err := GetUserDemographics()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}

func TestGetUserDemographic(t *testing.T) {
	expected := "25"
	result, err := GetUserDemographic("56")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.UserDemographicsId)

}

func TestCreateUserDemographics(t *testing.T) {
	userDemo := domain.UserDemographics{"56", "25", "86"}
	result, err := CreateUserDemographics(userDemo)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUserDemographics(t *testing.T) {
	userDemo := domain.UserDemographics{"56", "25", "86"}
	result, err := UpdateUserDemographics(userDemo)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestDeleteUserDemographics(t *testing.T) {
	userDemo := domain.UserDemographics{"56", "25", "86"}
	result, err := DeleteUserDemographics(userDemo)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}
