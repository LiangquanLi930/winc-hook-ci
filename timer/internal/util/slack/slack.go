package slack

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"timer/internal/util/yaml"
)

func createSlackMessage(message string) io.Reader {
	slackMessage := map[string]string{
		"text": message,
	}
	stringMessage, err := json.Marshal(slackMessage)
	if err != nil {
		panic(err)
	}
	return strings.NewReader(string(stringMessage))
}

func SendSlack(message string) error {
	client := &http.Client{}
	_, err := client.Post(yaml.GetConfig().SlackMessageUrl, "application/json", createSlackMessage(message))
	if err != nil {
		log.Println("slack send error")
		return err
	}
	return nil
}
