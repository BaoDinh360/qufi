/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example/qufi/features/config"
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"cfg"},
	Short:   "Set up configuration for this CLI app",
	Long: `This command perform configurations for this CLI app. Use this command with additional subcommand.
- Use with storage command to config directory path location for storing notes and records.
- Use with template command to config template directory path location for templates used in notes and records.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Begin config command...")
		defer fmt.Println("Config command exited...")

		config.ListAllConfigs()

	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
