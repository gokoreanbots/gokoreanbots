package gokoreanbots

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func post(url string, headers *map[string]string, jsonData []byte) error {
	client := http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if headers != nil {
		for key, value := range *headers {
			req.Header.Add(key, value)
		}
	}
	for tries := 0; tries < 5; tries++ {
		resp, _ := client.Do(req)
		switch resp.StatusCode {
		case http.StatusOK:
			return nil
		case http.StatusTooManyRequests:
			remainLimit, _ := strconv.Atoi(resp.Header.Get("x-ratelimit-remaining"))
			log.Printf("[GoKOREANBOTS] rate limited. retry after %d seconds", remainLimit)
			time.Sleep(time.Second * time.Duration(time.Now().Unix()-int64(remainLimit)))
			continue
		case http.StatusUnauthorized:
			return ErrUnauthorized
		case http.StatusBadRequest:
			return ErrBadRequest
		}
		err := resp.Body.Close()
		if err != nil {
			return nil
		}
		return nil
	}
	return ErrTooManyRequests
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
		responseByte, _ := io.ReadAll(resp.Body)
		response = string(responseByte)
		err = nil
	case http.StatusBadRequest:
		responseByte, _ := io.ReadAll(resp.Body)
		response = string(responseByte)
		err = ErrUnauthorized
	case http.StatusTooManyRequests:
		responseByte, _ := io.ReadAll(resp.Body)
		response = string(responseByte)
		err = ErrTooManyRequests
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	return
}
