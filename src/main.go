package main

import "fmt"

const configFile = "src/config.yaml"

var cfg = Config{}

func main() {
	cfg.loadConfig(configFile)
	startCron()
	fmt.Println("after Cron")
}
