package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
)

func startCron() {
	c := cron.New(cron.WithSeconds())
	if _, err := c.AddFunc(cfg.Cron.GetServerSteamID, runGetServerSteamID); err != nil {
		processError(err)
	}
	if _, err := c.AddFunc(cfg.Cron.GetMatchStats, runGetMatchJSON); err != nil {
		processError(err)
	}
	if _, err := c.AddFunc(cfg.Cron.WriteMatchStats, runWriteMatchStats); err != nil {
		processError(err)
	}
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func runGetServerSteamID() {
	currentServerSteamID = consoleLog()
	runGetMatchJSON()
	fmt.Printf("%v\n", currentServerSteamID)

}

func runGetMatchJSON() {
	if currentServerSteamID != "" {
		getJSONFromWebAPI(currentServerSteamID)
		writeMatchStats()
	}
}

func runWriteMatchStats() {
	//	writeMatchStats()
}
