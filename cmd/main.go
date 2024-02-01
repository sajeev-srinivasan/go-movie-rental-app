package main

import (
	"flag"
	"fmt"
	"movie-rental-app/cmd/commands"
)

const (
	configKey     = "configFile"
	defaultConfig = ""
	configUsage   = "this is config file path"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, configKey, defaultConfig, configUsage)
	flag.Parse()
	fmt.Println("In main")

	commands.Execute(flag.Args()[0], configFile)
}
