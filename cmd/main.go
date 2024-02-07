package main

import (
	"flag"
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

	commands.Execute(flag.Args()[0], configFile)
}
