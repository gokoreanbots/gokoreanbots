package gokoreanbots

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
)

func TestClient_GetBots(t *testing.T) {
	var client Client
	fmt.Println(client.GetBots(1))
}

func getBots(ch chan Bots, client Client, page int) {
	fmt.Println("req to koreanbots page " + strconv.Itoa(page))
	bots := client.GetBots(page)
	ch <- *bots
}

func TestGetAllBotsWithGoroutine(t *testing.T) {
	var client Client
	resp, err := get(baseURL+"/bots/get?page="+strconv.Itoa(34), nil)
	var respJSON getBotsResponse
	err = json.Unmarshal([]byte(resp), &respJSON)
	if err != nil {
		return
	}
	var botList Bots
	ch := make(chan Bots)
	for page := 0; page < respJSON.TotalPage+1; page++ {
		go getBots(ch, client, page+1)
	}
	for page := 1; page < respJSON.TotalPage; page++ {
		botList = append(botList, <-ch...)
	}
	for idx, bot := range botList {
		fmt.Printf("%d. %s (%d서버/하트 %d개)\n", idx+1, bot.Name, bot.Servers, bot.Votes)
	}
}

func TestGetBotsWithoutGoroutine(t *testing.T) {
	var client Client
	resp, err := get(baseURL+"/bots/get?page="+strconv.Itoa(34), nil)
	var respJSON getBotsResponse
	err = json.Unmarshal([]byte(resp), &respJSON)
	if err != nil {
		return
	}
	var botList Bots
	for page := 0; page < respJSON.TotalPage+1; page++ {
		bots := client.GetBots(page + 1)
		fmt.Println(page)
		botList = append(botList, *bots...)
	}
	for idx, bot := range botList {
		fmt.Printf("%d. %s (%d서버/하트 %d개)\n", idx+1, bot.Name, bot.Servers, bot.Votes)
	}
}

func TestClient_SearchBots(t *testing.T) {
	var client Client
	fmt.Println(client.SearchBots("초록", 1))
	fmt.Println(client.SearchBots("한마루", 1))
}

func TestClient_GetBotsByCategory(t *testing.T) {
	var client Client
	bots := client.GetBotsByCategory("뮤직", 1)
	for idx, bot := range *bots {
		fmt.Printf("%d. %s\n", idx+1, bot.Name)
	}
}

func TestClient_GetBot(t *testing.T) {
	var client Client
	fmt.Println(client.GetBot("770246143652397069").Name)
	fmt.Println(client.GetBot("687886541619462293").Name)
}
