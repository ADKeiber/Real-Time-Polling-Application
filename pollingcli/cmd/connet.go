package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a poll and allow for live updates",
	Run:   realTimePolling,
}

func realTimePolling(cmd *cobra.Command, args []string) {
	fmt.Println("Processing:") // Static line
	fmt.Println("Step 1...")   // The line that will be updated

	steps := []string{"Step 1...", "Step 2...", "Step 3...", "Done!"}

	for _, step := range steps {
		time.Sleep(1 * time.Second) // Simulate work
		fmt.Print("\033[A\033[K")   // Move up and clear line
		fmt.Println(step)           // Print the updated text
	}
}
