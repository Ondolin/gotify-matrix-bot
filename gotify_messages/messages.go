package gotify_messages

import (
	"github.com/gotify/go-api-client/v2/auth"
	"github.com/gotify/go-api-client/v2/gotify"
	"github.com/gotify/go-api-client/v2/models"
	"gotify_matrix_bot/cache"
	"gotify_matrix_bot/config"
	"log"
	"net/http"
	"net/url"
)

/*

This code is deprecated. It was replaced with the websocket functionality, and is kept for reference reasons

*/

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

func GetNewMessage() *models.MessageExternal {

	messages := updateMessages()

	ca := cache.GetCache()

	newMessagesCount := len(messages) - ca.ReadMessages

	if config.Configuration.Debug {
		log.Println("Currently there are ", newMessagesCount, "undelivered messages.")
	}

	if newMessagesCount < 0 {
		ca.ReadMessages = len(messages)
		cache.SetCache(*ca)
		log.Println("Possibly some messages got deleted!")
		return nil
	} else if newMessagesCount > 0 {
		ca.ReadMessages++
		cache.SetCache(*ca)
		return messages[newMessagesCount-1]
	}

	return nil

}
