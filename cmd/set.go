package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
	"strings"
)

const backgroundSchema = "org.gnome.desktop.background"
const baseCommand = "gsettings"
const key = "picture-uri"
const prepend = "file://"

// Settings has methods with return type []string as exec.Command takes, cmd name and ...arg string (string, []string)
type Settings interface {
	getCmd() (string, []string)             // gsettings get SCHEMA [:PATH] KEY
	setCmd(value string) (string, []string) // gsettings set SCHEMA [:PATH] KEY VALUE
}
type Gsettings struct {
	Base     string
	BgSchema string
	Key      string
}

func (gsettings Gsettings) getCmd() (string, []string) {
	// return gsettings.Base + " get " + gsettings.BgSchema + " " + gsettings.Key
	return gsettings.Base, []string{"get", gsettings.BgSchema, gsettings.Key}
}

func (gsettings Gsettings) setCmd(value string) (string, []string) {
	// return gsettings.Base + " set " + gsettings.BgSchema + " " + gsettings.Key + " " + prepend + value
	return gsettings.Base, []string{"set", gsettings.BgSchema, gsettings.Key, (prepend+value)}
}

var gSettings Settings

func init() {
	rootCmd.AddCommand(scheduleCmd)
	gSettings = Gsettings {
		Base:     baseCommand,
		BgSchema: backgroundSchema,
		Key:      key,
	}
}

var scheduleCmd = &cobra.Command{
	//TODO: add option for path argument to act as just setting wallpaper
	Use:   "set",
	Short: "Sets background to Bing wallpaper of the day",
	Long: `It downloads the latest Bing image. saves it to user specified directory and assigns as new background 
			depending on OS`,
	Run: func(cmd *cobra.Command, args []string) {
		name, osCmd := gSettings.getCmd()
		op, err := exec.Command(name, osCmd...).CombinedOutput()
		if err != nil {
			fmt.Printf("Error getting current wallpaper : %v\n %v\n", err, string(op))
			return
		}
		currentWallpaper := string(op)
		// Remove `prepend` i.e. `file://`
		currentWallpaper = strings.Replace(currentWallpaper, prepend, "", 1)
		fileUri := state.ImageDir + "/" + state.LatestImage
		name, osCmd = gSettings.setCmd(fileUri)
		op, err = exec.Command(name, osCmd...).CombinedOutput()
		if err != nil {
			fmt.Printf("Error setting wallpaper (%v) : %v\n %v\n", fileUri, err, string(op))
			return
		}
		fmt.Println("Output :")
		fmt.Println(op)
		//linux.AddCronEntry()

	},
}
