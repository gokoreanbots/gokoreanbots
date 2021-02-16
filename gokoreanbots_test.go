package gokoreanbots

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/bwmarrin/discordgo"
)


func TestRun(t *testing.T) {
	session, _ := discordgo.New("Bot DiscordBotToken")
	client := NewClient(session, "KoreanbotsToken", true)
	session.Open()
	fmt.Println(session.State.User.Username + "로 로그인했습니다.")
	fmt.Println(client.IsVoted(710687896809766952))
	sc := make(chan os.Signal)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}