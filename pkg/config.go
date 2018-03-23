package pkg

import (
	"crypto/sha256"
	"encoding/base64"
	"os"
	"path/filepath"
)

var (
	homeDirEnvironmentKey = "GOWORK_HOME"
	homeDirDefault        = ".gowork"

	ConfigFileName = "gowork.json"
	GoPathDir      = "go-source"
)

type Configuration struct {
	Repository     string
	CurrentDir     string
	CurrentDirHash string
	GoPath         string
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
