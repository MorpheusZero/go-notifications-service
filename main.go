package main

import (
	"fmt"
	"net/http"
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

	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
		var payload = &NotificationsPayload{NotificationType: "discord", Data: ""}
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
	fmt.Println(payload.NotificationType)
}
