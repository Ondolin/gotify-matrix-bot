package gotify_messages

import (
	"gotify_matrix_bot/config"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type callbackFunction func(string)

func OnNewMessage(callback callbackFunction) {

	websocketURL, urlError := url.Parse(config.Configuration.Gotify.URL + "/stream?token=" + config.Configuration.Gotify.ApiToken)

	if urlError != nil {
		log.Fatal("Error while trying to parse gotify url: ",
			config.Configuration.Gotify.URL+"/stream?token="+config.Configuration.Gotify.ApiToken, " ",
			urlError)
	}

	c, _, err := websocket.DefaultDialer.Dial(websocketURL.String(), nil)
	if err != nil {
		log.Fatal("Error while trying to connect to the gotify server:", err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Fatal("The websocket connection to gotify returned an error. Error message: ", err)
			}

			callback(string(message))

		}
	}()

}
