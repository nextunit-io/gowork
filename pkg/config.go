package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
)

var (
	homeDirEnvironmentKey  = "GOWORK_HOME"
	homeDirDefault         = ".gowork"
	envVarGoworkOldPath    = "GOWORK_GOPATH"
	envVarGoworkConfigPath = "GOWORK_CONFIG_PATH"

	// ConfigFileName is a vaiarble to configure the filename of othe workon config file.
	ConfigFileName = "gowork.json"

	// GoPathDir is a variable where the source is stored.
	GoPathDir = "src"
)

type Configuration struct {
	Repository    string
	CurrentDir    string
	WorkspaceName string
	GoPath        string
	GoExportPath  string
}

func GetHomePath() string {
	homeDir := os.Getenv(homeDirEnvironmentKey)

	if homeDir == "" {
		homeDir = filepath.Join(os.Getenv("HOME"), homeDirDefault)
	}

	return homeDir
}

func GetAndCreateHomePath() (string, error) {
	homeDir := GetHomePath()

	err := os.MkdirAll(homeDir, os.ModePerm)

	return homeDir, err
}

func GetHashForDirectory(dir string) string {
	h := sha256.New()
	h.Write([]byte(dir))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func OutputShell(output string) {
	fmt.Printf("echo %q\n", output)
}
