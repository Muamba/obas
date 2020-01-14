package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestCreateUserMatricInstitution(t *testing.T) {
	obj := domain.UserMatricInstitution{}
	result, err := CreateUserMatricInstitution(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserMatricInstitution(t *testing.T) {
	result, err := GetUserMatricInstitution("")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
