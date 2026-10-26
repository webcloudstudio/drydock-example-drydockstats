// Command ddkstat reports on Drydock's LLM execution log.
//
// The reporting pipeline lives in internal/cli. main only hands the arguments to cli.Run and exits
// with its code; it is the only package in this module that may call os.Exit.
//
// The report package is imported only for its init, which registers each report writer with cli.
package main

import (
	"os"

	"ddkstat/internal/cli"
	_ "ddkstat/internal/report"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
