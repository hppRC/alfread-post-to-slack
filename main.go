package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Token   string `json:"token"`
	Channel string `json:"channel"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postUrl := "https://slack.com/api/chat.postMessage"
	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	text := flag.String("m", "", "chatting text message")
	flag.Parse()

	if len(token) > 0 && len(*text) > 0 && len(channel) > 0 {
		values := url.Values{}
		values.Add("token", token)
		values.Add("channel", channel)
		values.Add("text", *text)

		client := http.Client{Timeout: time.Duration(10) * time.Second}
		req, err := http.NewRequest("POST", postUrl, strings.NewReader(values.Encode()))
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		defer res.Body.Close()

	} else {
		err := fmt.Sprintf("not set param. (t=%s, c=%s, m=%s)", token, channel, *text)
		log.Fatal(err)
		panic(err)
	}
}
