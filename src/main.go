package main

import (
	"fmt"
)

var cfg = Config{}
var currentServerSteamID string
var matchJson Match
var heroes Heroes
var gameMode GameMode

func main() {
	cfg.loadConfig()
	loadAllConst()
	startCron()
	fmt.Println("after Cron")
}
