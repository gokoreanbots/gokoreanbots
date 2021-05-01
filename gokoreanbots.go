package gokoreanbots

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	baseURL string = "https://api.koreanbots.dev/v1"
	// baseURLv2 string = "https://api.koreanbots.dev/v2" // For Future
)

// Client : 봇 클라이언트입니다.
type Client struct {
	token    string
	autoPost bool
	session  *discordgo.Session
}

type voteData struct {
	Code  int  `json:"code"`
	Voted bool `json:"voted"`
}

type serverSendTemplate struct {
	Servers int `json:"servers"`
}

// PostServers : 서버 수를 업데이트합니다.
func (c Client) PostServers() error {
	headers := map[string]string{"Content-Type": "application/json", "token": c.token}
	serverJSON, _ := json.Marshal(serverSendTemplate{
		Servers: len(c.session.State.Guilds),
	})
	err := post(baseURL+"/bots/servers", &headers, serverJSON)
	if err != nil {
	}
	return err
}

// IsVoted : 해당 유저의 봇 투표 여부를 불러옵니다.
// 받는 인자들
// userID: int / 유저의 ID
func (c Client) IsVoted(userID string) bool {
	resp, err := get(baseURL+"/bots/voted/"+userID, map[string]string{"token": c.token})
	if err != nil {
		log.Println("[GoKoreanbots] failed to get vote because of " + err.Error())
		return false
	}
	vD := voteData{}
	err = json.Unmarshal([]byte(resp), &vD)
	if err != nil {
		return false
	}
	return vD.Voted
}

func (c Client) autoPostServers() {
	for {
		err := c.PostServers()
		if err != nil {
			log.Println("[GoKoreanbots] failed to post servers because of " + err.Error())
		}
		time.Sleep(time.Minute * 30)
	}
}

func (c Client) GetBots(page int) *Bots {
	resp, err := get(baseURL+"/bots/get?page="+strconv.Itoa(page), nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	var parsedResp getBotsResponse
	err = json.Unmarshal([]byte(resp), &parsedResp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &parsedResp.Data
}

// NewClient : 새로운 클라이언트를 생성합니다.
// 받는 인자들
// session: discordgo.Session
// token: string / Koreanbots 토큰
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
