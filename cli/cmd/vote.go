package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(voteCmd)
}

var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "Vote on an active poll",
	Long:  `Retreives all active polls that a user is able to vote on.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Austin's Magic Polling application v0.1")
	},
}
