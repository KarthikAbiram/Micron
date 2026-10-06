/*
Copyright © 2025 KarthikAbiram, MIT License
*/
package cmd

import (
	"strings"
	"time"

	"microncli/library"

	"github.com/spf13/cobra"
)

var (
	stopNetworkFlag   string
	stopServiceIDFlag string
	stopTimeout_s     int
)

var stopCmd = &cobra.Command{
	Use:   "stop [network] [service]",
	Short: "stop a service in a network",
	Args:  cobra.ArbitraryArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var network, serviceID string
		var timeout time.Duration

		// First, try to get values from flags
		network = strings.ToLower(stopNetworkFlag)
		serviceID = strings.ToLower(stopServiceIDFlag)
		timeout = time.Duration(stopTimeout_s) * time.Second

		// If any required value is missing, try to fill from positional args
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}
		if serviceID == "" && len(args) > 1 {
			serviceID = strings.ToLower(args[1])
		}

		err := library.StopService(network, serviceID, timeout)
		return err
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)

	// Define flags
	stopCmd.Flags().StringVar(&stopNetworkFlag, "network", "", "Network name")
	stopCmd.Flags().StringVar(&stopServiceIDFlag, "service-id", "", "Service ID")
	stopCmd.Flags().IntVar(&stopTimeout_s, "timeout", 30, "Timeout in seconds")
}
