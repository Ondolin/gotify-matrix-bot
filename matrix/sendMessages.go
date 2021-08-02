// Code by tulir from https://mau.dev/-/snippets/6

package matrix

import (
	"fmt"
	"gotify_matrix_bot/config"
	"log"
	"maunium.net/go/mautrix"
	"maunium.net/go/mautrix/crypto"
	"maunium.net/go/mautrix/event"
	"maunium.net/go/mautrix/format"
	"maunium.net/go/mautrix/id"
)

func SendEncrypted(mach *crypto.OlmMachine, cli *mautrix.Client, roomID id.RoomID, text string) {

	if config.Configuration.Debug {
		log.Println("Sending new unencrypted message")
	}

	content := format.RenderMarkdown(text, true, true)
	encrypted, err := mach.EncryptMegolmEvent(roomID, event.EventMessage, content)
	// These three errors mean we have to make a new Megolm session
	if err == crypto.SessionExpired || err == crypto.SessionNotShared || err == crypto.NoGroupSession {
		err = mach.ShareGroupSession(roomID, getUserIDs(cli, roomID))
		if err != nil {
			panic(err)
		}
		encrypted, err = mach.EncryptMegolmEvent(roomID, event.EventMessage, content)
	}
	if err != nil {
		panic(err)
	}
	resp, err := cli.SendMessageEvent(roomID, event.EventEncrypted, encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Send response:", resp)
}

func SendUnencrypted(cli *mautrix.Client, roomID id.RoomID, text string) {

	if config.Configuration.Debug {
		log.Println("Sending new unencrypted message")
	}

	_, err := cli.SendMessageEvent(roomID, event.EventMessage, format.RenderMarkdown(text, true, true))

	if err != nil {
		panic(err)
	}

}
