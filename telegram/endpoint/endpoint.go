package endpoint

import (
	"fmt"
	"net/url"
	"strconv"
	"ukaska/config"
)

var baseUrl = fmt.Sprintf("https://api.telegram.org/bot%s/", config.BotToken)

func SendMessage(text, chatId string, disableNotification bool) string {
	return baseUrl + "sendMessage?" + url.PathEscape(fmt.Sprintf("text=%s&chat_id=%s&disable_notification=%s", text, chatId, strconv.FormatBool(disableNotification)))
}

func SendDocument(chatId string) string {
	return baseUrl + "sendDocument?" + url.PathEscape(fmt.Sprintf("&chat_id=%s", chatId))
}
