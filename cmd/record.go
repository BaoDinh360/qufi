/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"example/qufi/features/config"
	"example/qufi/features/record"
	"example/qufi/helpers"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// recordCmd represents the record command
// subcommand of create command
var recordCmd = &cobra.Command{
	Use:     "record",
	Aliases: []string{"r", "rec"},
	Short:   "(r, rec) Create a record file",
	Long: `This command create a record in a directory. Use this command with create command prefix.
- If record name is not specified, generate default name for this record.
- Use --directory/-d flag to specify directory for this record. Create a new directory if the specified directory not exists. Default directory is <YYYY>-<MM>-<DD>_temp.
- Use --extension/-e flag to specify this record extension. Default extension is markdown(.md).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Begin record command...")
		defer fmt.Println("Record command exited...")

		config.CheckRequiredConfig()

		var recordname string
		//check args
		if len(args) <= 0 {
			recordname = fmt.Sprintf("%s_temp", time.Now().Format("2006-01-02"))
			fmt.Printf("Record name not specified. Generate default record name : %s\n", recordname)
		} else {
			recordname = strings.Join(args, " ")
		}

		var dirname string
		//check for default flag
		dirflag, _ := cmd.Flags().GetString("directory")
		if dirflag == "" {
			dirname = fmt.Sprintf("%s_temp", time.Now().Format("2006-01-02"))
			fmt.Printf("Directory location not specified. This record will be in directory: %s\n", dirname)
		} else {
			dirname = dirflag
		}

		//get record extension
		recordext, _ := cmd.Flags().GetString("extension")

		//generate record info
		recordinfo := record.GenerateRecordInfo(recordname, dirname, recordext)
		recordexists, _ := helpers.IsPathExists(recordinfo.FullPath)
		if recordexists {
			newrecordname, err := record.GenerateNewRecordName(recordname, recordinfo.DirPath, recordinfo.Extension)
			if err != nil {
				fmt.Printf("Generate new record name error: %v\n", err)
				return
			}
			fmt.Printf("There is already a record name %s in %s. Renaming new record to %s\n", recordname, recordinfo.DirPath, newrecordname)
			recordinfo.SetRecordName(newrecordname)
		}

		//create record
		err := record.CreateRecord(recordinfo)
		if err != nil {
			fmt.Printf("Record command error: %v\n", err)
			return
		}
		fmt.Printf("Record: %s is created at: %s\n", recordinfo.RecordName, recordinfo.FullPath)

		// check if this record's extension is config to open in text editor after creation. If yes, open it
		err = record.OpenRecordAfterCreate(recordinfo.FullPath, recordinfo.Extension)
		if err != nil {
			fmt.Printf("Open record in text editor error: %v\n", err)
			return
		}

	},
}

func init() {
	createCmd.AddCommand(recordCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recordCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recordCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// recordCmd.Flags().StringP("archivename", "a", "", "Specify archive for record. If not used, create an archive with random name")
	recordCmd.Flags().StringP("extension", "e", "md", "Specify record extension")
	recordCmd.Flags().StringP("directory", "d", "", "Specify the directory for this record.")
}
