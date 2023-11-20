package helpers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"
)

// type Config struct {
// 	StoragePath            string
// 	TemplatePath           string
// 	AutoOpenFileExtensions []string
// }

// func GetBaseConfig() Config {

// 	cfgpath, err := config.GetConfigPath("config")
// 	if err != nil {
// 		panic(err)
// 	}

// 	//read JSON file
// 	content, err := os.ReadFile(cfgpath)
// 	if err != nil {
// 		panic(err)
// 	}

// 	var basecfg Config

// 	err = json.Unmarshal(content, &basecfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return basecfg
// }

func GetEntryTime(entryinfo os.FileInfo) (ctime time.Time) {
	//if current OS == "windows"
	//IMPLEMENT OS CHECK LATER
	stat := entryinfo.Sys().(*syscall.Win32FileAttributeData)
	ctime = time.Unix(0, stat.CreationTime.Nanoseconds())
	return
}

func CreateDirectory(dirpath string) (string, error) {
	err := os.MkdirAll(dirpath, 0750)
	if err != nil {
		return "", err
	}

	return dirpath, err
}

func IsPathExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			//dirname not exists
			return false, nil
		} else {
			return true, nil
		}
	}
	return true, nil
}

func GetDirectoryInfo(dirpath string) (os.FileInfo, error) {
	return os.Stat(dirpath)
}

func ListDirectoryEntries(dirpath string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(dirpath)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func OpenFileInTextEditor(filefullpath string) error {
	cmd := exec.Command("cmd", "/c", filefullpath)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("open file in text editor failed: %v", err)
	}
	return nil
}

func GetDataFromJSON(fullpath string, result interface{}) error {
	//read JSON file
	content, err := os.ReadFile(fullpath)
	if err != nil {
		return err
	}

	//only unmarshal file if it has content
	if len(content) > 0 {
		//convert JSON content to struct
		err = json.Unmarshal(content, result)
		if err != nil {
			return err
		}
	}

	return nil
}

func SaveDataToJSON(fullpath string, data interface{}) error {
	//convert struct to JSON content
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	//write to JSON file
	err = os.WriteFile(fullpath, content, 0666)
	if err != nil {
		return err
	}
	return nil
}

func GenerateRandomInt(min int, max int) int {
	return rand.Intn(max - min + 1)
}

// func GetRandomAffirmation() (string, error) {
// 	var affirmations []string

// 	err := GetDataFromJSON("./config/affirmations.json", &affirmations)
// 	if err != nil {
// 		return "", err
// 	}

// 	randint := generateRandomInt(0, len(affirmations)-1)

// 	return affirmations[randint], nil
// }

func ConvertDirectoryPath(path string) string {
	newpath := strings.TrimSpace(path)
	//replace WINDOWS \\ with /
	newpath = strings.ReplaceAll(newpath, "\\", "/")
	fmt.Println(newpath)
	return newpath
}

func CheckValidPath(path string) (bool, error) {
	dirpathregex := "^([a-zA-Z]:)*(/[a-zA-Z0-9]+([\\s_-]*[a-zA-Z0-9]*)*/*)+$"

	isvalid, err := regexp.MatchString(dirpathregex, path)
	if err != nil {
		return false, err
	}
	return isvalid, nil
}

func IndexOf[T comparable](el T, slice []T) int {
	for i, v := range slice {
		if el == v {
			return i
		}
	}
	return -1
}

func RemoveFromSlice[T comparable](slice []T, index int) []T {
	var result []T

	//if index is the last element in slice
	if index == len(slice)-1 {
		result = slice[:index]
	} else {
		result = append(slice[:index], slice[index+1])
	}

	return result
}
