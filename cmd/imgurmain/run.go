package imgurmain

import (
	"log"
	"os"

	imgur "github.com/rbo13/go-imgur"
)

var imgurClientID = os.Getenv("IMGUR_CLIENT_ID")

// Run is the main function for imgurmain package.
func Run() {

	if len(os.Args) < 1 {
		log.Fatalf("Must accept 1 valid image path")
		return
	}

	fileName := os.Args[1]

	imgr := imgur.New(imgurClientID)
	res, err := imgr.Upload(fileName)
	if err != nil {
		panic(err)
	}

	print(res.GetImageLink())
}
