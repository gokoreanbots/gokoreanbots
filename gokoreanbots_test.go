package gokoreanbots

import (
	"fmt"
	"testing"
)

func TestGetBots(t *testing.T) {
	var kbclient Client
	fmt.Println(kbclient.GetBots(5))
}

func TestGetBot(t *testing.T) {
	var kbclient Client
	fmt.Println(kbclient.GetBot("770246143652397069").Name)
	fmt.Println(kbclient.GetBot("687886541619462293").Name)
}
