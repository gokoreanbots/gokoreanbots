package gokoreanbots

import (
	"encoding/json"
	"fmt"
	"testing"
)


func TestGetBots(t *testing.T) {
	var kbclient Client
	fmt.Println(kbclient.GetBots(1))
}

func TestJSON(t *testing.T) {
	var parsed interface{}
	jsonString := "1"
	jsonString2 := "\"asdf\""
	err := json.Unmarshal([]byte(jsonString), &parsed)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsed)
	err = json.Unmarshal([]byte(jsonString2), &parsed)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parsed)
}
