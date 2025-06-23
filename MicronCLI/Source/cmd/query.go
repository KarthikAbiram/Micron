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
	queryNetwork   string
	queryServiceID string
)

var queryCmd = &cobra.Command{
	Use:   "query [network] [service]",
	Short: "Query a registered service's connection string",
	Long: `Query the connection string of a registered service.

You can use either flags or positional arguments:

Flags style:
  micronCLI query --network mynetwork --serviceID myservice

Positional style:
  micronCLI query mynetwork myservice
`,
	Args: cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Use flags if set
		network := strings.ToLower(queryNetwork)
		serviceID := strings.ToLower(queryServiceID)

		// Fallback to positional args if needed
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}
		if serviceID == "" && len(args) > 1 {
			serviceID = strings.ToLower(args[1])
		}

		// Validate
		if network == "" || serviceID == "" {
			fmt.Println("Error: both network and service are required.")
			fmt.Println(cmd.UsageString())
			os.Exit(1)
		}

		// Perform query
		connStr, _ := library.QueryService(network, serviceID)
		//Ignore error, which would occur if service has not been registered

		// connStr, err := library.QueryService(network, serviceID)
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// 	os.Exit(1)
		// }
		fmt.Println(connStr)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	queryCmd.Flags().StringVar(&queryNetwork, "network", "", "Network name")
	queryCmd.Flags().StringVar(&queryServiceID, "serviceID", "", "Service name")
}
