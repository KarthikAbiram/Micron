/*
Copyright © 2025 KarthikAbiram, MIT License
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"microncli/library"

	"github.com/spf13/cobra"
)

var (
	startNetworkFlag     string
	startServiceIDFlag   string
	startServicePathFlag string
	startTimeout_s       int
)

var startCmd = &cobra.Command{
	Use:   "start [network] [service]",
	Short: "Start a service in a network",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var network, serviceID, path string
		var timeout time.Duration

		// First, try to get values from flags
		network = strings.ToLower(startNetworkFlag)
		serviceID = strings.ToLower(startServiceIDFlag)
		path = strings.ToLower(startServicePathFlag)
		timeout = time.Duration(startTimeout_s) * time.Second

		// If any required value is missing, try to fill from positional args
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}
		if serviceID == "" && len(args) > 1 {
			serviceID = strings.ToLower(args[1])
		}

		if path == "" && len(args) > 2 {
			path = args[2]
		}

		connection, err := library.StartService(network, serviceID, path, timeout)
		fmt.Println(connection.ConnectionString)
		return err
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Define flags
	startCmd.Flags().StringVar(&startNetworkFlag, "network", "", "Network name")
	startCmd.Flags().StringVar(&startServiceIDFlag, "service-id", "", "Service ID")
	startCmd.Flags().StringVar(&startServicePathFlag, "path", "", "Service Path")
	startCmd.Flags().IntVar(&startTimeout_s, "timeout", 30, "Timeout in seconds")
}
