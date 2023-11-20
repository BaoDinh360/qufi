/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example/qufi/features/record"
	"fmt"

	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:     "open",
	Aliases: []string{"o"},
	Short:   "Open a record",
	Long:    `This command open a record file using your default text editor.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Begin opening record...")

		if len(args) <= 0 {
			fmt.Println("Error: Please type the name of the record")
			return
		}

		recordname := args[0]
		err := record.OpenRecord(recordname)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Println("Open command exited...")
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
