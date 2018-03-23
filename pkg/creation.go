package pkg

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// Creation interface to provide the functionallity for the creation workflow.
type Creation interface {
	Handle()
}

type creation struct {
	homeDir        string
	repository     string
	currentDir     string
	currentDirHash string
}

// NewCreation is a function, that returns a Creation interface object.
func NewCreation() Creation {
	homeDir, err := GetAndCreateHomePath()
	if err != nil {
		fmt.Printf("Cannot create homedir '%s': %s", homeDir, err.Error())
		os.Exit(1)
	}

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("Error while getting the current path: %s", err.Error())
		os.Exit(1)
	}

	createFlag := flag.NewFlagSet("create", flag.ExitOnError)
	repositoryName := createFlag.String("repo", "", "Imports within this folder will be symlinked into your environment.")
	createFlag.Parse(os.Args[2:])

	currentDirHash := GetHashForDirectory(currentDir)

	return &creation{
		homeDir:        homeDir,
		repository:     *repositoryName,
		currentDir:     currentDir,
		currentDirHash: currentDirHash,
	}
}
