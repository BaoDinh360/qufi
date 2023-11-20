/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example/qufi/features/config"
	"example/qufi/features/dailynote"
	"example/qufi/features/prompt"
	"example/qufi/helpers"
	"fmt"

	"github.com/spf13/cobra"
)

// dailynoteCmd represents the dailynote command
var dailynoteCmd = &cobra.Command{
	Use:     "dailynote",
	Aliases: []string{"d", "dno"},
	Short:   "(d, dno) Create a daily note",
	Long: `This command create a daily note, with the default template. Use this command with create command prefix. 
- Daily note use markdown .md file.
- Default daily note structure: <your_storage>/daily note/<YYYY-MM-DD>/<YYYY-MM-DD>.md.
- Daily note will be opened in text editor after its creation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Begin generating daily note command...")
		defer fmt.Println("Generate daily note command exited...")

		config.CheckRequiredConfig()

		dnoteinfo := dailynote.GenerateDailyNoteFileInfo()
		dnoteexists, _ := helpers.IsPathExists(dnoteinfo.FullPath)
		//if daily note already existed, display menu selection asking user to overwrite note or cancel command
		if dnoteexists {
			message := "There is already a daily note with name in this folder. Do you want to overwrite it?"
			options := []string{"Overwrite the existing note", "Cancel command"}

			choice, err := prompt.SelectPrompt(message, options)
			if err != nil {
				fmt.Printf("Option select error: %v\n", err)
				return
			}

			switch choice {
			case 0:
				fmt.Println("Overwriting existing note...")
			case 1:
				fmt.Println("Canceling command...")
				return
			default:
				fmt.Println("Invalid option! Command exiting...")
				return
			}
		}

		dnotepath, err := dailynote.CreateDailyNote(dnoteinfo)
		if err != nil {
			fmt.Printf("Generate daily note error: %v\n", err)
			return
		}
		fmt.Printf("Daily note for today is created at path %s\n", dnotepath)

		err = helpers.OpenFileInTextEditor(dnotepath)
		if err != nil {
			fmt.Printf("Open daily note in text editor error: %v\n", err)
			return
		}
		fmt.Println("Please wait while daily note is being opened in text editor...")

	},
}

func init() {
	createCmd.AddCommand(dailynoteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dailynoteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dailynoteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
