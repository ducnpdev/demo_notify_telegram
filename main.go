package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

var (
	ChatID   = "-4035316275"
	BotToken = "6970454443:AAHPBUSo2MFyELhL7mC16uyrB3N8xcYe6D8"
	Text     = "DucNPh notify telegram 123123"
)

// curl https://api.telegram.org/bot6970454443:AAHPBUSo2MFyELhL7mC16uyrB3N8xcYe6D8/getUpdates

func main() {
	fmt.Println("main")
	jsonStr := []byte(fmt.Sprintf(`{"chat_id": "%s", "text": "%s"}`, ChatID, Text))

	resp, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", BotToken), "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Default().Println("err:", err)
		return
	}
	defer resp.Body.Close()
	log.Default().Println("successfully")

}
