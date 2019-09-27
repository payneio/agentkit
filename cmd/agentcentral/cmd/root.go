package cmd

import (
	"io/ioutil"
	"os"

	"agentkit/pkg/agentkit/central"

	"cuelang.org/go/cue"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "agentcentral",
	Short: "Launch an agent central.",
	Long:  `Launch an agent central.`,
	Run: func(cmd *cobra.Command, args []string) {

		var (
			r          cue.Runtime
			config     *cue.Instance
			configData = []byte{}
			err        error
		)

		// Read in configData if a path was provided
		configPath, _ := cmd.Flags().GetString("config")
		if configPath != "" {
			configData, err = ioutil.ReadFile(configPath)
			if err != nil {
				log.Error(err)
				return
			}
		}

		// Compile configuration
		config, err = r.Compile("agent", configData)
		if err != nil {
			log.Error(err)
			return
		}

		// Assign port
		port, _ := cmd.Flags().GetInt(`port`)
		if port == 0 {
			port = 9100
		}
		config, _ = config.Fill(port, `_port`)

		central := central.New(config)
		central.Spin()

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("config", "c", "", "Path to agent configuration.")
	rootCmd.Flags().IntP("port", "p", 0, "Port to connect agent HTTP-JSON API to.")
}
