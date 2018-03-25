package pkg

import (
	"fmt"
	"os"
)

func (d *deactivate) Handle() {
	if !d.goworkUsageActive {
		OutputShell("Gowork is not activated.")
		os.Exit(1)
	}

	OutputShell("Gowork deactivated.")
	fmt.Print("shopt -s expand_aliases && unalias deactivate\n")
	fmt.Printf("export %s=\"\"\n", envVarGoworkOldPath)
	fmt.Printf("export GOPATH=%q\n", d.goworkOldPath)
}
