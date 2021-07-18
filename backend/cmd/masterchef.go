package cmd

import (
	"github.com/cosasdepuma/masterchef/backend/app/config"
	"github.com/cosasdepuma/masterchef/backend/app/server"
)

// Start the server
func Start() {
	var conf config.Config = config.Get()
	server.Start(&conf)
}
