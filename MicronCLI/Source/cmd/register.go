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

var (
	networkFlag    string
	serviceIDFlag  string
	connectionFlag string
	statusFlag     int
	infoFlag       string
)

var registerCmd = &cobra.Command{
	Use:   "register [network] [service] [connection]",
	Short: "Register a service with a connection string",
	Long: `Register a service in a network with its connection string.

You can provide arguments either as flags or as positional arguments:

Flags style:
  micronCLI register --network mynetwork --service-id myservice --connection localhost:50051 --status 0 --info "Sample Info"

Positional style:
  micronCLI register mynetwork myservice localhost:50051 0 "Sample Info"
`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var network, serviceID, connection string
		var status int
		var info string

		// First, try to get values from flags
		network = strings.ToLower(networkFlag)
		serviceID = strings.ToLower(serviceIDFlag)
		connection = connectionFlag
		status = statusFlag
		info = infoFlag

		// If any required value is missing, try to fill from positional args
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}
		if serviceID == "" && len(args) > 1 {
			serviceID = strings.ToLower(args[1])
		}
		if connection == "" && len(args) > 2 {
			connection = args[2]
		}

		// Validate required args
		if network == "" || serviceID == "" || (connection == "" && status == 0) {
			fmt.Println("Error: network, service, and connection/status are required.")
			fmt.Println(cmd.UsageString())
			os.Exit(100)
		}

		err := library.RegisterService(network, serviceID, connection, status, info)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(100)
		}

		// fmt.Println("Service successfully registered!")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Define flags
	registerCmd.Flags().StringVar(&networkFlag, "network", "", "Network name")
	registerCmd.Flags().StringVar(&serviceIDFlag, "service-id", "", "Service ID")
	registerCmd.Flags().StringVar(&connectionFlag, "connection", "", "Connection string")
	registerCmd.Flags().IntVar(&statusFlag, "status", 0, "Service status (default: 0)")
	registerCmd.Flags().StringVar(&infoFlag, "info", "", "Additional info about the service")

	// Ignore unknown flags
	registerCmd.FParseErrWhitelist.UnknownFlags = true
}
