package main

import (
	"gotify_matrix_bot/config"
	"gotify_matrix_bot/send"
	"log"
)

func main() {

	log.Println("The gotify matrix bot has started now.")

	if config.Configuration.Matrix.Encrypted {
		send.Encrypted()
	} else {
		send.Unencrypted()
	}
}
