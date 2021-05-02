package main

import (
	"encoding/json"
	"path/filepath"
)

// GameMode Struct
type GameMode map[string]GameModeValue

// GameModeValue Struct
type GameModeValue struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Balanced *bool  `json:"balanced,omitempty"`
}

func (g GameMode) loadConst() {
	f := readFile(filepath.Join("data", "dotaConst", "gameMode.json"))
	if err := json.Unmarshal(f, &gameMode); err != nil {
		processError(err)
	}
}
