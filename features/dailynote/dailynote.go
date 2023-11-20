package dailynote

import (
	"example/qufi/features/config"
	"example/qufi/features/template"
	"example/qufi/helpers"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type DailyNoteFileInfo struct {
	FileName string
	DirPath  string
	FullPath string
}

func GenerateDailyNoteFileInfo() DailyNoteFileInfo {
	basecfg := config.GetBaseConfig()
	dtfmt := "2006-01-02"
	tnow := time.Now().Format(dtfmt)

	dnotename := fmt.Sprintf("%s.md", tnow)
	dnotedirpath := filepath.Join(basecfg.StoragePath, "daily note", tnow)
	dnotepath := filepath.Join(dnotedirpath, dnotename)

	return DailyNoteFileInfo{
		FileName: dnotename,
		DirPath:  dnotedirpath,
		FullPath: dnotepath,
	}
}

func CreateDailyNote(dnoteinfo DailyNoteFileInfo) (string, error) {
	// basecfg := helpers.GetBaseConfig()
	// dtfmt := "2006-01-02"
	// tnow := time.Now().Format(dtfmt)

	// dnotename := fmt.Sprintf("%s.md", tnow)
	// dnotedirpath := filepath.Join(basecfg.DailyNotePath, tnow)
	// dnotepath := filepath.Join(dnotedirpath, dnotename)

	//check if daily dir exists, if not create new
	direxists, _ := helpers.IsPathExists(dnoteinfo.DirPath)
	if !direxists {
		_, err := helpers.CreateDirectory(dnoteinfo.DirPath)
		if err != nil {
			return "", err
		}
	}

	//create daily note
	dnotefile, err := os.Create(dnoteinfo.FullPath)
	if err != nil {
		return "", err
	}
	defer dnotefile.Close()
	//write the newly created file with template
	err = template.WriteTemplateIfExists("daily note", dnotefile)
	if err != nil {
		return "", err
	}

	return dnoteinfo.FullPath, nil

}
