package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	ServiceName string
	HttpPort    string
	Environment string
	Module      string
	AwsSqs      struct {
		QueueURL string
	}
}{}

func ConfigureEnvironment(path string, env ...string) {
	configor.Load(&Config, path+"config/config.json")
}
