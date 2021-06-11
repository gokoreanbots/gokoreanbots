package gokoreanbots

import (
	"github.com/bwmarrin/discordgo"
)

// NewClient : 새로운 클라이언트를 생성합니다.
// 인자
// session: discordgo.Session
// token: string / KOREANBOTS 토큰
// autoPost: bool / 서버 수 자동 포스트 여부
func NewClient(session *discordgo.Session, token string, autoPost bool) Client {
	client := Client{
		token:    token,
		autoPost: autoPost,
		session:  session,
	}
	if autoPost {
		go client.autoPostServers()
	}
	return client
}
