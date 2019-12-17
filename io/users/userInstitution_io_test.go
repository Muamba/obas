package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestGetUserIntstitutions(t *testing.T) {
	value, err := GetUserInstitutions()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetUserIntstitution(t *testing.T) {
	expected := "CUT"
	value, err := GetUserInstitution("87")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value.UserId, expected)
}

func TestCreateUserIntstitution(t *testing.T) {
	uInt := domain.UserInstitution{"86", "UNISA", 0,true}
	value, err := CreateUserInstitution(uInt)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
}

func TestUpdateUserIntstitution(t *testing.T) {
	uInt := domain.UserInstitution{"86", "UNISA", 0, false}
	value, err := DeleteUserInstitution(uInt)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteUserIntstitution(t *testing.T) {
	uInt := domain.UserInstitution{"86", "UNISA", 0,false}
	value, err := DeleteUserInstitution(uInt)
	assert.Nil(t, err)
	assert.True(t, value)
}
