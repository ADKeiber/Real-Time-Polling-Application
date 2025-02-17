/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	serverId string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "polling",
	Short: "cli allows for interaction with polls",
	Long: `A command line interface that allows for a user to interact with polls.
	
	Current supported actions include:
		1: Print out currently active polls
		2: Watch a polls results update live
		3: Vote on a poll
		4: Perform CRUD actions on a poll`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// Use the flag value

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	} else {

		// Check if the required flag is provided
		if serverId == "" {
			fmt.Println("Error: The 'server ID' flag is required. Can be set via --server=serverName or -s=serverName")
			os.Exit(1) // Exit with a non-zero status code
		}

		fmt.Printf("Welcome to server with ID '%s'!\n", serverId)

		//addCommands()
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ") // Prompt
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input) // Remove newline and extra spaces

			// Skip empty input
			if input == "" {
				continue
			}

			// Split input into arguments
			args := strings.Split(input, " ")

			// Set the arguments for Cobra
			rootCmd.SetArgs(args)

			// Execute the command
			if err := rootCmd.Execute(); err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}

func init() {

	// Define the required flag
	rootCmd.Flags().StringVarP(&serverId, "server", "s", "", "Server ID (required)")
	// Mark the flag as required
	rootCmd.MarkFlagRequired("server")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.PersistentFlags().StringP("server", "s", "", "ServerId (required)")
	// rootCmd.MarkPersistentFlagRequired("server")
}
