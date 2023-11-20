/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example/qufi/features/config"
	"example/qufi/features/list"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l", "lst"},
	Short:   "List all child directories and files in a directory inside your storage location",
	Long: `This command list all child directories and files in a directory.
- If directory not specified, list all child directories and files in your storage location.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Begin list command...")
		defer fmt.Println("List command exited...")

		config.CheckRequiredConfig()

		var dirname string
		//check args
		if len(args) <= 0 {
			//dirname will be the storage location
			dirname = ""
		} else {
			dirname = strings.Join(args, " ")
		}

		err := list.DisplayDirectoryContents(dirname)
		if err != nil {
			fmt.Printf("List contents in directory error: %v\n", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
