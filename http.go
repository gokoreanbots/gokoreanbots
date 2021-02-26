package gokoreanbots

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func post(url string, headers map[string]string, jsonData []byte) error {
	client := http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	res, err := client.Do(req)
	if err != nil {
		return ErrFailToPost
	}
	defer res.Body.Close()
	return nil
}

func get(url string, headers map[string]string) (string, error) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	responseByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	return string(responseByte), nil
}
