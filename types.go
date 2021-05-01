package gokoreanbots

// Bot type
type Bot struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Servers    int         `json:"servers"`
	Votes      int         `json:"votes"`
	Intro      string      `json:"intro"`
	Avatar     string      `json:"avatar"`
	URL        interface{} `json:"url"`
	Category   []string    `json:"category"`
	Tag        string      `json:"tag"`
	State      string      `json:"state"`
	Verified   int         `json:"verified"`
	Trusted    int         `json:"trusted"`
	Boosted    int         `json:"boosted"`
	Vanity     interface{} `json:"vanity"`
	Banner     interface{} `json:"banner"`
	BackGround interface{} `json:"bg"`
}

type Bots []Bot

type getBotsResponse struct {
	Code        int  `json:"code"`
	Data        Bots `json:"data"`
	CurrentPage int  `json:"currentPage"`
	TotalPage   int  `json:"totalPage"`
}
