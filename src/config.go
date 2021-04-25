package main

import (
	"gopkg.in/yaml.v2"
)

// Config
type Config struct {
	Output  Output  `yaml:"output"`
	Steam   Steam   `yaml:"steam"`
	Dataset Dataset `yaml:"dataset"`
	Cron	Cron	`yaml:"cron"`

}

// Output
type Output struct {
	MatchStats  string `yaml:"matchStats"`
	PlayerStats string `yaml:"playerStats"`
}

// Steam
type Steam struct {
	ApiKey string `yaml:"apiKey"`
	UserID string `yaml:"userID"`
}

// Dataset
type Dataset struct {
	Console     Console `yaml:"console"`
	SteamWebApi string  `yaml:"steamWebApi"`
	OpenDotaApi string  `yaml:"openDotaApi"`
}

// Cron
type Cron struct {
	GetServerSteamID string `yaml:"getServerSteamID"`
	GetMatchStats    string `yaml:"getMatchStats"`
}

// Console
type Console struct {
	Path           string `yaml:"path"`
	ClearAfterRead bool   `yaml:"clearAfterRead"`
}

// Metodo loadConfig
func (c *Config) loadConfig(configFile string) {

	// Abre o Arquivo
	f := openFile(configFile)

	// Fecha o arquivo no Final
	defer closeFile(f)

	// Instancia o decoder usando o arquivo aberto como origem
	decoder := yaml.NewDecoder(f)

	// Decodifica o YAML e joga para o apontamento em c (CONFIG STRUCT)
	err := decoder.Decode(&c)

	// Loga se der erro
	if err != nil {
		processError(err)
	}
}