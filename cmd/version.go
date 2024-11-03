package cmd

import (
	"fmt"
	
	"github.com/spf13/cobra"

	"github.com/gkwa/alpinerider/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of alpinerider",
	Long:  `All software has versions. This is alpinerider's`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := version.GetBuildInfo()
		fmt.Println(buildInfo)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
