/*
Copyright © 2025 KarthikAbiram, MIT License
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "micronCLI",
	Short: "A lightweight CLI to manage microservices",
	Long: `MicronCLI allows registering, querying, clearing, and unregistering services.

Usage examples:

  Register a service:
    micronCLI register --network mynetwork --service myservice --connection localhost:50051
    micronCLI register mynetwork myservice localhost:50051

  Query a service:
    micronCLI query --network mynetwork --service myservice
    micronCLI query mynetwork myservice

  Unregister a service:
    micronCLI unregister --network mynetwork --service myservice
    micronCLI unregister mynetwork myservice

  Clear a network:
    micronCLI clear --network mynetwork
    micronCLI clear mynetwork
`,
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true, // hides cmd
		// DisableDefaultCmd: true, // removes cmd
	},
	// Uncomment to add a default action:
	// Run: func(cmd *cobra.Command, args []string) {
	//     fmt.Println("MicronCLI called. Use --help to see available commands.")
	// },
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
	// You can define global flags here using PersistentFlags (e.g., a global --config flag).
	// Example:
	// rootCmd.PersistentFlags().String("config", "", "Path to config file")

	// Local flags only apply to this root command if it's invoked directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
