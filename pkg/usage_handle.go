package pkg

import (
	"fmt"
	"os"
	"strings"
)

func (c *usage) Handle() error {
	fmt.Print(getPathEntry())
	return nil
}

func getPathEntry() string {
	env := os.Environ()
	for i := range env {
		if strings.HasPrefix(env[i], "PATH=") {
			return env[i][5:]
		}
	}

	return ""
}
