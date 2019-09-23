package cmd

import (
	"agentkit/pkg/agentkit"
	"agentkit/pkg/agentkit/actuators"
	"agentkit/pkg/agentkit/sensors"
	"fmt"
	"io/ioutil"
	"os"

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
			os.Exit(-1)
		}

		// Compile configuration
		var r cue.Runtime
		config, err := r.Compile("agent", configData)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		fmt.Println(config)

		fmt.Println("MCP initializing.")

		percepts := agentkit.NewInMemoryPerceptQueue()
		actions := agentkit.NewInMemoryActionQueue()

		sensors := []agentkit.Sensor{
			&sensors.WebAPI{
				URL:  `https://api.openweathermap.org/data/2.5/weather?zip=98177,us&units=imperial&APPID=11c411febfa2057a80a18d89ff570383`,
				Rate: 0.1,
				Out:  percepts,
			},
		}

		actuators := []agentkit.Actuator{
			&actuators.StdOut{
				Label: `echo`,
				In:    actions,
			},
		}

		mind := &agentkit.LoopbackMind{
			Percepts: percepts,
			Actions:  actions,
		}

		actionDispatch := agentkit.NewActionDispatch(actions)
		actionDispatch.RegisterAll(actuators)

		agent := &agentkit.Agent{
			Sensors:        sensors,
			Actuators:      actuators,
			Mind:           mind,
			ActionDispatch: actionDispatch,
		}

		agent.Spin()

		fmt.Println("MCP Ready.")

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

	startCmd.Flags().String("name", "n", "Name of the agent. Must be unique. If not specified, one will be created.")
	startCmd.Flags().BoolP("register", "r", true, "Register to `agentd` or execute now.")
	startCmd.Flags().String("config", "c", "Path to agent configuration.")
	startCmd.MarkFlagRequired(`config`)
}
