package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Gocmd interface {
	Handle()
}

type gocmd struct {
	goworkUsageActive bool
	goworkOldPath     string
	config            *Configuration
}

func NewGocmd() Gocmd {
	oldPath := os.Getenv(envVarGoworkOldPath)
	usageActive := true
	if oldPath == "" {
		usageActive = false
	}

	config := &Configuration{}
	if usageActive {
		configPath := os.Getenv(envVarGoworkConfigPath)

		data, err := ioutil.ReadFile(configPath)

		if err != nil {
			OutputShell("Config file not found. Please use gowork create.")
			OutputShell(err.Error())
			os.Exit(1)
		}

		err = json.Unmarshal(data, config)
		if err != nil {
			OutputShell(fmt.Sprintf("Error while reading config file: %s", err.Error()))
			os.Exit(1)
		}
	}

	return &gocmd{
		goworkUsageActive: usageActive,
		goworkOldPath:     oldPath,
		config:            config,
	}
}
