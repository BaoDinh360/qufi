package list

import (
	"example/qufi/features/config"
	"example/qufi/helpers"
	"fmt"
	"io/fs"
	"path/filepath"
)

type EntryInfo struct {
	EntryType string
	Name      string
	Size      int64
	CreateAt  string
}

type DirectoryInfo struct {
	Name        string
	Path        string
	Size        int64
	CreateAt    string
	EntriesInfo []EntryInfo
}

func DisplayDirectoryContents(dirname string) error {
	basecfg := config.GetBaseConfig()

	//get dirpath
	var dirpath string
	//if dirname not specified, list contents from storage location directory
	if dirname == "" {
		dirpath = basecfg.StoragePath
	} else {
		dirpath = filepath.Join(basecfg.StoragePath, dirname)
	}

	//get DirectoryInfo
	directoryinfo, err := getDirectoryInfo(dirpath)
	if err != nil {
		return err
	}

	fmt.Printf("Directory of %s :\n", directoryinfo.Path)
	fmt.Printf("Directory created at: %s\n", directoryinfo.CreateAt)
	fmt.Printf("Total directory size: %d bytes\n", directoryinfo.Size)
	fmt.Printf("%-16s	%s	%-8s	%-8s\n", "Create date", "Type", "Size", "Name")
	fmt.Println("------------------------------------------------------------")
	for _, entry := range directoryinfo.EntriesInfo {
		printEntryInfo(entry)
	}
	fmt.Println("------------------------------------------------------------")
	return nil
}

func getDirectoryInfo(dirpath string) (DirectoryInfo, error) {
	dirinfo, err := helpers.GetDirectoryInfo(dirpath)
	if err != nil {
		return DirectoryInfo{}, nil
	}

	entries, err := helpers.ListDirectoryEntries(dirpath)
	if err != nil {
		return DirectoryInfo{}, nil
	}

	var entriesinfo []EntryInfo
	for _, entry := range entries {
		entryinfo, err := getEntryInfo(entry)
		if err != nil {
			return DirectoryInfo{}, nil
		}
		entriesinfo = append(entriesinfo, entryinfo)
	}

	ctime := helpers.GetEntryTime(dirinfo)

	return DirectoryInfo{
		Name:        dirinfo.Name(),
		Path:        dirpath,
		Size:        dirinfo.Size(),
		CreateAt:    ctime.Format("02/01/2006 03:04:05PM"),
		EntriesInfo: entriesinfo,
	}, nil
}

func getEntryInfo(entry fs.DirEntry) (EntryInfo, error) {

	entryinfo, err := entry.Info()
	if err != nil {
		return EntryInfo{}, err
	}

	var entrytype string
	if entry.IsDir() {
		entrytype = "DIR"
	} else {
		entrytype = "FILE"
	}

	name := entry.Name()
	size := entryinfo.Size()

	ctime := helpers.GetEntryTime(entryinfo)

	return EntryInfo{
		EntryType: entrytype,
		Name:      name,
		Size:      size,
		CreateAt:  ctime.Format("02/01/2006 03:04:05PM"),
	}, nil
}

func printEntryInfo(entryinfo EntryInfo) {
	fmt.Printf("%-9s	%s	%-9s	%-9s\n", entryinfo.CreateAt, entryinfo.EntryType, fmt.Sprintf("%d bytes", entryinfo.Size), entryinfo.Name)
}
