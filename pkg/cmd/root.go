package cmd

import (
	"github.com/nektro/internetarchive/pkg/cmd/download"
	"github.com/nektro/internetarchive/pkg/cmd/metadata"

	"github.com/spf13/cobra"
)

// Root is the root command
var Root = &cobra.Command{
	Use:   "ia",
	Short: "ia is a cli interface for Archive.org.",
}

func init() {
	Root.AddCommand(
		metadata.Cmd,
		download.Cmd,
	)
}
