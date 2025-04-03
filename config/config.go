// config/config.go
package config

import (
	"encoding/json"
	"os"
	"sync"
)

type ChannelConfig struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	RunCmd  string `json:"run_cmd"`
	Port    int    `json:"port"`
}

var (
	Channels []ChannelConfig
	mu       sync.RWMutex
)

func LoadChannels(path string) error {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var ch []ChannelConfig
	err = json.Unmarshal(file, &ch)
	if err != nil {
		return err
	}

	Channels = ch
	return nil
}

func GetChannels() []ChannelConfig {
	mu.RLock()
	defer mu.RUnlock()
	return Channels
}

func GetChannelByID(id string) *ChannelConfig {
	mu.RLock()
	defer mu.RUnlock()
	for _, ch := range Channels {
		if ch.ID == id {
			return &ch
		}
	}
	return nil
}
