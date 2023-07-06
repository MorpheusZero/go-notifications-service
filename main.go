package main

import (
	"encoding/json"
	"net/http"
	"notifications-service/providers"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", GetHealthCheck)
	mux.HandleFunc("/v1/notify", PostNotify)
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}

func GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func PostNotify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	apiKey := r.Header.Get("x-api-key")
	if !IsAPIKeyValid(&apiKey) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
		return
	}

	if r.Method == http.MethodPost {
		payload := &NotificationsPayload{}
		err := json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("bad_request"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		go NotificationsHandler(payload)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed!"))
	}
}

type NotificationsPayload struct {
	NotificationType string `json:"notification_type"`
	Data             string `json:"data"`
}

func NotificationsHandler(payload *NotificationsPayload) {
	switch payload.NotificationType {
	case DISCORD:
		providers.ExecuteDiscordProvider(&payload.Data)
	default:
		return
	}
}

func IsAPIKeyValid(givenKey *string) bool {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("[API_KEY] is not defined!")
	}
	return apiKey == *givenKey
}

const (
	DISCORD = "discord"
)
