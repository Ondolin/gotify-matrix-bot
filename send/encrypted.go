package send

import (
	"fmt"
	"github.com/robfig/cron"
	"gotify_matrix_bot/config"
	"gotify_matrix_bot/gotify_messages"
	"gotify_matrix_bot/matrix"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/crypto"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
)

func Encrypted() {

	log.Println("Encryption active")

	cli, err := mautrix.NewClient(config.Configuration.Matrix.HomeServerURL, "", "")
	if err != nil {
		panic(err)
	}

	// Log in to get access token and device ID.
	_, err = cli.Login(&mautrix.ReqLogin{
		Type: mautrix.AuthTypePassword,
		Identifier: mautrix.UserIdentifier{
			Type: mautrix.IdentifierTypeUser,
			User: config.Configuration.Matrix.Username,
		},
		Password:         config.Configuration.Matrix.Password,
		StoreCredentials: true,
	})
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

	// Create a store for the e2ee keys. In real apps, use NewSQLCryptoStore instead of NewGobStore.
	cryptoStore, err := crypto.NewGobStore("test.gob")
	if err != nil {
		panic(err)
	}

	mach := crypto.NewOlmMachine(cli, &matrix.FakeLogger{}, cryptoStore, &matrix.FakeStateStore{})
	// Load data from the crypto store
	err = mach.Load()
	if err != nil {
		panic(err)
	}

	// Hook up the OlmMachine into the Matrix client so it receives e2ee keys and other such things.
	syncer := cli.Syncer.(*mautrix.DefaultSyncer)
	syncer.OnSync(func(resp *mautrix.RespSync, since string) bool {
		mach.ProcessSyncResponse(resp, since)
		return true
	})
	syncer.OnEventType(event.StateMember, func(source mautrix.EventSource, evt *event.Event) {
		mach.HandleMemberEvent(evt)
	})
	// Start long polling in the background
	go func() {
		err = cli.Sync()
		if err != nil {
			panic(err)
		}
	}()
	// Put an internal room ID here, then type text into the program's input to send encrypted messages.
	// To stop the program, press enter without typing anything
	/*var sendToRoom id.RoomID = "!internalRoomID:example.com"
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		go matrix.SendEncrypted(mach, cli, sendToRoom, line)
	}*/

	c := cron.New()

	c.AddFunc(config.Configuration.Gotify.PollTime, func() {

		if config.Configuration.Debug {
			log.Println("Check for new Messages")
		}

		message := gotify_messages.GetNewMessage()

		if message != nil {
			matrix.SendEncrypted(mach, cli, id.RoomID(config.Configuration.Matrix.RoomID), message.Message)
		}
	})

	c.Start()

	for true {
	}
}
