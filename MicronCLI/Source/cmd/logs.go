/*
Copyright © 2025 KarthikAbiram, MIT License
*/
package cmd

import (
	"fmt"
	"microncli/library"
	"os"
	"strconv"
	"text/tabwriter" // Standard library for table formatting

	"github.com/spf13/cobra"
)

var logslimit int32

var logsCmd = &cobra.Command{
	Use:   "logs [limit]",
	Short: "Returns last n logs in table format",
	Run: func(cmd *cobra.Command, args []string) {
		limit := int(logslimit)

		if len(args) > 0 {
			if val, err := strconv.Atoi(args[0]); err == nil {
				limit = val
			}
		}

		logdata, err := library.GetLogs(limit)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error fetching logs:", err)
			os.Exit(1)
		}

		// Initialize TabWriter
		// minwidth, tabwidth, padding, padchar, flags
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)

		// Print Table Header
		fmt.Fprintln(w, "ID\tTIMESTAMP\tOPERATION\tNETWORK\tSERVICE\tCONNECTION\tSTATUS\tCALLER")
		fmt.Fprintln(w, "--\t---------\t---------\t-------\t-------\t----------\t------\t------")

		// Print Rows
		for _, log := range logdata {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\t%s\t%d\t%s\n",
				log.ID,
				log.Timestamp,
				log.Operation,
				log.Network,
				log.Service,
				log.Connection,
				log.Status,
				log.CallerID,
			)
		}

		// Flush buffer to screen
		w.Flush()

		fmt.Printf("\n(%d entries shown)\n", len(logdata))
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
	logsCmd.Flags().Int32Var(&logslimit, "limit", 100, "Number of Logs to return")
}
