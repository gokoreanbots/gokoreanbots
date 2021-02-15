package http

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
)

var baseURL string  = "https://api.koreanbots.dev/v1"


// PostServers Post Servers on Koreanbots
func PostServers(servers int, token string) {
	headers := map[string]string{"Content-Type": "application/json", "token": token}
	client := http.Client{}
	seversJSON := []byte(`{"servers": ` + strconv.Itoa(servers) + `}`)
	req, _ := http.NewRequest("POST", baseURL + "/bots/servers", bytes.NewBuffer(seversJSON))
	for key, value := range headers {
		req.Header.Add(key, value)
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
}