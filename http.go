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
	resp, _ := client.Do(req)
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusTooManyRequests:
		return ErrRateLimited
	case http.StatusUnauthorized:
		return ErrUnauthorized
	}
	defer resp.Body.Close()
	return nil
}

func get(url string, headers map[string]string) (response string, err error) {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	resp, _ := client.Do(req)
	switch resp.StatusCode {
	case http.StatusOK:
		responseByte, _ := ioutil.ReadAll(resp.Body)
		response = string(responseByte)
		err = nil
	case http.StatusBadRequest:
		responseByte, _ := ioutil.ReadAll(resp.Body)
		response = string(responseByte)
		err = ErrUnauthorized
	}
	defer resp.Body.Close()
	return
}
