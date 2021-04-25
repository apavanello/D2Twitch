package main

import "fmt"

const configFile = "src/config.yaml"

var cfg = Config{}
var currentServerSteamID string
var matchJson Match

func main() {
	cfg.loadConfig(configFile)
	startCron()
	fmt.Println("after Cron")
}
