package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func (cmd *gocmd) Handle() {
	goPath, err := exec.LookPath("go")
	if err != nil {
		OutputShell("Cannot find go executable")
	}

	if !cmd.goworkUsageActive {
		OutputShell("Gowork is not activated.")
		os.Exit(1)
	}

	repositoryPath := filepath.Join(cmd.config.GoExportPath, cmd.config.Repository)
	cmd.syncDirectory(repositoryPath)

	// Jump into repository path
	fmt.Print("CURRENT_GO_PATH=$(pwd) && ")
	fmt.Printf("cd %q && ", repositoryPath)
	fmt.Printf("%q %s || true && ", goPath, strings.Join(os.Args[2:], " "))
	fmt.Print("cd $CURRENT_GO_PATH ")
}

func (cmd *gocmd) syncDirectory(repositoryPath string) {
	os.MkdirAll(repositoryPath, os.ModePerm)

	fmt.Printf("rm -rf %q && ", repositoryPath)
	fmt.Printf("cp -R . %q && ", repositoryPath)
}
