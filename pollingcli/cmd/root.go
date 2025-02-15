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
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	} else {
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pollingcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
