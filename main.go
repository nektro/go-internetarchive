package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/nektro/go-util/ansi/style"
	"github.com/spf13/cobra"
)

// Version takes in version string from build_all.sh
var Version = "vMASTER"

var cmdRoot = &cobra.Command{
	Use:   "ia",
	Short: "ia is a cli interface for Archive.org.",
}

func main() {
	dieOnError(cmdRoot.Execute())
}

func dieOnError(err error, args ...string) {
	if err != nil {
		logError(err.Error())
		for _, item := range args {
			logError(item)
		}
		os.Exit(1)
	}
}

func logError(err string) {
	fmt.Print(style.FgRed)
	log.Println(err)
	fmt.Print(style.ResetFgColor)
}
