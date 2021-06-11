package gokoreanbots

import "github.com/bwmarrin/discordgo"

// Client : 봇 클라이언트입니다.
type Client struct {
	token    string
	autoPost bool
	session  *discordgo.Session
}

type Bot struct {
	Id       string      `json:"id"`
	Flags    int         `json:"flags"`
	Owners   []Owner     `json:"owners"`
	Lib      string      `json:"lib"`
	Prefix   string      `json:"prefix"`
	Votes    int         `json:"votes"`
	Servers  int         `json:"servers"`
	Intro    string      `json:"intro"`
	Desc     string      `json:"desc"`
	Web      string      `json:"web"`
	Git      string      `json:"git"`
	Url      string      `json:"url"`
	Category []string    `json:"category"`
	Status   string      `json:"status"`
	Discord  string      `json:"discord"`
	State    string      `json:"state"`
	Vanity   string      `json:"vanity"`
	Bg       string      `json:"bg"`
	Banner   string      `json:"banner"`
	Tag      string      `json:"tag"`
	Avatar   interface{} `json:"avatar"`
	Name     string      `json:"name"`
}

type User struct {
	Id       string      `json:"id"`
	Flags    int         `json:"flags"`
	Github   interface{} `json:"github"`
	Tag      string      `json:"tag"`
	Username string      `json:"username"`
	Bots     *Bots       `json:"bots"`
}

type Owner struct {
	User
	Bots []string `json:"bots"`
}

type Bots []Bot

type strMap map[string]string

type voteResp struct {
	Code int `json:"code"`
	Data struct {
		Voted    bool  `json:"voted"`
		LastVote int64 `json:"lastVote"`
	} `json:"data"`
	Version int `json:"version"`
}

type postServer struct {
	Servers int `json:"servers"`
}

type getBotResp struct {
	Code    int `json:"code"`
	Data    Bot `json:"data"`
	Version int `json:"version"`
}

type getBotsResp struct {
	Code int `json:"code"`
	Data struct {
		Type        string `json:"type"`
		Data        Bots   `json:"data"`
		CurrentPage int    `json:"currentPage"`
		TotalPage   int    `json:"totalPage"`
	} `json:"data"`
	Version int `json:"version"`
}
