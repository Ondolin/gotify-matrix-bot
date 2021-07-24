package gotify_messages

import (
	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
	"gotify_matrix_bot/config"
	"log"
	"net/http"
	"net/url"
)

var myURL, _ = url.Parse(config.Configuration.Gotify.URL)
var client = gotify.NewClient(myURL, &http.Client{})
var authToken = auth.TokenAuth(config.Configuration.Gotify.ApiToken)

func updateMessages() []*models.MessageExternal {
	res, err := client.Message.GetMessages(nil, authToken)

	if err != nil {
		log.Fatal("Could not get Messages for gotify server with url: ", myURL, err)
	}

	return res.Payload.Messages

}

var messageCount = 0

func GetNewMessage() *models.MessageExternal {

	messages := updateMessages()

	newMessagesCount := len(messages) - messageCount

	if newMessagesCount < 0 {
		messageCount = len(messages)
		log.Println("Possibly some messages got deleted!")
		return nil
	} else if newMessagesCount > 0 {
		messageCount++
		return messages[newMessagesCount-1]
	}

	return nil

}
