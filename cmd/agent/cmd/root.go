package cmd

import (
	"agentkit/pkg/agentkit/agent"
	"agentkit/pkg/agentkit/util"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"cuelang.org/go/cue"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "agent",
	Short: "Start a Sentinovo Agent",
	Long:  `Start a Sentinovo Agent`,
	Run: func(cmd *cobra.Command, args []string) {

		// Logging
		log.SetReportCaller(true)
		log.New().Formatter = &log.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				repopath := fmt.Sprintf("%s/src/github.com/bob", os.Getenv("GOPATH"))
				filename := strings.Replace(f.File, repopath, "", -1)
				return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
			},
		}
		log.SetFormatter(&log.JSONFormatter{})

		// Read in configuration
		configPath, _ := cmd.Flags().GetString("config")
		configData, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Println(err)
			return
		}

		// Compile configuration
		var r cue.Runtime
		config, err := r.Compile("agent", configData)
		if err != nil {
			log.Error(err)
			return
		}

		// Make a name
		name, _ := cmd.Flags().GetString(`name`)
		if name == `` {
			name = util.GenerateName()
		}
		// TODO: check for uniqueness

		// Create PID file
		// workdir, _ := cmd.Flags().GetString(`workdir`)
		// pid := os.Getpid()
		// pidFilepath := path.Join(workdir, name)
		// err = ioutil.WriteFile(pidFilepath, []byte(strconv.Itoa(pid)), 0644)
		// if err != nil {
		// 	log.Error(err)
		// 	return
		// }
		// defer os.Remove(pidFilepath)

		// Assign a free TCP port for agent communication
		port, _ := cmd.Flags().GetInt(`port`)
		if port == 0 {
			port = util.FindFreeTCPPort()
		}

		publicAddress, _ := cmd.Flags().GetString(`public`)
		centralAddress, _ := cmd.Flags().GetString(`central`)

		agentConfig := map[string]interface{}{
			`name`: name,
			//`workdir`:       workdir,
			`port`:          port,
			`publicAddress`: publicAddress,
			`central`: map[string]interface{}{
				`address`: centralAddress,
			},
		}
		config, _ = config.Fill(agentConfig, `_agent`)

		log.WithFields(log.Fields{`name`: name}).Info("Agent rezzing.")
		agent, err := agent.New(config)
		if err != nil {
			log.Error(err)
			return
		}

		// Capture CTRL-C
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			log.Info(`Goodbye. --` + name)
			//os.Remove(pidFilepath)
			os.Exit(-1)
		}()

		agent.Spin()

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
	rootCmd.PersistentFlags().StringP("workdir", "w", "./.agent/", "Path to working directory.")
	rootCmd.Flags().StringP("name", "n", "", "Name of the agent. Must be unique. If not specified, one will be created.")
	rootCmd.Flags().StringP("config", "c", "", "Path to agent configuration.")
	rootCmd.Flags().String("central", "", "Central address.")
	rootCmd.Flags().String("public", "", "Public address to advertise to other agents.")
	rootCmd.Flags().IntP("port", "p", 0, "Port to connect agent HTTP-JSON API to.")
	rootCmd.MarkFlagRequired(`config`)
}
