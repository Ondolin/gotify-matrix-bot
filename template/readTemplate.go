package template

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type Message struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func GetFormattedMessageString(message string) string {

	var m Message

	err := json.Unmarshal([]byte(message), &m)
	if err != nil {
		log.Println("[ERROR] Could not parse message from: " + message)
		return "Could not parse message from: " + message
	}

	templateString, err := ioutil.ReadFile("messageTamplate.md")

	if err != nil {
		log.Fatal("Could not find / read messageTamplate.md!")
	}

	content := strings.ReplaceAll(string(templateString), "[TITLE]", m.Title)

	content = strings.ReplaceAll(content, "[MESSAGE]", m.Message)

	return content
}
