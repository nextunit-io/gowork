package pkg

import "os"

type Deactivate interface {
	Handle()
}

type deactivate struct {
	goworkUsageActive bool
	goworkOldPath     string
}

func NewDeactivate() Deactivate {
	oldPath := os.Getenv(envVarGoworkOldPath)
	usageActive := true
	if oldPath == "" {
		usageActive = false
	}

	return &deactivate{
		goworkUsageActive: usageActive,
		goworkOldPath:     oldPath,
	}
}
