package main

import (
	"github.com/nektro/internetarchive/pkg/cmd"
	. "github.com/nektro/internetarchive/pkg/util"
)

// Version takes in version string from build_all.sh
var Version = "vMASTER"

func main() {
	DieOnError(cmd.Root.Execute())
}
