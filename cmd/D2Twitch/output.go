package main

import (
	"fmt"
	"strconv"
)

func writeMatchStats() {

	//testes
	for k := range matchJSON.Teams {
		if teamID := matchJSON.Teams[k].TeamNumber; teamID == 2 {
			fmt.Println("Radiant")
		} else {
			fmt.Println("Dire")
		}

		for kk := range matchJSON.Teams[k].Players {

			player := matchJSON.Teams[k].Players[kk]
			playerName := player.Name
			heroName := heroes[strconv.FormatInt(player.Heroid, 10)].LocalizedName
			kills := strconv.FormatInt(player.KillCount, 10)
			deaths := strconv.FormatInt(player.DeathCount, 10)
			assists := strconv.FormatInt(player.AssistsCount, 10)

			KDA := kills + "/" + deaths + "/" + assists

			fmt.Printf("%v \t %v ==> %v \n", playerName, heroName, KDA)

		}

	}

}
