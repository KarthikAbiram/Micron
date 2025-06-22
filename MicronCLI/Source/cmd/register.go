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
	networkFlag    string
	serviceFlag    string
	connectionFlag string
)

var registerCmd = &cobra.Command{
	Use:   "register [network] [service] [connection]",
	Short: "Register a service with a connection string",
	Long: `Register a service in a network with its connection string.

You can provide arguments either as flags or as positional arguments:

Flags style:
  micronCLI register --network mynetwork --service myservice --connection localhost:50051

Positional style:
  micronCLI register mynetwork myservice localhost:50051
`,
	Args: cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		var network, service, connection string

		// First, try to get values from flags
		network = strings.ToLower(networkFlag)
		service = strings.ToLower(serviceFlag)
		connection = connectionFlag

		// If any required value is missing, try to fill from positional args
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}
		if service == "" && len(args) > 1 {
			service = strings.ToLower(args[1])
		}
		if connection == "" && len(args) > 2 {
			connection = args[2]
		}

		// Validate required args
		if network == "" || service == "" || connection == "" {
			fmt.Println("Error: network, service, and connection are required.")
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}

		err := library.RegisterService(network, service, connection)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Define flags
	registerCmd.Flags().StringVar(&networkFlag, "network", "", "Network name")
	registerCmd.Flags().StringVar(&serviceFlag, "service", "", "Service name")
	registerCmd.Flags().StringVar(&connectionFlag, "connection", "", "Connection string")
}
