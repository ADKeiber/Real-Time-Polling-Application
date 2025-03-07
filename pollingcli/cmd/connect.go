package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-stomp/stomp"
	"github.com/spf13/cobra"
)

/*
Testing this it doesn't currently work as expected
*/
func init() {
	rootCmd.AddCommand(connectCmd)
}

var batchFileName = "websocket_output.bat"
var wsURL = "ws://localhost:8080/ws-guide-websocket"

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a poll and allow for live updates",
	Run:   realTimePolling,
}

func realTimePolling(cmd *cobra.Command, args []string) {
	//Step 1: open new command prompt
	err := openNewCommandPrompt()
	if err != nil {
		fmt.Println("Error opening command prompt:", err)
		return
	}

	// Step 2: Connect to WebSocket
	fmt.Println("Connecting to WebSocket server at", wsURL)

	conn, err := stomp.Dial("tcp", "localhost:8080", nil)

	if err != nil {
		fmt.Println("WebSocket connection error:", err)
		return
	}
	defer conn.Disconnect()

	sub, err := conn.Subscribe("/topic/greetings", stomp.AckAuto)

	if err != nil {
		fmt.Println("Failed to subscribe to destination: %v", err)
	}
	defer sub.Unsubscribe()

	// Step 3: Listen for messages from WebSocket
	for {
		msg := <-sub.C
		if err != nil {
			fmt.Println("WebSocket read error:", err)
			break
		}

		// Step 4: Append message to batch script
		err = appendMessageToBatch(string(msg.Body))
		if err != nil {
			fmt.Println("Error updating batch script:", err)
		}
	}

}

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
