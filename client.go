package gokoreanbots

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

const (
	baseURL string = "https://koreanbots.dev/api/v2"
)

// PostServers : 서버 수를 업데이트합니다.
func (c Client) PostServers() error {
	headers := strMap{"Content-Type": "application/json", "Authorization": c.token}
	serverJSON, _ := json.Marshal(postServer{
		Servers: len(c.session.State.Guilds),
	})
	err := post(
		urlFmt(
			baseURL,
			"/bots/"+c.session.State.User.ID+"/stats",
			nil,
		),
		&headers,
		serverJSON,
	)
	return err
}

// IsVoted : 해당 유저의 봇 투표 여부를 불러옵니다.
// 인자
// userID: int / 확인할 유저 ID
func (c Client) IsVoted(userID string) bool {
	resp, err := get(
		urlFmt(
			baseURL,
			"/bots/"+c.session.State.User.ID+"/vote",
			&strMap{
				"userID": userID,
			},
		),
		strMap{
			"Authorization": c.token,
		},
	)
	if err != nil {
		log.Println("[GoKOREANBOTS] failed to get vote because of " + err.Error())
		return false
	}
	vD := voteResp{}
	err = json.Unmarshal([]byte(resp), &vD)
	if err != nil {
		return false
	}
	return vD.Data.Voted
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
// 인자
// page: int / 조회할 페이지 번호
func (c Client) GetBots(page int) *Bots {
	resp, err := get(
		urlFmt(
			baseURL,
			"/list/bots/votes",
			&strMap{
				"page": strconv.Itoa(page),
			},
		),
		nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	var parsedResp getBotsResp
	err = json.Unmarshal([]byte(resp), &parsedResp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &parsedResp.Data.Data
}

// GetBot : KOREANBOTS에 등록된 봇 정보를 가져옵니다.
// 인자
// id: string / 봇의 ID
func (c Client) GetBot(id string) *Bot {
	resp, err := get(
		urlFmt(
			baseURL,
			"/bots/"+id,
			nil),
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	var parsedResp getBotResp
	err = json.Unmarshal([]byte(resp), &parsedResp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &parsedResp.Data
}

// SearchBots : KOREANBOTS에서 해당 쿼리로 봇을 검색합니다.
// 인자
// query: string / 검색어
// page: int / 조회할 페이지 번호
func (c Client) SearchBots(query string, page int) *Bots {
	resp, err := get(
		urlFmt(
			baseURL,
			"/search/bots",
			&strMap{
				"query": query,
				"page":  strconv.Itoa(page),
			},
		),
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	var parsedResp getBotsResp
	err = json.Unmarshal([]byte(resp), &parsedResp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &parsedResp.Data.Data
}

// GetNewBots : 새로 등록된 봇들을 가져옵니다
func (c Client) GetNewBots() *Bots {
	resp, err := get(
		urlFmt(baseURL,
			"/list/bots/new",
			nil,
		),
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil
	}
	var parsedResp getBotsResp
	err = json.Unmarshal([]byte(resp), &parsedResp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &parsedResp.Data.Data
}
