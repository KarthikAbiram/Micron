/*
Copyright © 2025 KarthikAbiram, MIT License
*/
package cmd

import (
	"MicronCLI/library"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	unregNetwork string
	unregService string
)

var unregisterCmd = &cobra.Command{
	Use:   "unregister [network] [service]",
	Short: "Unregister a service from a network",
	Long: `Unregister a registered service from a specific network.

You can provide input either as flags or positional arguments:

Flag style:
  micronCLI unregister --network mynetwork --service myservice

Positional style:
  micronCLI unregister mynetwork myservice
`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Read flags first
		network := strings.ToLower(unregNetwork)
		serviceName := strings.ToLower(unregService)

		// Fallback to positional args
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}
		if serviceName == "" && len(args) > 1 {
			serviceName = strings.ToLower(args[1])
		}

		// Validate inputs
		if network == "" || serviceName == "" {
			fmt.Println("Error: both network and service are required.")
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}

		// Call the unregister logic
		err := library.UnregisterService(network, serviceName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully unregistered '%s' from network '%s'\n", serviceName, network)
	},
}

func init() {
	rootCmd.AddCommand(unregisterCmd)

	unregisterCmd.Flags().StringVar(&unregNetwork, "network", "", "Network name")
	unregisterCmd.Flags().StringVar(&unregService, "service", "", "Service name")
}
