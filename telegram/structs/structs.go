package structs

type Message struct {
	Id       int      `json:"message_id"`
	From     User     `json:"from"`
	Chat     Chat     `json:"chat"`
	Date     int      `json:"date"`
	Text     string   `json:"text"`
	Document Document `json:"document"`
}

type Document struct {
	FileName     string `json:"file_name"`
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
}

type User struct {
	Id        int64  `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Title     string `json:"title"`
}

type Chat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type MessageResponse struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}
