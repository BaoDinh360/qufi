package record

import (
	"errors"
	"example/qufi/features/config"
	"example/qufi/features/template"
	"example/qufi/helpers"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
)

type RecordInfo struct {
	RecordName string
	DirName    string
	DirPath    string
	FullPath   string
	Extension  string
}

// ==========RecordInfo methods=================
func (recordinfo *RecordInfo) SetRecordName(recordname string) {
	recordinfo.RecordName = recordname
	newrecordpath := filepath.Join(recordinfo.DirPath, fmt.Sprintf("%s.%s", recordinfo.RecordName, recordinfo.Extension))
	recordinfo.FullPath = newrecordpath
}

//=============================================

func GenerateRecordInfo(recordname string, dirname string, extension string) RecordInfo {
	basecfg := config.GetBaseConfig()

	recordfull := fmt.Sprintf("%s.%s", recordname, extension)
	dirpath := filepath.Join(basecfg.StoragePath, dirname)
	recordpath := filepath.Join(dirpath, recordfull)

	return RecordInfo{
		RecordName: recordname,
		DirName:    dirname,
		DirPath:    dirpath,
		FullPath:   recordpath,
		Extension:  extension,
	}
}

func CreateRecord(recordinfo RecordInfo) error {
	isdirexists, _ := helpers.IsPathExists(recordinfo.DirPath)
	//if directory not exists, create it
	if !isdirexists {
		_, err := helpers.CreateDirectory(recordinfo.DirPath)
		if err != nil {
			return err
		}
	}

	//get record file path
	recordfile, err := os.Create(recordinfo.FullPath)
	if err != nil {
		return err
	}
	defer recordfile.Close()

	//write to created file with template if template exist
	err = template.WriteTemplateIfExists(recordinfo.DirName, recordfile)
	if err != nil {
		return err
	}

	return nil

}

// check if this record's extension is config to open in text editor after creation. If yes, open it
func OpenRecordAfterCreate(recordpath string, extension string) error {
	basecfg := config.GetBaseConfig()
	hasvalue := slices.Contains(basecfg.AutoOpenFileExtensions, extension)
	//if this record extension is not in config for open after created, return
	if !hasvalue {
		return nil
	}

	err := helpers.OpenFileInTextEditor(recordpath)
	if err != nil {
		return err
	}

	fmt.Println("Please wait for record to open in your text editor...")
	return nil
}

func OpenRecord(recordname string) error {

	cmd := exec.Command("cmd", "/c", "D:\\Learn\\Learning Golang\\qufi\\archive_test\\archive_1\\record_271023.txt")
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("start failed: %v", err)
	}
	fmt.Printf("Waiting for command to finish.\n")
	err = cmd.Wait()
	fmt.Printf("Command finished with error: %v\n", err)

	return nil
}

// func GetRecord() error {
// 	fullpath := "D:\\Learn\\Learning Golang\\qufi\\archive_test\\.qufi\\.record.json"
// 	return helpers.GetDataFromJSON(fullpath, &recordList)
// }

// func SaveRecord(record Record) error {
// 	if len(recordList) <= 0 {
// 		err := GetRecord()
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	//add new record to slice
// 	recordList = append(recordList, record)

// 	fullpath := "D:\\Learn\\Learning Golang\\qufi\\archive_test\\.qufi\\.record.json"
// 	return helpers.SaveDataToJSON(fullpath, recordList)
// }

func GenerateNewRecordName(recordname string, dirpath string, extension string) (string, error) {
	//loop 100 times
	for i := 1; i <= 100; i++ {
		//generate new record name as <recordname>_(<i>)
		newname := fmt.Sprintf("%s_(%d)", recordname, i)
		newpath := filepath.Join(dirpath, fmt.Sprintf("%s.%s", newname, extension))
		fmt.Printf("Newpath: %v\n", newpath)
		isexists, _ := helpers.IsPathExists(newpath)
		//if new record name is not exists in this directory, return it, else keep looping
		if !isexists {
			return newname, nil
		}
	}

	return "", errors.New("cannot generate new record name")
}
