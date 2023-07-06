package providers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// We set it globally, so we don't have to keep reading it.
var webhookURL string = os.Getenv("DISCORD_WEBHOOK_URL")

func ExecuteDiscordProvider(data *string) {
	if webhookURL == "" {
		fmt.Println("[DISCORD_WEBHOOK_URL] is not defined!")
		return
	}
	_, err := http.Post(webhookURL, "application/json", bytes.NewReader([]byte(*data)))
	if err != nil {
		fmt.Println(err.Error())
	}
}
