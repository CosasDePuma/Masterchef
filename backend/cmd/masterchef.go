package cmd

import (
	"github.com/cosasdepuma/misterchef/backend/app/config"
	"github.com/cosasdepuma/misterchef/backend/app/server"
)

// Start the server
func Start() {
	var conf config.Config = config.Get()
	server.Start(&conf)
}
