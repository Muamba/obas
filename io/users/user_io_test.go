package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
	"time"
)

func TestGetUsers(t *testing.T) {
	result, err := GetUsers()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}

func TestGetUser(t *testing.T) {
	expected := "JEAN"
	result, err := GetUser("7896541230")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.FirstName)

}

func TestCreateUser(t *testing.T) {
	userC := domain.User{"m@gt.com", "", "JEAN", "PAUL", "MATUTO", time.Time{}}
	result, err := CreateUser(userC)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)

}

func TestUpdateUser(t *testing.T) {
	userC := User{"m@gt.com", "", "JEAN", "PAUL", "MATUTO", time.Time{}}
	result, err := UpdateUser(userC, "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, result)
}

func TestDeleteUser(t *testing.T) {
	userC := domain.User{"m@gt.com", "", "JEAN", "PAUL", "MATUTO", time.Time{}}
	result, err := DeleteUser(userC)
	assert.Nil(t, err)
	assert.True(t, result)

}
