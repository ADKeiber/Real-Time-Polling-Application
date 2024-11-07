package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

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
		newPrompt = exec.Command("cmd", "/k", "start") // Windows
		err := newPrompt.Start()

		if err != nil {
			panic(err)
		}

		i := 0

		for {

			data := fmt.Sprintf("Update %d: Hello, World!", i+1)

			newPrompt.Stdout = os.Stdout
			newPrompt.Stderr = os.Stderr

			fmt.Print(data)
			// newPrompt := exec.Command("cmd", "/k", fmt.Sprintf("start; echo %s", data))

			time.Sleep(2 * time.Second)

			i++

			err := newPrompt.Start()
			if err != nil {
				panic(err)
			}
		}
	default:
		panic("Unsupported operating system")
	}

}
