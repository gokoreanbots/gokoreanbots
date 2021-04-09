package gokoreanbots

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/bwmarrin/discordgo"
)

var (
	session, _ = discordgo.New("Bot ")
	kbclient = NewClient(session, "", true)
)

func TestRun(t *testing.T) {
	session.Open()
	fmt.Println(session.State.User.Username + "로 로그인했습니다.")
	session.AddHandler(messageCreate)
	sc := make(chan os.Signal, 2)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s.State.User.ID == m.Author.ID {
		return
	}
	switch m.Content {
	case "!votecheck":
		isVoted := kbclient.IsVoted(m.Author.ID)
		fmt.Println(isVoted)
		if isVoted {
			s.ChannelMessageSend(m.ChannelID, "하트를 누르셨군요")
		} else {
			s.ChannelMessageSend(m.ChannelID, "하트를 누르지 않으셨군요")
		}
	}
}
