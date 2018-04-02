package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Handle is a function to create the workspace for the current environment.
func (c *creation) Handle() {
	configFilePath := filepath.Join(c.homeDir, c.workspaceName)
	goPath := filepath.Join(configFilePath, GoPathDir)

	_, err := os.Stat(configFilePath)
	if err == nil || !os.IsNotExist(err) {
		fmt.Printf("There is already a gowork workspace fo this directory. Use ./gowork use")
		os.Exit(1)
	}

	err = os.MkdirAll(goPath, os.ModePerm)
	if err != nil {
		fmt.Printf("Cannot create a configuration directory: %s", err.Error())
		os.Exit(1)
	}

	jsonBody, err := json.Marshal(c.createConfiguration(goPath, configFilePath))
	if err != nil {
		fmt.Printf("Cannot create json body: %s", err.Error())
		os.Exit(1)
	}

	ioutil.WriteFile(filepath.Join(configFilePath, ConfigFileName), jsonBody, 0644)

	fmt.Printf("Created a workspace. (%s)", c.workspaceName)
}

func (c *creation) createConfiguration(goPath, goExportPath string) *Configuration {
	return &Configuration{
		CurrentDir:    c.currentDir,
		WorkspaceName: c.workspaceName,
		Repository:    c.repository,
		GoPath:        goPath,
		GoExportPath:  goExportPath,
	}
}
