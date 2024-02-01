package main

import (
	"flag"
	"fmt"
)

var (
	configKey     = "configFile"
	defaultConfig = ""
	configUsage   = "this is config file path"
)

func main() {
	//engine := gin.Default()
	//var config config.Config
	//config.GetConfig(&config)
	//
	//router.RegisterRoutes(engine, config)
	//
	//err := engine.Run(fmt.Sprint(config.Server.Host, ":", config.Server.Port))
	//if err != nil {
	//	return
	//}
	//println("Listening and serving at 8080")

	var config string
	flag.StringVar(&config, configKey, defaultConfig, configUsage)
	flag.Parse()
	fmt.Println("flag", flag.Args()[0])
}
