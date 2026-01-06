/*
Copyright © 2025 KarthikAbiram, MIT License
*/
package cmd

import (
	"fmt"
	"microncli/library"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var clearNetwork string

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear [network]",
	Short: "Clear (remove) all services from a network",
	Long: `Clear/remove all services registered in a specified network.

Usage examples:

  micronCLI clear --network mynetwork
  micronCLI clear mynetwork
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		network := strings.ToLower(clearNetwork)

		// Fallback to positional arg if not set via flag
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}

		if network == "" {
			fmt.Println("Error: network name is required.")
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}

		// Perform clear operation
		err := library.Clear(network)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully cleared network '%s'\n", network)
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	clearCmd.Flags().StringVar(&clearNetwork, "network", "", "Network name to clear")
}
