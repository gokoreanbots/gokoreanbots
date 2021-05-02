package gokoreanbots

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestClient_GetBots(t *testing.T) {
	var client Client
	fmt.Println(client.GetBots(1))
}

func TestGet(t *testing.T) {
	resp, err := get(baseURL+"/bots/get?page="+strconv.Itoa(34), nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
}

func TestClient_SearchBots(t *testing.T) {
	var client Client
	fmt.Println(client.SearchBots("초록", 1))
	fmt.Println(client.SearchBots("한마루", 1))
}

func TestClient_GetBotsByCategory(t *testing.T) {
	var client Client
	fmt.Println(client.GetBotsByCategory("뮤직", 1))
}

func TestClient_GetBot(t *testing.T) {
	var client Client
	fmt.Println(client.GetBot("770246143652397069").Name)
	fmt.Println(client.GetBot("687886541619462293").Name)
}
