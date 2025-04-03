// main.go
package main

import (
	"log"
	"net/http"

	"mu-server-manager/config"
	"mu-server-manager/controllers"
)

func main() {
	err := config.LoadChannels("config/channels.json")
	if err != nil {
		log.Fatalf("Failed to load channels: %v", err)
	}

	http.HandleFunc("/start/", controllers.StartChannelHandler)
	http.HandleFunc("/stop/", controllers.StopChannelHandler)
	http.HandleFunc("/status", controllers.StatusHandler)
	http.HandleFunc("/start-all", controllers.StartAllHandler)
	http.HandleFunc("/stop-all", controllers.StopAllHandler)

	log.Println("Channel Manager running on :8080")
	http.ListenAndServe(":8080", nil)
}
