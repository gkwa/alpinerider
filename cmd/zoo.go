package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/gkwa/alpinerider/core"
	"github.com/spf13/cobra"
)

var zooCmd = &cobra.Command{
	Use:   "zoo",
	Short: "Show all configs",
	Long:  `Show configs for all animal types (cat, dog, owl)`,
	Run: func(cmd *cobra.Command, args []string) {
		composer := core.NewConfigComposer(core.BaseConfig)

		cat, err := composer.Cat()
		if err != nil {
			fmt.Printf("Error creating cat config: %v\n", err)
			return
		}
		fmt.Println("-- cat --")
		catJson, _ := json.MarshalIndent(cat, "", "  ")
		fmt.Println(string(catJson))
		fmt.Println()

		dog, err := composer.Dog()
		if err != nil {
			fmt.Printf("Error creating dog config: %v\n", err)
			return
		}
		fmt.Println("-- dog --")
		dogJson, _ := json.MarshalIndent(dog, "", "  ")
		fmt.Println(string(dogJson))
		fmt.Println()

		owl, err := composer.Owl()
		if err != nil {
			fmt.Printf("Error creating owl config: %v\n", err)
			return
		}
		fmt.Println("-- owl --")
		owlJson, _ := json.MarshalIndent(owl, "", "  ")
		fmt.Println(string(owlJson))
	},
}

func init() {
	rootCmd.AddCommand(zooCmd)
}
