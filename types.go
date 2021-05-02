package gokoreanbots

// Bot type
type Bot struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Servers  int         `json:"servers"`
	Votes    int         `json:"votes"`
	Intro    string      `json:"intro"`
	Avatar   interface{} `json:"avatar"`
	URL      interface{} `json:"url"`
	Category []string    `json:"category"`
	Tag      string      `json:"tag"`
	State    string      `json:"state"`
	Verified int         `json:"verified"`
	Trusted  int         `json:"trusted"`
	Boosted  int         `json:"boosted"`
	Vanity   interface{} `json:"vanity"`
	Banner   interface{} `json:"banner"`
	Bg       interface{} `json:"bg"`
}

type BotDetail struct {
	ID     string `json:"id"`
	Date   int    `json:"date"`
	Owners []struct {
		Id       string      `json:"id"`
		Username string      `json:"username"`
		Tag      string      `json:"tag"`
		Avatar   interface{} `json:"avatar"`
		Bots     []string    `json:"bots"`
	} `json:"owners"`
	Lib      string      `json:"lib"`
	Prefix   string      `json:"prefix"`
	Votes    int         `json:"votes"`
	Servers  int         `json:"servers"`
	Intro    string      `json:"intro"`
	Desc     string      `json:"desc"`
	Web      interface{} `json:"web"`
	Git      interface{} `json:"git"`
	Url      string      `json:"url"`
	Category []string    `json:"category"`
	Status   string      `json:"status"`
	Name     string      `json:"name"`
	Avatar   interface{} `json:"avatar"`
	Tag      string      `json:"tag"`
	Verified int         `json:"verified"`
	Trusted  int         `json:"trusted"`
	Discord  string      `json:"discord"`
	Boosted  int         `json:"boosted"`
	State    string      `json:"state"`
	Vanity   interface{} `json:"vanity"`
	Bg       interface{} `json:"bg"`
	Banner   interface{} `json:"banner"`
}

type Bots []Bot

type getBotsResponse struct {
	Code        int  `json:"code"`
	Data        Bots `json:"data"`
	CurrentPage int  `json:"currentPage"`
	TotalPage   int  `json:"totalPage"`
}

type getBotResponse struct {
	Code int       `json:"code"`
	Data BotDetail `json:"data"`
}
