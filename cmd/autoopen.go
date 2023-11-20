/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example/qufi/features/config"
	"example/qufi/features/prompt"
	"fmt"

	"github.com/spf13/cobra"
)

// autoopenCmd represents the autoopen command
var autoopenCmd = &cobra.Command{
	Use:   "autoopen",
	Short: "Set up which records extension will be opened in default text editor automatically after creation.",
	Long: `This command set up which records extension will be opened in default text editor automatically after creation. 
- In default, the extensions specified will be add to the config.
- If --delete/-d flag used, remove the extensions specified from the config.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Begin auto open config command...")
		defer fmt.Println("Auto open config command exited...")

		//check flags
		isdelete, _ := cmd.Flags().GetBool("delete")
		var message, autoopenop string
		if isdelete {
			autoopenop = "delete"
			fmt.Println("Remove record and notes extensions from being opened after its creation")
			message = "Select the extensions that will be removed"
		} else {
			autoopenop = "add"
			fmt.Println("Add record and notes extensions from being opened after its creation")
			message = "Select the extensions that will be added"
		}

		options := []string{"md", "txt", "json"}

		choices, err := prompt.MultiSelectPrompt(message, options)
		if err != nil {
			fmt.Printf("Select errors: %v\n", err)
		}

		err = config.EditConfig(cmd.Name(), choices, autoopenop)
		if err != nil {
			fmt.Printf("Edit config errors: %v\n", err)
			return
		}

		fmt.Println("Config for auto open extensions has been updated")

	},
}

func init() {
	configCmd.AddCommand(autoopenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// autoopenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// autoopenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	autoopenCmd.Flags().BoolP("delete", "d", false, "Delete the extensions from the config")
}
