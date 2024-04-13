package sentreq

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
)

// Sent post type request, use basic auth
func FormDataReq(url, username, password string, formData map[string]string) (map[string]interface{}, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range formData {
		_ = writer.WriteField(k, v)
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(username, password)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseData map[string]interface{}

	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}

// Sent get type request, use basic auth
func GET(url, username, password string, queryData map[string]string) (map[string]interface{}, error) {
	if queryData != nil {
		url += "?"
		for k, v := range queryData {
			url += k + "=" + v + "&"
		}
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(username, password)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var responseData map[string]interface{}

	err = json.Unmarshal(responseBody, &responseData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
