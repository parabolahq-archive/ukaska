package config

import (
	"encoding/json"
	"os"
)

var (
	BotToken             = os.Getenv("BOT_TOKEN")
	DbUrl                = os.Getenv("DB_URL")
	ChannelId            = os.Getenv("CHANNEL_ID")
	CollectionNames      = []string{}
	DisableNotification = true
	CollectionsCount     int
)

func Init() {
	json.Unmarshal([]byte(os.Getenv("COLLECTION_NAMES")), &CollectionNames)
	if os.Getenv("DISABLIE_NOTIFICATION") == "false" {
		DisableNotification = false
	}
}
