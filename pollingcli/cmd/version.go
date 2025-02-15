package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of the application",
	Long:  `The current version of the application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Austin's Magic Polling application v0.1")
	},
}
