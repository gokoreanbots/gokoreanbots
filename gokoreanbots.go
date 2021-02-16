package gokoreanbots

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/simsimler/gokoreanbots/gkbhttp"
)

var baseURL string  = "https://api.koreanbots.dev/v1"

// Client : 봇 클라이언트입니다.
type Client struct {
	token 		string
}

type voteData struct {
	Code int
	Voted bool
}

// PostServers : 서버 수를 업데이트합니다.
// 받는 인자들
// servers: int / 봇의 서버 수
func (c Client) PostServers(servers int) {
	headers := map[string]string{"Content-Type": "application/json", "token": c.token}
	seversJSON := []byte(`{"servers": ` + strconv.Itoa(servers) + `}`)
	err := gkbhttp.Post(baseURL + "/bots/servers", headers, seversJSON)
	if err != nil {
		fmt.Println("Error on post servers to koreanbots:", err)
	}
}

// IsVoted : 해당 유저의 봇 투표 여부를 불러옵니다.
// 받는 인자들
// userID: int / 유저의 ID
func (c Client) IsVoted(userID int) bool {
	resp := gkbhttp.Get(baseURL + "/bots/voted/" + strconv.Itoa(userID), map[string]string{"token": c.token})
	vD := voteData{}
	json.Unmarshal([]byte(resp), &vD)
	return vD.Voted
}

// NewClient : 새로운 클라이언트를 생성합니다.
// 받는 인자들
// session: discordgo.Session
// token: string / Koreanbots 토큰
// autoPost: bool / 서버 수 자동 포스트 여부
func NewClient(session *discordgo.Session, token string, autoPost bool) Client {
	client := Client{
		token: token,
	}
	if autoPost {
		go client.autoPostServers(session)
	}
	return client
}

func (c Client) autoPostServers(session *discordgo.Session) {
	for {
		c.PostServers(len(session.State.Guilds))
		time.Sleep(time.Minute * 30)
	}
}
