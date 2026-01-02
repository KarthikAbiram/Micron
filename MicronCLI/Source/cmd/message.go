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
	msgNetwork string
	msgService string
	msgCommand string
	msgPayload string
)

var messageCmd = &cobra.Command{
	Use:   "message [network] [service] [command] [payload]",
	Short: "Message a service in a network",
	Long: `Message a service in a network.

You can use either flags or positional arguments:

Flags style:
  micronCLI message --network mynetwork --service myservice --command mycommand --payload mypayload

Positional style:
  micronCLI message mynetwork myservice mycommand mypayload
`,
	Args: cobra.MaximumNArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		// Use flags if set
		network := strings.ToLower(msgNetwork)
		service := strings.ToLower(msgService)
		command := strings.ToLower(msgCommand)
		payload := strings.ToLower(msgPayload)

		// Fallback to positional args if needed
		if network == "" && len(args) > 0 {
			network = strings.ToLower(args[0])
		}
		if service == "" && len(args) > 1 {
			service = strings.ToLower(args[1])
		}
		if command == "" && len(args) > 2 {
			command = strings.ToLower(args[2])
		}
		if payload == "" && len(args) > 3 {
			payload = strings.ToLower(args[3])
		}

		// List available network and services
		response, err := library.MessageService(network, service, command, payload)
		fmt.Println(response)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)

	messageCmd.Flags().StringVar(&msgNetwork, "network", "", "Network name")
	messageCmd.Flags().StringVar(&msgService, "service", "", "Service name")
	messageCmd.Flags().StringVar(&msgCommand, "command", "", "Command name")
	messageCmd.Flags().StringVar(&msgPayload, "payload", "", "Payload")
}
