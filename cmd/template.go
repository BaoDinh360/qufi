/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example/qufi/features/config"
	"example/qufi/helpers"
	"fmt"

	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Set up template directory path for templates files used in notes and records.",
	Long: `This command set up directory path for templates files used in notes and records.
The template directory path must be absolute path.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Begin template directory config command...")
		defer fmt.Println("Template directory config command exited...")

		//convert path
		templatepath := helpers.ConvertDirectoryPath(args[0])

		//directory path validation
		isvalid, err := helpers.CheckValidPath(templatepath)
		if err != nil {
			fmt.Printf("Directory path validation error: %v\n", err)
			return
		}
		if !isvalid {
			fmt.Println("This is not a valid directory path")
			return
		}

		err = config.EditConfig(cmd.Name(), templatepath, "")
		if err != nil {
			fmt.Printf("Edit config errors: %v\n", err)
			return
		}

		fmt.Println("Config for template directory path has been updated")

	},
}

func init() {
	configCmd.AddCommand(templateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// templateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
