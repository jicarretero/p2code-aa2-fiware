package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jicarretero/p2code-aa2-fiware/brokerld"
	"github.com/jicarretero/p2code-aa2-fiware/config"
	"github.com/jicarretero/p2code-aa2-fiware/idm"
	"github.com/jicarretero/p2code-aa2-fiware/mqtt"
)

// main function to initialize the application
func main() {
	configPath := "config/config.toml"
	config, err := config.ReadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	idm.SetConfig(config)

	if len(os.Args) > 1 && os.Args[1] == "token" {
		fmt.Println(idm.GetOIDCV4Token())
	} else {
		err = mqtt.Subscribe(config, brokerld.MapDeviceProfile)
		brokerld.SetConfig(config)

		if err != nil {
			fmt.Println(err)
		} else {
			select {}
		}
	}
}
