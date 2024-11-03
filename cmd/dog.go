package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gkwa/alpinerider/core"
	"github.com/spf13/cobra"
)

var dogCmd = &cobra.Command{
	Use:   "dog",
	Short: "Generate Go config",
	Long:  `Generate renovate config for Go projects without tests`,
	Run: func(cmd *cobra.Command, args []string) {
		composer := core.NewConfigComposer(core.BaseConfig)
		config, err := composer.Dog()
		if err != nil {
			fmt.Printf("Error creating Go config: %v\n", err)
			return
		}
		json, _ := json.MarshalIndent(config, "", "  ")
		fmt.Println(string(json))
	},
}

func init() {
	rootCmd.AddCommand(dogCmd)
}
