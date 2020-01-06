package academics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCourse(t *testing.T) {
	expected := "Test Course"
	result, err := GetCourse("AACI-XJJMU")
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.Equal(t, expected, result.CourseName)
}

func TestSaveCourse(t *testing.T) {
	course := Course{"123", "Test Course", "This is a test course"}
	result, err := SaveCourse(course)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestUpdateCourse(t *testing.T) {
	result, err := GetCourse("123")
	assert.Nil(t, err)
	result.CourseName = "New Test Course"
	r, err := UpdateCourse(result, "")
	assert.Nil(t, err)
	assert.True(t, r)
}

func TestDeleteCourse(t *testing.T) {
	r, err := GetCourse("123")
	assert.Nil(t, err)
	result, err := DeleteCourse(r)
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestGetAllCourses(t *testing.T) {
	result, err := GetAllCourses()
	assert.Nil(t, err)
	fmt.Println(" The Results", result)
	assert.True(t, len(result) > 0)
}
