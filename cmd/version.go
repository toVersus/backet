package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Backet",
	Long:  `Print the version number of Backet based on the Semantic Versioning 2.0.0`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Backet v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
