package endpoint

import (
	"fmt"
	"net/url"
	"ukaska/config"
)

var baseUrl = fmt.Sprintf("https://api.telegram.org/bot%s/", config.BotToken)

func SendMessage(text, chatId string) string {
	return baseUrl + "sendMessage?" + url.PathEscape(fmt.Sprintf("text=%s&chat_id=%s", text, chatId))
}

func SendDocument(chatId string) string {
	return baseUrl + "sendDocument?" + url.PathEscape(fmt.Sprintf("&chat_id=%s", chatId))
}
