package main

import (
	"regexp"
)

const regex = `(?m)\((?P<steamID>\d{17})\)`

func consoleLog() string {

	re := regexp.MustCompile(regex)
	str := string(readFile(cfg.Dataset.Console.Path))
	match := re.FindAllStringSubmatch(str, -1)
	return match[len(match) -1][1]

}