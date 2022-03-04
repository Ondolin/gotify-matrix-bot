package send

import (
	"gotify_matrix_bot/config"
	"gotify_matrix_bot/gotify_messages"
	"gotify_matrix_bot/matrix"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
)

func Unencrypted() {

	log.Println("Encryption inactive")

	cli, err := mautrix.NewClient(
		config.Configuration.Matrix.HomeServerURL,
		id.UserID("@"+config.Configuration.Matrix.Username+":"+config.Configuration.Matrix.MatrixDomain),
		config.Configuration.Matrix.Token)

	if err != nil {
		panic(err)
	}

	gotify_messages.OnNewMessage(func(message string) {

		err := matrix.SendUnencrypted(cli, id.RoomID(config.Configuration.Matrix.RoomID), message)
		if err != nil {
			log.Fatal("Could not send encrypted message to matrix. ", err)
		}

	})

	select {}
}
