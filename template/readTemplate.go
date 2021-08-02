package template

import (
	"github.com/gotify/go-api-client/v2/models"
	"io/ioutil"
	"log"
	"strings"
)

func GetFormattedMessageString(message *models.MessageExternal) string {
	templateString, err := ioutil.ReadFile("messageTamplate.md")

	if err != nil {
		log.Fatal(err)
	}

	content := strings.ReplaceAll(string(templateString), "[TITLE]", message.Title)

	content = strings.ReplaceAll(content, "[MESSAGE]", message.Message)

	return content
}
