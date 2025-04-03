// controllers/channel_controller.go
package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"mu-server-manager/config"
	"mu-server-manager/process"
)

func StartChannelHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/start/")
	channel := config.GetChannelByID(id)
	if channel == nil {
		http.Error(w, "Channel not found", http.StatusNotFound)
		return
	}

	err := process.StartProcess(channel.ID, channel.RunCmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Write([]byte("Started " + channel.ID))
}

func StopChannelHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/stop/")
	err := process.StopProcess(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write([]byte("Stopped " + id))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	channels := config.GetChannels()
	type status struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Running bool   `json:"running"`
		Port    int    `json:"port"`
	}
	var result []status
	for _, ch := range channels {
		result = append(result, status{
			ID:      ch.ID,
			Name:    ch.Name,
			Running: process.IsRunning(ch.ID),
			Port:    ch.Port,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func StartAllHandler(w http.ResponseWriter, r *http.Request) {
	channels := config.GetChannels()
	for _, ch := range channels {
		_ = process.StartProcess(ch.ID, ch.RunCmd)
	}
	w.Write([]byte("Started all channels"))
}

func StopAllHandler(w http.ResponseWriter, r *http.Request) {
	process.StopAll()
	w.Write([]byte("Stopped all channels"))
}
