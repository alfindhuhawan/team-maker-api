package main

import (
	"flag"
	"fmt"
	"team-maker-api/application"
	"team-maker-api/shared/driver"
	"team-maker-api/shared/infrastructure/config"
)

func main() {

	// test every config is ready
	config.ReadConfig()

	appMap := map[string]func() driver.RegistryContract{
		"appteammakerapi": application.NewAppTeamMaker(),
	}
	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		driver.Run(app())
	} else {
		fmt.Println("You may try 'go run main.go <app_name>' :")
		for appName := range appMap {
			fmt.Printf(" - %s\n", appName)
		}
	}

}
