package gotify_messages

import (
	"github.com/gorilla/websocket"
	"gotify_matrix_bot/config"
	"log"
	"net/url"
)

var websocketURL, urlError = url.Parse("wss://" + config.Configuration.Gotify.URL + "/stream?token=" + config.Configuration.Gotify.ApiToken)

type callbackFunction func(string)

func OnNewMessage(callback callbackFunction) {

	if urlError != nil {
		log.Fatal("Error while trying to cast as url: ",
			"wss://"+config.Configuration.Gotify.URL+"/stream?token="+config.Configuration.Gotify.ApiToken, " ",
			urlError)
	}

	c, _, err := websocket.DefaultDialer.Dial(websocketURL.String(), nil)
	if err != nil {
		log.Fatal("Error while trying to connect to the webserver:", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Fatal("The websocket connection returned an error. Error message: ", err)
			}

			callback(string(message))

		}
	}()

}
