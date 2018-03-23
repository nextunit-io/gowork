package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

type Usage interface {
	Handle() error
}

type usage struct {
	homeDir    string
	currentDir string
}

func NewUsage() Usage {
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

	return &usage{
		homeDir:    homeDir,
		currentDir: currentDir,
	}
}
