package types

// Bot type
type Bot struct {
	ID			string
	Name		string
	Servers		int
	Votes		int
	Intro		string
	Avatar		string
	URL			string
	Category 	Category
	Tag 		string
	State		BotState
	verified	bool
	trusted		bool
	boosted		bool
	vanity		string
	banner		string
	bg			string
}

// Category of bot
type Category struct {
	Mod 		bool
	Music 		bool
	GameRecord	bool
	WebDash 	bool
	Logging 	bool
	Gamble 		bool
	Game		bool
	Meme		bool
	Leveling	bool
	Utility		bool
	Translate	bool
	Talk		bool
	NSFW		bool
	Search		bool
}

// BotState state of bot
type BotState struct {
	OK 			bool
	Archived	bool
	Private		bool
}
