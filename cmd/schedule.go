package cmd

import (
	"github.com/siddhant94/gloin/cmd/linux"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	rootCmd.AddCommand(scheduleCmd)
}

var scheduleCmd = &cobra.Command{
	//TODO: add option for path argument to act as just setting wallpaper
	Use:   "schedule",
	Short: "Schedules background change to user configured time.",
	Long: `It schedules a job which updates background on user configured settings, if some error happened in the 
			aforementioned time, default 09:00 hours is used`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check OS TODO: Remove stringg literal and add support for windows
		if state.Os == "linux" {
			// Get time from state
			hr := state.Scheduler.ChangeTime.Hour()
			min := state.Scheduler.ChangeTime.Minute()
			freq := getCronFreq(hr, min)
			linux.AddCronEntry(freq)
		}

	},
}

// Cron Frequency : *  *  *  *  * (min hr day(month) month day(week))
// *	any value
// ,	value list separator
// -	range of values
// /	step values
func getCronFreq(hr, min int) string {
	return strconv.Itoa(min) + " " + strconv.Itoa(hr) + " * * *"
}
