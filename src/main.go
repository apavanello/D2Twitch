package main

import "fmt"

const configFile = "src/config.yaml"

var cfg = Config{}
var currentServerSteamID string
var matchJson Match
var heroes Heroes
var gameMode GameMode

func main() {
	cfg.loadConfig(configFile)
	loadAllConst()
	startCron()
	fmt.Println("after Cron")
}
