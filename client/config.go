package client

type Config struct {
	Schema          string
	TelegramHost    string
	TelegramPath    string
	BotKey          string
	SendMessagePath string
	GetUpdatePath   string
	Timeout         int
}

var config = &Config{
	Schema:          "https",
	TelegramHost:    "api.telegram.org",
	TelegramPath:    "bot",
	BotKey:          "XXXXXXXXXXXXXXXXXXXXX",
	SendMessagePath: "sendMessage",
	GetUpdatePath:   "getUpdates",
	Timeout:         5,
}
