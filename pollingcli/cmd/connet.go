package cmd

import (
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

/*
Testing this
*/
func init() {
	rootCmd.AddCommand(connectCmd)
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a poll and allow for live updates",
	Run:   realTimePolling,
}

func realTimePolling(cmd *cobra.Command, args []string) {
	var newPrompt *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		newPrompt = exec.Command("cmd", "/k", "start; echo Hello World && pause >nul") // Windows
	default:
		panic("Unsupported operating system")
	}

	err := newPrompt.Start()
	if err != nil {
		panic(err)
	}
}
