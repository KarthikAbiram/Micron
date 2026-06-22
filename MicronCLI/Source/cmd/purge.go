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

var purgeKeep int32

var purgeCmd = &cobra.Command{
	Use:   "purge [keep]",
	Short: "Purges old system logs, keeping only the most recent N records",
	Run: func(cmd *cobra.Command, args []string) {
		keep := int(purgeKeep)

		// Overwrite default flag value if a positional argument is provided
		if len(args) > 0 {
			if val, err := strconv.Atoi(args[0]); err == nil {
				keep = val
			} else {
				fmt.Fprintln(os.Stderr, "Error: 'keep' argument must be an integer")
				os.Exit(1)
			}
		}

		// Ensure we don't accidentally pass a negative value to the database query
		if keep < 0 {
			fmt.Fprintln(os.Stderr, "Error: Number of logs to keep cannot be negative")
			os.Exit(1)
		}

		// Call the library function
		deletedCount, err := library.PurgeLogs(keep)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error purging logs:", err)
			os.Exit(1)
		}

		// Success output
		fmt.Printf("Successfully purged old logs. Kept the %d most recent entries.\n", keep)
		fmt.Printf("Rows removed: %d\n", deletedCount)
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	// Defaulting to keeping 1000 logs if no argument/flag is passed
	purgeCmd.Flags().Int32Var(&purgeKeep, "keep", 1000, "Number of recent logs to retain")
}
