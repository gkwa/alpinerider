package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gkwa/alpinerider/core"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Generate Node.js config",
	Long:  `Generate renovate config for Node.js projects without tests`,
	Run: func(cmd *cobra.Command, args []string) {
		composer := core.NewConfigComposer(core.BaseConfig)
		config, err := composer.Cat()
		if err != nil {
			fmt.Printf("Error creating Node config: %v\n", err)
			return
		}
		json, _ := json.MarshalIndent(config, "", "  ")
		fmt.Println(string(json))
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
}
