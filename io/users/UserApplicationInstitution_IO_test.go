package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestCreateUserApplicationInstitution(t *testing.T) {
	obj := domain.UserApplicationInstitution{}
	result, err := CreateUserApplicationInstitution(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteUserApplicationInstitution(t *testing.T) {
	obj := domain.UserApplicationInstitution{}
	result, err := DeleteUserApplicationInstitution(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserApplicationInstitutionAllForUser(t *testing.T) {
	result, err := GetUserApplicationInstitutions("")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateUserApplicationInstitution(t *testing.T) {
	obj := domain.UserApplicationInstitution{}
	result, err := UpdateUserApplicationInstitution(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserApplicationInstitutionForAppl(t *testing.T) {
	result, err := GetUserApplicationInstitution("", "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserApplicationInstitutions(t *testing.T) {
	result, err := GetAllUserApplicationInstitutions()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
