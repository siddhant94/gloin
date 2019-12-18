package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gloin",
	Short: "Gloin is a CLI tool for scheduling background wallpaper.",
	Long: `Gloin is a CLI utility tool which enable users to schedule background wallpapers change on a customizable 
			frequency. It downloads Bing Wallpaper of the day, saves it to local filesystem and changes it to the freshly
			downloaded wallpaper.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Show CLI Usage.
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
