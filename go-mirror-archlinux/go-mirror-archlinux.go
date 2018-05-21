package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	configPath := flag.String("config", "", "The path to your config.json file.")
	flag.Parse()
	if *configPath == "" {
		log.Fatal("Please specify a config path")
	}
	fmt.Println("Loading config...")
	err := loadConfig(*configPath)
	if err != nil {
		fmt.Println("Unable to load config.json (does it exist?)")
		return
	}
	fmt.Println("Creating directory (if not preexisting)")
	os.MkdirAll(config.RepoDirectory, 0770)
	fmt.Println("Starting service...")
	serve()
}
