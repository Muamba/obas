package academics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/academics"
	"testing"
)

func TestCreateCourseSubject(t *testing.T) {
	obj := domain.CourseSubject{"0001", "A0000"}
	result, err := CreateCourseSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, obj, result)
}

func TestGetCourseSubject(t *testing.T) {
	result, err := GetCourseSubject("0001", "002")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}

func TestDeleteCourseSubject(t *testing.T) {
	obj := domain.CourseSubject{"AAAD-OTKLY", "AMNN-00MWQ"}
	result, err := DeleteCourseSubject(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}

func TestGetAllCourseSubject(t *testing.T) {
	result, err := GetAllCourseSubject()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.NotNil(t, result)
}
