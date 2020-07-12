package cmd

import (
	"github.com/spf13/cobra"
)

// Root is the root command
var Root = &cobra.Command{
	Use:   "ia",
	Short: "ia is a cli interface for Archive.org.",
}
