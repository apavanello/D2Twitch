package main

import (
	"fmt"
	"strconv"
)

func writeMatchStats() {

	//testes
	for k := range matchJson.Teams {
		if teamID := matchJson.Teams[k].TeamNumber; teamID == 2 {
			fmt.Println("Radiant")
		} else {
			fmt.Println("Dire")
		}

		for kk := range matchJson.Teams[k].Players {

			playerName := matchJson.Teams[k].Players[kk].Name
			heroName := heroes[strconv.FormatInt(matchJson.Teams[k].Players[kk].Heroid, 10)].LocalizedName
			kills := strconv.FormatInt(matchJson.Teams[k].Players[kk].KillCount, 10)
			deaths := strconv.FormatInt(matchJson.Teams[k].Players[kk].DeathCount, 10)
			assists := strconv.FormatInt(matchJson.Teams[k].Players[kk].AssistsCount, 10)

			KDA := kills + "/" + deaths + "/" + assists

			fmt.Printf("%v \t %v ==> %v \n", playerName, heroName, KDA)

		}

	}

}
