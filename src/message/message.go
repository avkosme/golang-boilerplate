package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var ChatId string = os.Getenv("CHAT_ID")
var ApiKey string = os.Getenv("API_KEY")

func SendMessage(text string) bool {
	reqBody, err := json.Marshal(map[string]string{
		"text":    text,
		"chat_id": ChatId,
	})
	Url := fmt.Sprintf("https://api.telegram.org/%v/%v", ApiKey, "sendMessage")

	res, err := http.Post(
		Url,
		"application/json",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	return true
}
