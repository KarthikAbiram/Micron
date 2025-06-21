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
	queryNetwork string
	queryService string
)

var queryCmd = &cobra.Command{
	Use:   "query [network] [service]",
	Short: "Query a registered service's connection string",
	Long: `Query the connection string of a registered service.

You can use either flags or positional arguments:

Flags style:
  micronCLI query --network mynetwork --service myservice

Positional style:
  micronCLI query mynetwork myservice
`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Use flags if set
		network := strings.ToLower(queryNetwork)
		serviceName := strings.ToLower(queryService)

		// Fallback to positional args if needed
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}
		if serviceName == "" && len(args) > 1 {
			serviceName = strings.ToLower(args[1])
		}

		// Validate
		if network == "" || serviceName == "" {
			fmt.Println("Error: both network and service are required.")
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}

		// Perform query (replace with real logic)
		connStr, err := library.QueryService(network, serviceName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(connStr)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	queryCmd.Flags().StringVar(&queryNetwork, "network", "", "Network name")
	queryCmd.Flags().StringVar(&queryService, "service", "", "Service name")
}
