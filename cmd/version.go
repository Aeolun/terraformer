package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.9.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Terraformer",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Terraformer " + version)
	},
}
