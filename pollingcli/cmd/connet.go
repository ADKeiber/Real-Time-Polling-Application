package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

/*
Testing this
*/
func init() {
	rootCmd.AddCommand(connectCmd)
}

var batchFileName = "websocket_output.bat"
var wsURL = "ws://localhost:8080/ws"

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a poll and allow for live updates",
	Run:   realTimePolling,
}

func realTimePolling(cmd *cobra.Command, args []string) {
	// var newPrompt *exec.Cmd

	// switch runtime.GOOS {
	// case "windows":
	// 	newPrompt = exec.Command("cmd", "/k", "start") // Windows
	// 	err := newPrompt.Start()

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	i := 0

	// 	for {

	// 		data := fmt.Sprintf("Update %d: Hello, World!", i+1)

	// 		newPrompt.Stdout = os.Stdout
	// 		newPrompt.Stderr = os.Stderr

	// 		fmt.Print(data)
	// 		// newPrompt := exec.Command("cmd", "/k", fmt.Sprintf("start; echo %s", data))

	// 		time.Sleep(2 * time.Second)

	// 		i++

	// 		err := newPrompt.Start()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 	}
	// default:
	// 	panic("Unsupported operating system")
	// }

	err := openNewCommandPrompt()
	if err != nil {
		fmt.Println("Error opening command prompt:", err)
		return
	}

	// Step 2: Connect to WebSocket
	fmt.Println("Connecting to WebSocket server at", wsURL)
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		fmt.Println("WebSocket connection error:", err)
		return
	}
	defer conn.Close()

	// Step 3: Listen for messages from WebSocket
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		// Step 4: Append message to batch script
		err = appendMessageToBatch(string(message))
		if err != nil {
			fmt.Println("Error updating batch script:", err)
		}
	}

}

// func openNewCommandPrompt(cmd *cobra.Command, args []string) {
// 	// Create a temporary batch script
// 	batchFileName := "temp_script.bat"
// 	batchFile, err := os.Create(batchFileName)
// 	if err != nil {
// 		fmt.Println("Error creating batch file:", err)
// 		return
// 	}
// 	defer batchFile.Close()

// 	// Write initial command to open CMD and keep it open
// 	batchFile.WriteString("@echo off\n")

// 	// Write messages to be displayed
// 	for i := 1; i <= 10; i++ {
// 		batchFile.WriteString(fmt.Sprintf("echo Message %d\n", i))
// 	}

// 	// Keep the CMD window open
// 	batchFile.WriteString("pause\n")

// 	// Close the batch file to flush data
// 	batchFile.Close()

// 	// Open the new command prompt and execute the batch file
// 	newCmd := exec.Command("cmd", "/C", "start", batchFileName)
// 	err = newCmd.Start()
// 	if err != nil {
// 		fmt.Println("Error opening new cmd:", err)
// 		return
// 	}

// 	// Wait a bit for the command prompt to launch
// 	time.Sleep(2 * time.Second)

// 	fmt.Println("New command prompt opened and printing messages...")
// }

func openNewCommandPrompt() error {
	// Create batch script with initial message
	err := os.WriteFile(batchFileName, []byte("@echo off\n"), 0644)
	if err != nil {
		return err
	}

	// Open a new CMD window and run the batch script
	cmd := exec.Command("cmd", "/C", "start", batchFileName)
	return cmd.Start()
}

func appendMessageToBatch(message string) error {
	// Open the batch script in append mode
	file, err := os.OpenFile(batchFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the new message to the batch script
	_, err = file.WriteString(fmt.Sprintf("echo %s\n", message))
	if err != nil {
		return err
	}

	// Ensure the new CMD window reloads the updated batch script
	_, err = file.WriteString("timeout /t 1 >nul\n") // Short pause to display new messages
	if err != nil {
		return err
	}

	return nil
}
