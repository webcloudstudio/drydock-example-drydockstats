// Command ddkstat reports on Drydock's LLM execution log.
package main

import (
	"os"

	"ddkstat/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
