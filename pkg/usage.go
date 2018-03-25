package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	envVarGoworkOldPath = "GOWORK_GOPATH"
)

type Usage interface {
	Handle()
}

type usage struct {
	homeDir           string
	currentDir        string
	currentDirHash    string
	goworkUsageActive bool
	goworkOldPath     string
	config            *Configuration
}

func NewUsage() Usage {
	homeDir, err := GetAndCreateHomePath()

	if err != nil {
		outputShell(fmt.Sprintf("Cannot create homedir '%s': %s", homeDir, err.Error()))
		os.Exit(1)
	}

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		outputShell(fmt.Sprintf("Error while getting the current path: %s", err.Error()))
		os.Exit(1)
	}

	currentDirHash := GetHashForDirectory(currentDir)

	oldPath := os.Getenv(envVarGoworkOldPath)
	usageActive := true
	if oldPath == "" {
		usageActive = false
		oldPath = os.Getenv("GOPATH")

		if os.PathSeparator == '\\' {
			oldPath = strings.Replace(oldPath, "\\", "/", -1)
		}
	}

	configFilePath := filepath.Join(homeDir, currentDirHash)
	data, err := ioutil.ReadFile(filepath.Join(configFilePath, ConfigFileName))

	if err != nil {
		outputShell("Config file not found. Please use gowork create.")
		outputShell(err.Error())
		os.Exit(1)
	}

	config := &Configuration{}
	err = json.Unmarshal(data, config)
	if err != nil {
		outputShell(fmt.Sprintf("Error while reading config file: %s", err.Error()))
		os.Exit(1)
	}

	return &usage{
		homeDir:           homeDir,
		currentDir:        currentDir,
		currentDirHash:    currentDirHash,
		goworkUsageActive: usageActive,
		goworkOldPath:     oldPath,
		config:            config,
	}
}

func outputShell(output string) {
	fmt.Printf("echo %q", output)
}
