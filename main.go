package main

import (
	"github.com/robfig/cron"
	"gotify_matrix_bot/gotify_messages"
	"gotify_matrix_bot/matrix"
	"log"
)

func main() {

	matrixClient := matrix.New()

	c := cron.New()

	c.AddFunc("0 * * * * *", func() {

		log.Println("Check")

		message := gotify_messages.GetNewMessage()

		if message != nil {
			log.Println(message.Message)
			matrixClient.Send(message.Message, message.Message)
		}
	})

	c.Start()

	for true {
	}

}
