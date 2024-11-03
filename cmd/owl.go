package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gkwa/alpinerider/core"
	"github.com/spf13/cobra"
)

var owlCmd = &cobra.Command{
	Use:   "owl",
	Short: "Generate Go config with tests",
	Long:  `Generate renovate config for Go projects with tests`,
	Run: func(cmd *cobra.Command, args []string) {
		composer := core.NewConfigComposer(core.BaseConfig)
		config, err := composer.Owl()
		if err != nil {
			fmt.Printf("Error creating Go with tests config: %v\n", err)
			return
		}
		json, _ := json.MarshalIndent(config, "", "  ")
		fmt.Println(string(json))
	},
}

func init() {
	rootCmd.AddCommand(owlCmd)
}
