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
	homeDir       string
	repository    string
	currentDir    string
	workspaceName string
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

	currentDirHash := GetHashForDirectory(currentDir)

	createFlag := flag.NewFlagSet("create", flag.ExitOnError)
	repositoryName := createFlag.String("repo", "", "Imports within this folder will be symlinked into your environment.")
	workspaceName := createFlag.String("name", currentDirHash, "Name of the workspace - if not used the hash of the current directory will be used.")
	createFlag.Parse(os.Args[2:])

	return &creation{
		homeDir:       homeDir,
		repository:    *repositoryName,
		currentDir:    currentDir,
		workspaceName: *workspaceName,
	}
}
