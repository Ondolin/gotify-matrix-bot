package send

import (
	"fmt"
	"github.com/robfig/cron"
	"gotify_matrix_bot/config"
	"gotify_matrix_bot/gotify_messages"
	"gotify_matrix_bot/matrix"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/id"
)

func Unencrypted() {

	log.Println("Encryption inactive")

	cli, err := mautrix.NewClient(config.Configuration.Matrix.HomeServerURL, "bot", "syt_Ym90_gBxYUIJyjVpThFvaAKmD_1Qa58F")
	if err != nil {
		panic(err)
	}

	// Log in to get access token and device ID.
	/*_, err = cli.Login(&mautrix.ReqLogin{
		Type: mautrix.AuthTypeToken,
		Identifier: mautrix.UserIdentifier{
			Type: mautrix.IdentifierTypeUser,
			User: config.Configuration.Matrix.Username,
		},
		Token: config.Configuration.Matrix.Token,
		StoreCredentials: true,
	})*/
	if err != nil {
		panic(err)
	}

	// Log out when the program ends (don't do this in real apps)
	defer func() {
		fmt.Println("Logging out")
		resp, err := cli.Logout()
		if err != nil {
			fmt.Println("Logout error:", err)
		}
		fmt.Println("Logout response:", resp)
	}()

	c := cron.New()

	c.AddFunc(config.Configuration.Gotify.PollTime, func() {

		if config.Configuration.Debug {
			log.Println("Check for new Messages")
		}

		message := gotify_messages.GetNewMessage()

		if message != nil {
			matrix.SendUnencrypted(cli, id.RoomID(config.Configuration.Matrix.RoomID), message.Message)
		}
	})

	c.Start()

	for true {
	}
}
