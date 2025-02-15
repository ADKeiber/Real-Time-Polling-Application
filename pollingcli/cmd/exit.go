package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(exitCmd)
}

var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Exit the interactive shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Thanks for dropping by! See ya!")
		os.Exit(0)
	},
}
