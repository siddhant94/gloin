package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"time"
)

type State struct {
	ImageDir    string `json:"image_dir"`// Path to default pictures directory for storing and changing wallpapers.
	LatestImage string `json:"latest_image"`
	Scheduler   scheduleMeta `json:"scheduler"`
	Os          string `json:"os"`
}

type scheduleMeta struct {
	Scheduled            bool `json:"scheduled"`
	Interval             time.Duration `json:"interval"`
	LocalChangeTimestamp time.Time `json:"local_change_timestamp"`
}

var state State

const ConfigFilename = "state.json"

func init() {
	rootCmd.AddCommand(configureCmd)
	const goOS string = runtime.GOOS
	state = State{Os: goOS}
	// Get State if state json exists
	byteValue, err := ioutil.ReadFile(ConfigFilename)
	if err != nil {
		log.Printf("Error opening state file : %v \n", err)
		return
	}
	err = json.Unmarshal(byteValue, &state)
	if err != nil {
		log.Printf("Error unmarshalling state file: %v \n")
		return

	}
	fmt.Printf("STATE : \n %+v \n", state)
}

var configureCmd = &cobra.Command{
	Use:   "configure </path/to/directory>",
	Short: "Configure user-settings for Gloin.",
	Long:  `Takes in desired values for Gloin's settings from user and stores in json file.'`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		state.ImageDir = args[0] // /home/sid/Desktop/Workspace/bing-wallpapers
		// Check if directory exists, if not create one
		if _, err := os.Stat(state.ImageDir); os.IsNotExist(err) {
			fmt.Println("Directory path does not exists, creating directory")
			err = os.Mkdir(state.ImageDir, 0664)
			if err != nil {
				log.Printf("Error creating directory - %s, \n %v \n", state.ImageDir, err)
				return
			}
			fmt.Println("Successfully created directory - " + state.ImageDir)
		}
		// Write to a file for persistence
		writeState(state)
	},
}

func writeState(state State) error {
	jsonData, err := json.Marshal(state)
	if err != nil {
		log.Printf("Error JSON Marshall - %v \n", err)
		return err
	}
	err = ioutil.WriteFile(ConfigFilename, jsonData, 0644)
	if err != nil {
		log.Printf("Error Writing State JSON to file - %v \n", err)
		return err
	}
	fmt.Println("Successfully saved configuration.")
	return nil
}