package main

import (
	"github.com/robfig/cron"
	"gotify_matrix_bot/config"
	"gotify_matrix_bot/gotify_messages"
	"gotify_matrix_bot/matrix"
	"log"
)

func main() {

	matrixClient := matrix.New()

	c := cron.New()

	c.AddFunc("*/10 * * * * *", func() {

		if config.Configuration.Debug {
			log.Println("Check for new Messages")
		}

		message := gotify_messages.GetNewMessage()

		if message != nil {
			matrixClient.Send(message.Message, message.Message)
		}
	})

	c.Start()

	for true {
	}

}
