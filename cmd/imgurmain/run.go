package imgurmain

import (
	"os"

	imgur "github.com/rbo13/go-imgur"
)

var imgurClientID = os.Getenv("IMGUR_CLIENT_ID")

// Run is the main function for imgurmain package.
func Run() {
	fileName := os.Args[1]

	imgur := imgur.New(imgurClientID)
	res, err := imgur.Upload(fileName)
	if err != nil {
		panic(err)
	}
	print(res)
}
