package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"strconv"
)

func startCron() {
	c := cron.New(cron.WithSeconds())
	if _, err := c.AddFunc(cfg.Cron.GetServerSteamID, runGetServerSteamID); err != nil {
		processError(err)
	}
	if _, err := c.AddFunc(cfg.Cron.GetMatchStats, runGetMatchJson); err != nil {
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
	runGetMatchJson()
	fmt.Printf("%v\n", currentServerSteamID)

}

func runGetMatchJson() {
	if currentServerSteamID != "" {
		getJsonFromWebAPI(currentServerSteamID)

		for k, _ := range matchJson.Teams {
			if teamID := matchJson.Teams[k].TeamNumber; teamID == 2 {
				fmt.Println("Radiant")
			} else {
				fmt.Println("Dire")
			}

			for kk, _ := range matchJson.Teams[k].Players {
				playerName := matchJson.Teams[k].Players[kk].Name
				heroName := heroes[strconv.FormatInt(matchJson.Teams[k].Players[kk].Heroid, 10)].LocalizedName
				kills := strconv.FormatInt(matchJson.Teams[k].Players[kk].KillCount, 10)
				deaths := strconv.FormatInt(matchJson.Teams[k].Players[kk].DeathCount, 10)
				assists := strconv.FormatInt(matchJson.Teams[k].Players[kk].AssistsCount, 10)
				KDA := kills + "/" + deaths + "/" + assists

				fmt.Printf("%v \t\t %v ==> %v ", playerName, heroName, KDA)
				fmt.Println()
			}

		}
	}
}

func runWriteMatchStats() {

}
