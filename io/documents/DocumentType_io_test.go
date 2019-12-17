package documents

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	domain "obas/domain/documents"
	"testing"
)

var token = "eyJraWQiOiJURVNUX1BIUkFTRSIsImFsZyI6IkVTMjU2In0.eyJpc3MiOiJIQVNIQ09ERS5aTSIsImF1ZCI6IlNJVEVVU0VSUyIsImV4cCI6MTU3NjE1NTIxOSwianRpIjoiMkhkT0NyRmM0SHVkZXVyeURmSHZsZyIsImlhdCI6MTU3NjA2ODgxOSwibmJmIjoxNTc2MDY4Njk5LCJzdWIiOiJTaXRlIEFjY2VzcyIsImVtYWlsIjoiZXNwb2lyZGl0ZWtlbWVuYUBnbWFpbC5jb20iLCJyb2xlIjoiU1RSMDAxIn0.3SoiDBLI-ubU7ArWjPuMKh36aVAA6Cm0HfVdnP_ta3pMmGVCdFXEYHw1WMP_lNJS7MUPZlNp9ISfQWhqThhPcQ"

func TestGetDocumentTypes(t *testing.T) {
	value, err := GetDocumentTypes()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetDocumentsType(t *testing.T) {
	expected := "COURSES"
	value, err := GetDocument("2")
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.Equal(t, value, expected)
}

func TestCreateDocumentTypes(t *testing.T) {
	//docType := domain.DocumentType{"2", "COURSES"}
	docType := DocumentType{"2", "COURSES"}
	value, err := CreateDocument(docType, token)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocumentTypes(t *testing.T) {
	docType := domain.DocumentType{"2", "COURSES"}
	value, err := UpdateDocument(docType)
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, value)
}

func TestDeleteDocumentTypes(t *testing.T) {
	docType := domain.DocumentType{"2", "COURSES"}
	value, err := DeleteDocumentType(docType)
	assert.Nil(t, err)
	assert.True(t, value)
}
