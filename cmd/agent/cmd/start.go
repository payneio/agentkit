package cmd

import (
	"agentkit/pkg/agentkit/agent"
	"agentkit/pkg/agentkit/util"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path"
	"strconv"
	"syscall"

	"cuelang.org/go/cue"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Read in configuration
		configPath, _ := cmd.Flags().GetString("config")
		fmt.Println(`Agent configuration: ` + configPath)
		configData, err := ioutil.ReadFile(configPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Compile configuration
		var r cue.Runtime
		config, err := r.Compile("agent", configData)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Make a name
		name, _ := cmd.Flags().GetString(`name`)
		if name == `` {
			name = agent.GenerateName()
		}
		config, _ = config.Fill(name, `_name`)
		// TODO: check for uniqueness

		// Create PID file
		workdir, _ := cmd.Flags().GetString(`workdir`)
		pid := os.Getpid()
		pidFilepath := path.Join(workdir, name)
		err = ioutil.WriteFile(pidFilepath, []byte(strconv.Itoa(pid)), 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer os.Remove(pidFilepath)
		config, _ = config.Fill(workdir, `_workdir`)

		// Assign a free TCP port for agent communication
		port := util.FindFreeTCPPort()
		config, _ = config.Fill(port, `_port`)
		fmt.Printf("Agent %s rezzing.\n", name)

		agent, err := agent.New(config)
		if err != nil {
			fmt.Print(err)
			return
		}

		// Capture CTRL-C
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			fmt.Println(`Goodbye. --` + name)
			os.Remove(pidFilepath)
			os.Exit(-1)
		}()

		agent.Spin()

	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	startCmd.Flags().StringP("name", "n", "", "Name of the agent. Must be unique. If not specified, one will be created.")
	startCmd.Flags().BoolP("register", "r", true, "Register to `agentd` or execute now.")
	startCmd.Flags().StringP("config", "c", "", "Path to agent configuration.")
	startCmd.MarkFlagRequired(`config`)
}
