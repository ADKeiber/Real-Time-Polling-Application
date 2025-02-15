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

// Another example command
var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "Exit the interactive shell",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Goodbye!")
		os.Exit(0)
	},
}

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
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Goodbye!")
	// 	os.Exit(0)
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(exitCmd)

	fmt.Println("Welcome to Austin's Magical Polling Command Line Application! Type 'help' for a list of commands.")
	startInteractiveShell()
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//rootCmd.PersistentFlags().StringP("server", "s", "66", "Server ID (Required)")
	//rootCmd.MarkPersistentFlagRequired("server")
}

func startInteractiveShell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		args := strings.Split(input, " ")

		rootCmd.SetArgs(args)
		if err := rootCmd.Execute(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
