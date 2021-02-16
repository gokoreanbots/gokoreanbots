package gokoreanbots

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// Post somthing
func Post(url string, headers map[string]string, jsonData []byte) error{
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

// Get something
func Get(url string, headers map[string]string) string {
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	res, _ := client.Do(req)
	responseByte, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return string(responseByte)
}
