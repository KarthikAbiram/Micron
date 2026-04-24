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
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Use flags if set
		network := strings.ToLower(listNetwork)

		// Fallback to positional args if needed
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}

		// List available network and services
		connections, err := library.ListNetworkAndServices(network)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(100)
		}

		if len(connections) == 0 {
			fmt.Println("No active network/services present")
		} else {
			fmt.Printf("%-15s %-15s %-15s %-15s %-15s\n", "Network", "Service", "Connection", "Status", "Info")
			fmt.Println(strings.Repeat("-", 100))

			for _, c := range connections {
				fmt.Printf(
					"%-15s %-15s %-15s %-15d %-25s\n",
					c.Network,
					c.Service,
					c.ConnectionString,
					c.Status,
					c.Info,
				)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVar(&listNetwork, "network", "", "Network name")

	//Ignore unknown flags
	registerCmd.FParseErrWhitelist.UnknownFlags = true
}
