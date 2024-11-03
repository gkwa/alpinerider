package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gkwa/alpinerider/core"
	"github.com/spf13/cobra"
)

var monkeyCmd = &cobra.Command{
	Use:   "monkey",
	Short: "Generate hybrid config",
	Long:  `Generate hybrid renovate config with Node.js and Go settings`,
	Run: func(cmd *cobra.Command, args []string) {
		composer := core.NewConfigComposer(core.BaseConfig)
		config, err := composer.Monkey()
		if err != nil {
			fmt.Printf("Error creating monkey config: %v\n", err)
			return
		}
		json, _ := json.MarshalIndent(config, "", "  ")
		fmt.Println(string(json))
	},
}

func init() {
	rootCmd.AddCommand(monkeyCmd)
}
