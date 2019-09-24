package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all active agents",
	Long:  `List all active agents. By default, use PIDs in working directory. If an agent registry is provided, it will be used.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO: implement.")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
