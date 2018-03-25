package pkg

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Usage interface {
	Handle()
}

type usage struct {
	homeDir           string
	currentDir        string
	goworkUsageActive bool
	goworkOldPath     string
	config            *Configuration
}

func NewUsage() Usage {
	homeDir, err := GetAndCreateHomePath()

	if err != nil {
		OutputShell(fmt.Sprintf("Cannot create homedir '%s': %s", homeDir, err.Error()))
		os.Exit(1)
	}

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		OutputShell(fmt.Sprintf("Error while getting the current path: %s", err.Error()))
		os.Exit(1)
	}

	currentDirHash := GetHashForDirectory(currentDir)

	oldPath := os.Getenv(envVarGoworkOldPath)
	usageActive := true
	if oldPath == "" {
		usageActive = false
		oldPath = os.Getenv("GOPATH")
	}

	usageFlag := flag.NewFlagSet("use", flag.ExitOnError)
	workspaceName := usageFlag.String("name", currentDirHash, "Name of the workspace - if not used the hash of the current directory will be used.")
	usageFlag.Parse(os.Args[2:])

	configFilePath := filepath.Join(homeDir, *workspaceName)
	data, err := ioutil.ReadFile(filepath.Join(configFilePath, ConfigFileName))

	if err != nil {
		OutputShell("Config file not found. Please use gowork create.")
		OutputShell(err.Error())
		os.Exit(1)
	}

	config := &Configuration{}
	err = json.Unmarshal(data, config)
	if err != nil {
		OutputShell(fmt.Sprintf("Error while reading config file: %s", err.Error()))
		os.Exit(1)
	}

	return &usage{
		homeDir:           homeDir,
		currentDir:        currentDir,
		goworkUsageActive: usageActive,
		goworkOldPath:     oldPath,
		config:            config,
	}
}
