package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "agentctl",
	Short: "Sentinovo AgentKit CLI",
	Long: `Sentinovo AgentKit Client is a command-line interface (CLI) to manage the
  lifecycle of software agents. With this cli, you can list, start, pause, update,
  and stop agents.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("workdir", "w", "./.agent/", "Path to working directory.")
}
