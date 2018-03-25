package pkg

import (
	"fmt"
	"os"
)

func (u *usage) Handle() {
	if u.goworkUsageActive {
		OutputShell("GOWork already activated.")
		os.Exit(1)
	}

	OutputShell("Gowork activated.")
	fmt.Print(" && ")
	fmt.Printf("export %s=%q\n", envVarGoworkOldPath, u.goworkOldPath)
	fmt.Printf("export GOPATH=%q\n", u.config.GoExportPath)
	fmt.Printf("&& shopt -s expand_aliases && alias deactivate=%q\n", "eval `./gowork deactivate`")
}
