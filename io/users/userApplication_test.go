package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/users"
	"testing"
	"time"
)

func TestCreateUserApplication(t *testing.T) {

	obj := domain.UserApplication{"0001", "0303445", time.Now()}
	resp, err := CreateUserApplication(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestDeleteUserApplication(t *testing.T) {
	obj := domain.UserApplication{"espoirditekemena@gmail.com", "AARR-6AHMZ", time.Now()}
	resp, err := DeleteUserApplication(obj)
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestGetUserApplication(t *testing.T) {
	resp, err := GetUserApplication("0001", "0303445")
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
func TestGetUserApplications(t *testing.T) {
	resp, err := GetUserApplications("espoirditekemena@gmail.com")
	if len(resp) <= 0 {
		fmt.Println(" The Results in if", resp)
		return
	}
	assert.Nil(t, err)
	fmt.Println(" The Results", resp)
	assert.NotNil(t, resp)
}
