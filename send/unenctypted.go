package send

import (
	"github.com/robfig/cron"
	"gotify_matrix_bot/config"
	"gotify_matrix_bot/gotify_messages"
	"gotify_matrix_bot/matrix"
	"gotify_matrix_bot/template"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
	"strings"
)

func Unencrypted() {

	log.Println("Encryption inactive")

	cli, err := mautrix.NewClient(
		config.Configuration.Matrix.HomeServerURL,
		id.UserID("@"+config.Configuration.Matrix.Username+":"+strings.ReplaceAll(config.Configuration.Matrix.HomeServerURL, "https://", "")),
		config.Configuration.Matrix.Token)

	if err != nil {
		panic(err)
	}

	c := cron.New()

	c.AddFunc(config.Configuration.Gotify.PollTime, func() {

		if config.Configuration.Debug {
			log.Println("Check for new Messages")
		}

		message := gotify_messages.GetNewMessage()

		if message != nil {

			c.Stop()

			sendErr := matrix.SendUnencrypted(cli, id.RoomID(config.Configuration.Matrix.RoomID), template.GetFormattedMessageString(message))

			// resend message until it is successful
			for sendErr != nil {
				log.Println("Try to resend message...")
				sendErr = matrix.SendUnencrypted(cli, id.RoomID(config.Configuration.Matrix.RoomID), template.GetFormattedMessageString(message))
			}

			c.Start()
		}
	})

	c.Start()

	for true {
	}
}
