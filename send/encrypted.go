package send

import (
	"github.com/robfig/cron"
	"gotify_matrix_bot/config"
	"gotify_matrix_bot/gotify_messages"
	"gotify_matrix_bot/matrix"
	"gotify_matrix_bot/template"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/crypto"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/id"
	"strings"
)

func Encrypted() {

	log.Println("Encryption active")

	cli, err := mautrix.NewClient(
		config.Configuration.Matrix.HomeServerURL,
		id.UserID("@"+config.Configuration.Matrix.Username+":"+strings.ReplaceAll(config.Configuration.Matrix.HomeServerURL, "https://", "")),
		config.Configuration.Matrix.Token)

	if err != nil {
		panic(err)
	}

	// Create a store for the e2ee keys. In real apps, use NewSQLCryptoStore instead of NewGobStore.
	cryptoStore, err := crypto.NewGobStore("cryptoStore.gob")
	if err != nil {
		panic(err)
	}

	mach := crypto.NewOlmMachine(cli, &matrix.Logger{}, cryptoStore, &matrix.FakeStateStore{})
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

	c := cron.New()

	c.AddFunc(config.Configuration.Gotify.PollTime, func() {

		if config.Configuration.Debug {
			log.Println("Check for new Messages")
		}

		message := gotify_messages.GetNewMessage()

		if message != nil {
			matrix.SendEncrypted(mach, cli, id.RoomID(config.Configuration.Matrix.RoomID), template.GetFormattedMessageString(message))
		}
	})

	c.Start()

	for true {
	}
}
