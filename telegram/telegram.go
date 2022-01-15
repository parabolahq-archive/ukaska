package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"ukaska/backup"
	"ukaska/config"
	"ukaska/telegram/endpoint"
	"ukaska/telegram/structs"
)

func SendMessage(messageText string) structs.MessageResponse {
	uri := endpoint.SendMessage(messageText, config.ChannelId, config.DisableNotification)
	resp, err := http.Get(uri)
	backup.HandleIfError(err)
	body, _ := ioutil.ReadAll(resp.Body)
	return parseMessageResponse(body)
}

func SendDocument(documentName string) structs.MessageResponse {
	client := &http.Client{}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, _ := writer.CreateFormFile("document", documentName)
	file, _ := os.Open(documentName)
	_, _ = io.Copy(fw, file)
	writer.WriteField("disableNotifications", strconv.FormatBool(config.DisableNotification))
	_ = writer.Close()
	req, err := http.NewRequest("POST", endpoint.SendDocument(config.ChannelId), bytes.NewReader(body.Bytes()))
	backup.HandleIfError(err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	backup.HandleIfError(err)
	responseBody, _ := io.ReadAll(resp.Body)
	return parseMessageResponse(responseBody)
}

func parseMessageResponse(responseBody []byte) (message structs.MessageResponse) {
	err := json.Unmarshal([]byte(responseBody), &message)
	backup.HandleIfError(err)
	return
}
