/*
Copyright © 2025 KarthikAbiram, MIT License
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version = "0.4"

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "micronCLI",
	Short: "A lightweight CLI to manage microservices",
	Long: `MicronCLI allows registering, querying, clearing, and unregistering services.

Usage examples:

  Register a service:
    micronCLI register --network mynetwork --service-id myservice --connection localhost:50051
    micronCLI register mynetwork myservice localhost:50051

  List available networks and services:
	micronCLI list
	micronCLI list --network mynetwork
	micronCLI list mynetwork

  Query a service:
    micronCLI query --network mynetwork --service-id myservice
    micronCLI query mynetwork myservice

  Unregister a service:
    micronCLI unregister --network mynetwork --service-id myservice
    micronCLI unregister mynetwork myservice

  Clear a network:
    micronCLI clear --network mynetwork
    micronCLI clear mynetwork
`,
	Version: Version,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true, // hides cmd
	},
}

// versionCmd represents the version command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		// Only print the version number
		fmt.Println(Version)
	},
}

// Execute is called by main.main(). It runs the CLI.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func init() {
	// Add the version command as a subcommand to rootCmd
	rootCmd.AddCommand(versionCmd)

	// You can define global flags here using PersistentFlags (e.g., a global --config flag).
	// Example:
	// rootCmd.PersistentFlags().String("config", "", "Path to config file")

	// Local flags only apply to this root command if it's invoked directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
