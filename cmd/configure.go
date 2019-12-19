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
	ImageDir    string       `json:"image_dir"` // Path to default pictures directory for storing and changing wallpapers.
	LatestImage string       `json:"latest_image"`
	Scheduler   scheduleMeta `json:"scheduler"`
	Os          string       `json:"os"`
}

type scheduleMeta struct {
	Scheduled  bool          `json:"scheduled"`
	Interval   time.Duration `json:"interval"`
	ChangeTime time.Time     `json:"change_time"`
}

var state State
var defaultTime, _ = time.Parse("15:04", "09:00")

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
	Use:   "configure",
	Short: "Configure user-settings for Gloin.",
	Long:  `Takes in desired values for Gloin's settings from user and stores in json file.'`,
	//Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: keep the questioin and state json key in map
		directoryPath := getUserInput("Enter Directory path (absolute) where Backgrounds would be downloaded. ")
		if !pathExists(directoryPath) { // If Path doesn't exists, create it.
			err := os.Mkdir(directoryPath, 0664)
			if err != nil {
				log.Printf("Error creating directory - %s, \n %v \n", directoryPath, err)
				return
			}
		}
		state.ImageDir = directoryPath
		schedule := getUserInput("Do you want to schedule it for  daily change? Press y for Yes | n for No")
		if schedule == "n" {
			fmt.Println("Generating Configuration")
		} else if schedule == "y" {
			fmt.Println("Scheduling Daily update. ")
			t := getUserInput("Please enter update timestamp in hh:mm (24 hr format)")
			parsedTime, err := time.Parse("15:04", t)
			if err != nil {
				fmt.Println("Error parsing time, setting to default time i.e. 09:00 hrs : %v \n", err)
				state.Scheduler.ChangeTime = defaultTime
			} else {
				state.Scheduler.ChangeTime = parsedTime // "0 8 * * *"
			}
		}
		writeState(state)
		return
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
