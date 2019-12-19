package cmd

import (
	"github.com/siddhant94/gloin/cmd/linux"
	"github.com/spf13/cobra"
)

// Settings has methods with return type []string as exec.Command takes, cmd name and ...arg string (string, []string)
//type Settings interface {
//	GetCmd() (string, []string)             // gsettings get SCHEMA [:PATH] KEY
//	SetCmd(value string) (string, []string) // gsettings set SCHEMA [:PATH] KEY VALUE
//}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}

var scheduleCmd = &cobra.Command{
	//TODO: add option for path argument to act as just setting wallpaper
	Use:   "set",
	Short: "Sets background to Bing wallpaper of the day",
	Long: `It downloads the latest Bing image. saves it to user specified directory and assigns as new background 
			depending on OS`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check OS TODO: Remove stringg literal and add support for windows
		if state.Os == "linux" {
			linux.SetWallpaper(state.ImageDir, state.LatestImage)
		}
		//linux.AddCronEntry()

	},
}
