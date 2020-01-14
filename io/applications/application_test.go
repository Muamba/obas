package applications

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/application"
	"testing"
)

func TestCreateApplication(t *testing.T) {
	obj := domain.Application{"0022", "ERRY-OPBMV", "AEEP-OCOTZ"}
	resp, err := CreateApplication(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestGetApplication(t *testing.T) {
	resp, err := GetApplication("ORXX-BFZBW")
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestDeleteApplication(t *testing.T) {
	obj := domain.Application{"0022", "ERRY-OPBMV", "AEEP-OCOTZ"}
	resp, err := DeleteApplication(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestUpdateApplicatio(t *testing.T) {
	obj := domain.Application{"0022", "ERRY-OPBMV", "AEEP-OCOTZ"}
	resp, err := UpdateApplication(obj, "")
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestGetApplications(t *testing.T) {
	resp, err := GetApplications()
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
