package main

import (
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Config struct
type Config struct {
	Steam   Steam   `yaml:"steam"`
	Dataset Dataset `yaml:"dataset"`
	Output  Output  `yaml:"output"`
	Cron    Cron    `yaml:"cron"`
	Mock    Mock    `yaml:"mock"`
}

// Steam struct
type Steam struct {
	APIKey string `yaml:"apiKey"`
	UserID string `yaml:"userID"`
}

// Dataset struct
type Dataset struct {
	Console     Console `yaml:"console"`
	SteamWebAPI string  `yaml:"steamWebApi"`
	OpenDotaAPI string  `yaml:"openDotaApi"`
}

// Console struct
type Console struct {
	ClearAfterRead bool   `yaml:"clearAfterRead"`
	Path           string `yaml:"path"`
}

// Output struct
type Output struct {
	MatchStats  string `yaml:"matchStats"`
	PlayerStats string `yaml:"playerStats"`
}

// Cron struct
type Cron struct {
	GetServerSteamID string `yaml:"getServerSteamID"`
	WriteMatchStats  string `yaml:"writeMatchStats"`
	GetMatchStats    string `yaml:"getMatchStats"`
}

// Mock struct
type Mock struct {
	Enable    bool      `yaml:"enable"`
	Endpoints Endpoints `yaml:"endpoints"`
}

// Endpoints struct
type Endpoints struct {
	MatchStats string `yaml:"matchStats"`
}

// Metodo loadConfig
func (c *Config) loadConfig() {
	var scriptPath, err = os.Getwd()
	if err != nil {
		processError(err)
	}
	configFile := filepath.Join(scriptPath, "config.yaml")

	// Abre o Arquivo
	f := openFile(configFile)

	// Fecha o arquivo no Final
	defer closeFile(f)

	// Instancia o decoder usando o arquivo aberto como origem
	decoder := yaml.NewDecoder(f)

	// Decodifica o YAML e joga para o apontamento em c (CONFIG STRUCT)
	err = decoder.Decode(&c)

	// Loga se der erro
	if err != nil {
		processError(err)
	}
}
