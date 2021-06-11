package gokoreanbots

import (
	"fmt"
	"testing"
)

func TestClient_GetBot(t *testing.T) {
	client := Client{}
	fmt.Println(client.GetBot("770246143652397069"))
}

func TestClient_GetBots(t *testing.T) {
	client := Client{}
	fmt.Println(client.GetBots(1))
}

func TestClient_GetNewBots(t *testing.T) {
	client := Client{}
	fmt.Println(client.GetNewBots())
}
