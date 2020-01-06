package academics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/academics"
	"testing"
)

func TestCreateSubject(t *testing.T) {
	obj := domain.Subject{"0000", "adp3"}
	result, err := CreateSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestDeleteSubject(t *testing.T) {
	obj := domain.Subject{"0000", "adp3"}
	result, err := DeleteSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetSubject(t *testing.T) {
	result, err := GetCourseSubject("0000", "")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestGetSubjects(t *testing.T) {
	result, err := GetSubjects()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
func TestUpdateSubject(t *testing.T) {
	obj := domain.Subject{"0000", "adp3"}
	result, err := UpdateSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
