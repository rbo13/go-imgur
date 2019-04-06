package main

import (
	"fmt"
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

	fmt.Println("Uploaded")
	fmt.Println(res)
	fmt.Println(err)
}
