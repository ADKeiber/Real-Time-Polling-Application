package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a poll and allow for live updates",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Connected to a specific poll!")
		//need to have a loop here displaying the information of the poll and allow for a response
	},
}
