package gokoreanbots

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

var baseURL string  = "https://api.koreanbots.dev/v1"

// Client : 봇 클라이언트입니다.
type Client struct {
	token 		string
	autoPost	bool
	session		*discordgo.Session
}

type voteData struct {
	Code	int
	Voted 	bool
}

// Start : 서버 수 자동 업데이트를 시작합니다.
// discordgo.Session이 오픈된 후 서버 수 자동 업데이트를 시작해주세요.
// autoPost를 false로 하실 경우 서버 수가 자동으로 업데이트 되지 않습니다.
func (c Client) Start() {
	if c.autoPost {
		go c.autoPostServers(c.session)
	}
}

// PostServers : 서버 수를 업데이트합니다.
// 받는 인자들
// servers: int / 봇의 서버 수
func (c Client) PostServers(servers int) {
	headers := map[string]string{"Content-Type": "application/json", "token": c.token}
	seversJSON := []byte(`{"servers": ` + strconv.Itoa(servers) + `}`)
	err := post(baseURL + "/bots/servers", headers, seversJSON)
	if err != nil {
		fmt.Println(err)
	}
}

// IsVoted : 해당 유저의 봇 투표 여부를 불러옵니다.
// 받는 인자들
// userID: int / 유저의 ID
func (c Client) IsVoted(userID string) (bool, error) {
	resp, err := get(baseURL + "/bots/voted/" + userID, map[string]string{"token": c.token})
	vD := voteData{}
	json.Unmarshal([]byte(resp), &vD)
	return vD.Voted, err
}

func (c Client) autoPostServers(session *discordgo.Session) {
	for {
		c.PostServers(len(session.State.Guilds))
		time.Sleep(time.Minute * 30)
	}
}

// NewClient : 새로운 클라이언트를 생성합니다.
// 받는 인자들
// session: discordgo.Session
// token: string / Koreanbots 토큰
// autoPost: bool / 서버 수 자동 포스트 여부
func NewClient(session *discordgo.Session, token string, autoPost bool) Client {
	client := Client{
		token: token,
		autoPost: autoPost,
	}
	return client
}
