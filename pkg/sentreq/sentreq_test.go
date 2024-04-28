package sentreq

import (
	"testing"
)

func TestFormDataReq(t *testing.T) {
	url := "https://example.com/upload"
	username := "test_user"
	password := "test_password"
	formData := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	responseData, err := FormDataReq(url, username, password, formData)
	if err != nil {
		t.Errorf("FormDataReq() returned an error: %v", err)
	}

	if len(responseData) == 0 {
		t.Error("FormDataReq() returned an empty response")
	}

	t.Logf("FormDataReq() test passed successfully")
}

func TestGET(t *testing.T) {
	url := "https://example.com/api"
	username := "test_user"
	password := "test_password"
	queryData := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	responseData, err := GET(url, username, password, queryData)
	if err != nil {
		t.Errorf("GET() returned an error: %v", err)
	}

	if len(responseData) == 0 {
		t.Error("GET() returned an empty response")
	}

	t.Logf("GET() test passed successfully")
}
