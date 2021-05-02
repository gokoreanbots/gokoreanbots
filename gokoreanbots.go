package gokoreanbots

import (
	"encoding/json"
	"log"
	"net/url"
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
	return err
}

// IsVoted : 해당 유저의 봇 투표 여부를 불러옵니다.
// 받는 인자들
// userID: int / 확인할 유저 ID
func (c Client) IsVoted(userID string) bool {
	resp, err := get(baseURL+"/bots/voted/"+userID, map[string]string{"token": c.token})
	if err != nil {
		log.Println("[GoKOREANBOTS] failed to get vote because of " + err.Error())
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
			log.Println("[GoKOREANBOTS] failed to post servers because of " + err.Error())
		}
		time.Sleep(time.Minute * 30)
	}
}

// GetBots : 봇 순위를 가져옵니다.
// 받는 인자들
// page: int / 조회할 페이지 번호
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

// GetBot : KOREANBOTS에 등록된 봇 정보를 가져옵니다.
// 받는 인자들
// id: string / 봇의 ID
func (c Client) GetBot(id string) *BotDetail {
	resp, err := get(baseURL+"/bots/get/"+id, nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	var parsedResp getBotResponse
	err = json.Unmarshal([]byte(resp), &parsedResp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &parsedResp.Data
}

// SearchBots : KOREANBOTS에서 해당 쿼리로 봇을 검색합니다.
// 받는 인자들
// query: string / 검색어
// page: int / 조회할 페이지 번호
func (c Client) SearchBots(query string, page int) *Bots {
	resp, err := get(baseURL+"/bots/search"+"?q="+url.QueryEscape(query)+"&page="+strconv.Itoa(page), nil)
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

func (c Client) GetBotsByCategory(category string, page int) *Bots {
	resp, err := get(baseURL+"/bots/category/"+category+"?page="+strconv.Itoa(page), nil)
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
