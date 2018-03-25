package pkg

import (
	"fmt"
	"os"
)

func (c *usage) Handle() {
	if c.goworkUsageActive {
		outputShell(fmt.Sprintf("GOWork already activated."))
		os.Exit(1)
	}

	fmt.Printf("export %s=%q\n", envVarGoworkOldPath, c.goworkOldPath)
	fmt.Printf("export GOPATH=%q\n", c.config.GoExportPath)
	// fmt.Printf("alias deactivate=\"eval `./gowork deactivate`\"\n")
	outputShell("Gowork activated.")
}
