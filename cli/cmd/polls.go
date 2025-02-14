package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pollsCmd)
}

var pollsCmd = &cobra.Command{
	Use:   "polls",
	Short: "Retreive active polls",
	Long: `Retreives all active polls that a user is able to vote on.
	User can vote using the vote`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Austin's Magic Polling application v0.1")
	},
}
