package template

import (
	"github.com/gotify/go-api-client/v2/models"
	"io/ioutil"
	"strings"
)

func GetFormattedMessageString(message *models.MessageExternal) string {
	templateString, err := ioutil.ReadFile("messageTamplate.md")

	if err != nil {
		return message.Message
	}

	content := strings.ReplaceAll(string(templateString), "[TITLE]", message.Title)

	content = strings.ReplaceAll(content, "[MESSAGE]", message.Message)

	return content
}
