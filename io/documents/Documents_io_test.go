package documents

/**
var doc = documents{"test@test.go", "25", "FR", "MATRIC", "DS", "QA", "", "NONE"}
var token=""
func TestDocuments(t *testing.T) {
	value, err := GetDocuments()
	assert.Nil(t, err)
	fmt.Println(" The Results", value)
	assert.True(t, len(value) > 0)
}

func TestGetDocument(t *testing.T) {
	expected := doc
	value, err := GetDocument(doc.DocumentId)
	assert.Nil(t, err)
	assert.Equal(t, value, expected)
}

func TestCreateDocument(t *testing.T) {
	value, err := CreateDocument(,token)
	assert.Nil(t, err)
	assert.True(t, value)
}

func TestUpdateDocument(t *testing.T) {
	var expected = "MATRIC"
	var doc = domain.Documents{"FG", "27", "FR", "MATRIC", "DS", "QA", "", "NONE"}
	result, err := UpdateDocument(doc)
	assert.Nil(t, err)
	assert.True(t, result)
	value, err := GetDocument(doc.DocumentsId)
	assert.Equal(t, expected, value.Description)
}

func TestDeleteDocument(t *testing.T) {
	value, err := DeleteDocument(doc)
	assert.Nil(t, err)
	assert.True(t, value)
}*/
