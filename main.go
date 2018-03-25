package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	// TODO: Remove relative path
	pkg "./pkg"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	switch os.Args[1] {
	case "create":
		pkg.NewCreation().Handle()
		break
	case "use":
		pkg.NewUsage().Handle()
		break
	case "deactivate":
		pkg.NewDeactivate().Handle()
		break
	default:
		fmt.Printf("Please use not %s", os.Args[1])
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Printf("Help for ./gowork")
}
