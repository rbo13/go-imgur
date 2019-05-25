package main

import (
	"log"
	"os"

	imgur "github.com/rbo13/go-imgur"
)

var (
	imgurSecretKey = os.Getenv("IMGUR_CLIENT_SECRET")
	imgurClientID  = os.Getenv("IMGUR_CLIENT_ID")
)

func main() {
	imgr := imgur.New(imgurClientID)
	res, err := imgr.Upload("./examples/4k.jpg")

	if err != nil {
		log.Println(err)
		return
	}

	println(res.GetImageLink())
	println(res.GetImageID())
	println("Image Link successfully copied to clipboard")
	res.Clipboard()
}
