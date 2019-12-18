package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Gloin",
	Long:  `All software has versions. This is Gloin's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gloin Desktop Background Scheduler v0.1 -- HEAD")
	},
}