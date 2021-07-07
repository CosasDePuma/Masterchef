package cmd

import (
	"cosasdepuma/masterchef/app/config"
	"cosasdepuma/masterchef/app/server"
)

// Start the server
func Start() {
	var conf config.Config = config.Get()
	server.Start(&conf)
}
