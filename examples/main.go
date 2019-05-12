package main

import (
	"encoding/json"
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

	imgRes := imgur.Response{}
	json.NewDecoder(res.Body).Decode(&imgRes)

	print(imgRes.Data.Link)
}
