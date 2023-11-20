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

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Set up directory path for notes and records storage.",
	Long: `This command set up directory path for your notes and records storage.
The directory path must be absolute path.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Begin storage config command...")
		defer fmt.Println("Storage config command exited...")

		//convert path
		storagepath := helpers.ConvertDirectoryPath(args[0])

		//directory path validation
		isvalid, err := helpers.CheckValidPath(storagepath)
		if err != nil {
			fmt.Printf("Directory path validation error: %v\n", err)
			return
		}
		if !isvalid {
			fmt.Println("This is not a valid directory path")
			return
		}

		err = config.EditConfig(cmd.Name(), storagepath, "")
		if err != nil {
			fmt.Printf("Edit config errors: %v\n", err)
			return
		}

		fmt.Println("Config for storage path has been updated")
	},
}

func init() {
	configCmd.AddCommand(storageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
