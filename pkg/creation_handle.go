package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

// Handle is a function to create the workspace for the current environment.
func (c *creation) Handle() {
	configFilePath := filepath.Join(c.homeDir, c.currentDirHash)
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

	// Create symlink if repository name is set.
	if c.repository != "" {
		c.createRepositorySymlink(goPath)
	}

	fmt.Printf("Created a workspace here. (%s)", c.currentDirHash)
}

func (c *creation) createConfiguration(goPath, goExportPath string) *Configuration {
	return &Configuration{
		CurrentDir:     c.currentDir,
		CurrentDirHash: c.currentDirHash,
		Repository:     c.repository,
		GoPath:         goPath,
		GoExportPath:   goExportPath,
	}
}

func (c *creation) createRepositorySymlink(goPath string) {
	repositoryPath := filepath.Join(goPath, c.repository)

	os.MkdirAll(repositoryPath, os.ModePerm)
	os.Remove(repositoryPath)

	if runtime.GOOS != "windows" {
		err := os.Symlink(c.currentDir, repositoryPath)
		if err != nil {
			fmt.Printf("Cannot create repository symlink, but I will continue: %s", err.Error())
			fmt.Println()
		}
	} else {
		// Wired Windows needs admin rights to create symlinks...
		fmt.Println("Please execute the following command via cmd with administrator access:")
		fmt.Printf("mklink /J \"%s\" \"%s\"", repositoryPath, c.currentDir)
		fmt.Println()
	}
}
