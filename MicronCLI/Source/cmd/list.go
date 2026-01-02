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
	listNetwork string
)

var listCmd = &cobra.Command{
	Use:   "list or list [network]",
	Short: "List available networks and services",
	Long: `List available networks and services.

You can use either flags or positional arguments:

Flags style:
  micronCLI list --network mynetwork

Positional style:
  micronCLI list mynetwork
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Use flags if set
		network := strings.ToLower(listNetwork)

		// Fallback to positional args if needed
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}

		// List available network and services
		connections, err := library.ListNetworkAndServices(network)
		if len(connections) == 0 {
			fmt.Println("No active network/services present")
		} else {
			fmt.Printf("%-15s %-15s %s\n", "Network", "Service", "Connection")
			fmt.Println(strings.Repeat("-", 60))

			for _, c := range connections {
				fmt.Printf(
					"%-15s %-15s %s\n",
					c.Network,
					c.Service,
					c.ConnectionString,
				)
			}
		}

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVar(&listNetwork, "network", "", "Network name")
}
