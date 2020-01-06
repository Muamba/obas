package institutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/institutions"
	"testing"
)

func TestCreateInstitutionCourse(t *testing.T) {
	obj := domain.InstitutionCourse{"00012", "0040"}
	resp, err := CreateInstitutionCourse(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestDeleteInstitutionCourse(t *testing.T) {
	obj := domain.InstitutionCourse{"00012", "0040"}
	resp, err := DeleteInstitutionCourse(obj)
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestGetInstitutionCourse(t *testing.T) {
	resp, err := GetInstitutionCourse("0040")
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestGetInstitutionCourses(t *testing.T) {
	resp, err := GetInstitutionCourses()
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
func TestReadInstitutionCourse(t *testing.T) {
	resp, err := ReadInstitutionCourse("00012", "0040")
	assert.NotNil(t, resp)
	fmt.Println(" The Results", resp)
	assert.Nil(t, err)
}
