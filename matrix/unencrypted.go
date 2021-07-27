//
// Code inspired by rakshazi from https://gitlab.com/rakshazi/desktop2gotify/-/blob/master/matrix/matrix.go
//

package matrix

import (
	"github.com/matrix-org/gomatrix"
	"gotify_matrix_bot/config"
	"log"
)

// Matrix bot config
type Matrix struct {
	HomeServer string
	User       string
	Token      string
	Room       string
	Client     *gomatrix.Client
}

// New matrix client
func New() *Matrix {

	hs := config.Configuration.Matrix.HomeServerURL
	user := config.Configuration.Matrix.Username
	token := config.Configuration.Matrix.Token
	room := config.Configuration.Matrix.RoomID

	log.Println(hs, user, token, room)

	client, err := gomatrix.NewClient(hs, user, token)
	if err != nil {
		log.Println("Could not connect with homeserver", err)
		panic(err)
	}

	/*_, err = client.JoinRoom(room, "", nil)
	if err != nil {
		fmt.Println("[matrix] cannot join room, did you invite that user to room?")
		panic(err)
	}*/

	return &Matrix{
		HomeServer: hs,
		User:       user,
		Token:      token,
		Room:       room,
		Client:     client,
	}
}

func (mx *Matrix) Send(message string, formattedMessage string) {

	if config.Configuration.Debug {
		log.Println("Send new unencrypted message to matrix server", mx, ". Message content: ", message)
	}

	_, err := mx.Client.SendFormattedText(mx.Room, message, formattedMessage)

	if err != nil {
		log.Println("Could not send the message to the matrix server!", err)
	}

}
