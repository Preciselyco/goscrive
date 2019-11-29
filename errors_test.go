package scrive_test

/*
func TestParseResponseErrorOk(t *testing.T) {
	errBody := []byte(`
		{
			"error_type": "resource_not_found",
			"error_message": "The resource was not found",
			"http_code": 404
		}
	`)
	se, err := scrive.ParseResponseError(errBody)
	assert.NotNil(t, se)
	assert.Nil(t, err)
	assert.Equal(t, se.ErrorType, "resource_not_found")
	assert.Equal(t, se.ErrorMessage, "The resource was not found")
	assert.Equal(t, se.HttpCode, 404)
}

func TestParseResponseErrorNotOk(t *testing.T) {
	errBody := []byte(`{}`)
	se, err := scrive.ParseResponseError(errBody)
	assert.Nil(t, se)
	assert.NotNil(t, err)
}
*/
