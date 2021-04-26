package main

import "encoding/json"

type GameMode map[string]GameModeValue

type GameModeValue struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Balanced *bool  `json:"balanced,omitempty"`
}

func (g GameMode) loadConst() {
	f := readFile("data/dotaConst/gameMode.json")
	if err := json.Unmarshal(f, &gameMode); err != nil {
		processError(err)
	}
}