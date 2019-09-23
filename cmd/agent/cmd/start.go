package cmd

import (
	"agentkit/pkg/agentkit"
	"agentkit/pkg/agentkit/actuators"
	kactuators "agentkit/pkg/agentkit/actuators"
	"agentkit/pkg/agentkit/minds"
	"agentkit/pkg/agentkit/queues"
	ksensors "agentkit/pkg/agentkit/sensors"
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

		fmt.Println("Agent rezzing.")

		// Queues
		percepts := queues.NewInMemoryPerceptQueue()
		actions := queues.NewInMemoryActionQueue()

		// Sensors
		var sensorConfigs []*ksensors.SensorConfig
		err = config.Lookup(`sensors`).Decode(&sensorConfigs)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		sensors := []ksensors.Sensor{}
		for _, sensorConfig := range sensorConfigs {
			sensor := ksensors.New(sensorConfig, percepts)
			sensors = append(sensors, sensor)
		}

		// Actuators
		var actuatorConfigs []*actuators.ActuatorConfig
		err = config.Lookup(`actuators`).Decode(&actuatorConfigs)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		actuators := []actuators.Actuator{}
		for _, actuatorConfig := range actuatorConfigs {
			actuator := kactuators.New(actuatorConfig, actions)
			actuators = append(actuators, actuator)
		}

		// ActionDispatch
		actionDispatch := agentkit.NewActionDispatch(actions)
		actionDispatch.RegisterAll(actuators)

		// Mind
		var mindConfig *minds.Config
		err = config.Lookup(`mind`).Decode(&mindConfig)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		mind := minds.New(mindConfig, percepts, actions)

		agent := &agentkit.Agent{
			Sensors:        sensors,
			Actuators:      actuators,
			Mind:           mind,
			ActionDispatch: actionDispatch,
		}

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

	startCmd.Flags().String("name", "n", "Name of the agent. Must be unique. If not specified, one will be created.")
	startCmd.Flags().BoolP("register", "r", true, "Register to `agentd` or execute now.")
	startCmd.Flags().String("config", "c", "Path to agent configuration.")
	startCmd.MarkFlagRequired(`config`)
}
