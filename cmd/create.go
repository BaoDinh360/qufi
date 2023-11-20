/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c", "cre"},
	Short:   "Create a file (dailynote, record)",
	Long: `This command create an archive or create a record.
- Use with command dailynote/d/dno to create a daily note for today.
- Use with command record/r/rec to create a record.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Error: Please using create command with either archive or record subcommand")

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// createCmd.Flags().StringP("archivename", "a", "", "Specify archive name")
}
