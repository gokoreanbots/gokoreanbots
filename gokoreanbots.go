package gokoreanbots

import (
	"encoding/json"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	baseURL string  = "https://api.koreanbots.dev/v1"
	baseURLv2 string = "https://api.koreanbots.dev/v2" // For Future
)


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

type serverSendTemplate struct {
	Servers int `json:"servers"`
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

func (c Client) PostServers() {
	headers := map[string]string{"Content-Type": "application/json", "token": c.token}
	strJSON, _ := json.Marshal(serverSendTemplate{
		Servers: len(c.session.State.Guilds),
	})
	seversJSON := strJSON
	err := post(baseURL + "/bots/servers", headers, seversJSON)
	if err != nil {
		log.Println(err)
	}
}

// IsVoted : 해당 유저의 봇 투표 여부를 불러옵니다.
// 받는 인자들
// userID: int / 유저의 ID
func (c Client) IsVoted(userID string) bool {
	resp, err := get(baseURL + "/bots/voted/" + userID, map[string]string{"token": c.token})
	if err != nil {
		log.Println(err)
		return false
	}
	vD := voteData{}
	json.Unmarshal([]byte(resp), &vD)
	return vD.Voted
}

func (c Client) autoPostServers(session *discordgo.Session) {
	for {
		c.PostServers()
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
