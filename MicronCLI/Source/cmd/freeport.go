/*
Copyright © 2025 KarthikAbiram, MIT License
*/
package cmd

import (
	"fmt"
	"microncli/library"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var freeportPrefer int32

var freeportCmd = &cobra.Command{
	Use:   "freeport [prefer]",
	Short: "Finds an available TCP port",
	Long: `Finds an available TCP port on the system.

You can use either a flag or a positional argument:

Flag style:
  micronCLI freeport --prefer 50051

Positional style:
  micronCLI freeport 50051

If the preferred port is omitted or 0, the utility will find any random free port. 
If the preferred port is occupied, it will automatically fall back to finding a free port.`,
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		prefer := int(freeportPrefer)

		// Fallback to positional argument if flag isn't set
		if prefer == 0 && len(args) > 0 {
			if val, err := strconv.Atoi(args[0]); err == nil {
				prefer = val
			} else {
				fmt.Fprintln(os.Stderr, "Error: 'prefer' argument must be an integer")
				os.Exit(1)
			}
		}

		// Call the library function
		port, err := library.FreePort(prefer)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error finding free port:", err)
			os.Exit(1)
		}

		// Print only the port number to stdout so it can be easily grabbed in bash scripts
		fmt.Println(port)
	},
}

func init() {
	rootCmd.AddCommand(freeportCmd)

	freeportCmd.Flags().Int32Var(&freeportPrefer, "prefer", 0, "Preferred port number to check first")
}
