package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
)

func TestGetUserApplicationCourseAllForUser(t *testing.T) {
	result, err := GetUserApplicationCourses("")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestCreateUserApplicationCourse(t *testing.T) {
	obj := domain.UserApplicationCourse{}
	result, err := CreateUserApplicationCourse(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteUserApplicationCourse(t *testing.T) {
	obj := domain.UserApplicationCourse{}
	result, err := DeleteUserApplicationCourse(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserApplicationCourses(t *testing.T) {
	result, err := GetAllUserApplicationCourses()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateUserApplicationCourse(t *testing.T) {
	obj := domain.UserApplicationCourse{}
	result, err := UpdateUserApplicationCourse(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetUserApplicationCourseForAppl(t *testing.T) {
	result, err := GetUserApplicationCourse("", "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
