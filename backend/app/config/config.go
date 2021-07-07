package config

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config contains all the information related to the variables necessary to customize the program execution
	Config struct {
		// Server contains the information related to the web service
		Server Server
		// Threads indicates the maximum number of concurrent executions per API call
		Threads int `envconfig:"MC_THREADS" default:"100"`
	}

	// Server contains the information related to the web service
	Server struct {
		Address string `envconfig:"MC_ADDR" default:"localhost:7767"`
	}
)

// Get returns the configuration required for program execution
func Get() Config {
	var config Config = Config{}
	_ = envconfig.Process("", &config)
	return config
}
