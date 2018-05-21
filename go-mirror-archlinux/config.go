package main

import (
	"encoding/json"
	"log"
	"os"
)

var (
	config *Config
)

// Config : Service configuration struct
type Config struct {
	Port           int    `json:"PORT"`
	RepoDirectory  string `json:"REPO_DIRECTORY"`
	PrimaryServer  string `json:"PRIMARY_SERVER"`
	BackupServer   string `json:"BACKUP_SERVER"`
	BandwidthLimit int    `json:"BANDWIDTH_LIMIT_KB"`
	SyncInterval   int    `json:"SYNC_INTERVAL_HOURS"`
	SyncISO        bool   `json:"SYNC_ISO"`
	SyncOther      bool   `json:"SYNC_OTHER"`
	SyncSources    bool   `json:"SYNC_SOURCES"`
}

func loadConfig(configPath string) error {
	confFile, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&config)
	return nil
}
