package config

import (
	"encoding/json"
	"errors"
	"example/qufi/helpers"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type Config struct {
	StoragePath            string
	TemplatePath           string
	AutoOpenFileExtensions []string
}

func GetBaseConfig() Config {

	cfgpath, err := GetConfigPath("config")
	if err != nil {
		panic(err)
	}

	//read JSON file
	content, err := os.ReadFile(cfgpath)
	if err != nil {
		panic(err)
	}

	var basecfg Config

	err = json.Unmarshal(content, &basecfg)
	if err != nil {
		panic(err)
	}

	return basecfg
}

func CheckRequiredConfig() {
	basecfg := GetBaseConfig()
	if basecfg.StoragePath == "" {
		fmt.Println(`Please set up the storage path config before using this command.
- Use the following command to set up config: qufi config storage "<your_directory_path>"`)
		os.Exit(1)
	}
}

func SetUpConfigFiles() {
	userhome, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	configdir := filepath.Join(userhome, ".qufi")
	configfiles := []struct{ Cfgpath, Cfgtype string }{
		{
			Cfgpath: filepath.Join(configdir, "config.json"),
			Cfgtype: "config",
		},
		{
			Cfgpath: filepath.Join(configdir, "affirmations.json"),
			Cfgtype: "affirmations",
		},
	}

	isconfigexists, _ := helpers.IsPathExists(configdir)
	// fmt.Printf("Is .qufi config exists? %v\n", isconfigexists)
	//if config dir not exists, create dir and child config files
	if !isconfigexists {
		_, err := helpers.CreateDirectory(configdir)
		if err != nil {
			panic(err)
		}

		for _, value := range configfiles {
			//create config files
			cfile, err := os.Create(value.Cfgpath)
			if err != nil {
				panic(err)
			}

			//write default config files
			err = writeDefaultConfig(value.Cfgpath, value.Cfgtype)
			if err != nil {
				panic(err)
			}

			cfile.Close()
		}
	}

}

func GetConfigPath(cfgname string) (string, error) {
	userhome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	cfgfile := fmt.Sprintf("%s.json", cfgname)

	path := filepath.Join(userhome, ".qufi", cfgfile)
	return path, nil
}

func ListAllConfigs() {
	basecfg := GetBaseConfig()

	value := reflect.ValueOf(basecfg)

	fmt.Println("Configurations for this CLI app:")
	fmt.Println("----------------------------------")
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("%s : %v\n", value.Type().Field(i).Name, value.Field(i))
	}
	fmt.Println("----------------------------------")
}

func EditConfig(configfield string, value interface{}, autoopenop string) error {

	basecfg := GetBaseConfig()

	switch configfield {
	case "storage":
		basecfg.StoragePath = value.(string)
	case "template":
		basecfg.TemplatePath = value.(string)
	case "autoopen":
		if autoopenop == "" {
			return errors.New("invalid auto open extensions operation")
		}

		newcfg, err := editAutoOpenExtensions(autoopenop, basecfg.AutoOpenFileExtensions, value.([]string))
		if err != nil {
			return err
		}

		basecfg.AutoOpenFileExtensions = newcfg

	default:
		return errors.New("invalid params configuration")
	}

	fmt.Println("Begin saving config...")
	//save config to JSON
	err := saveConfig(basecfg)
	if err != nil {
		return err
	}

	return nil
}

func GetRandomAffirmation() (string, error) {
	var affirmations []string

	affpath, err := GetConfigPath("affirmations")
	if err != nil {
		return "", err
	}

	err = helpers.GetDataFromJSON(affpath, &affirmations)
	if err != nil {
		return "", err
	}

	randint := helpers.GenerateRandomInt(0, len(affirmations)-1)

	return affirmations[randint], nil
}

func editAutoOpenExtensions(operation string, oldcfg []string, newcfg []string) ([]string, error) {
	fmt.Printf("%v, %v, %v\n", operation, oldcfg, newcfg)
	switch operation {
	case "add":
		for _, value := range newcfg {
			index := helpers.IndexOf[string](value, oldcfg)
			//if extensions is not in cfg, append
			if index == -1 {
				oldcfg = append(oldcfg, value)
			}
		}
	case "delete":
		for _, value := range newcfg {
			index := helpers.IndexOf[string](value, oldcfg)
			//if extensions is in cfg, remove
			if index != -1 {
				oldcfg = helpers.RemoveFromSlice[string](oldcfg, index)
			}
		}
	default:
		return nil, errors.New("invalid operation")
	}
	fmt.Println(oldcfg)
	return oldcfg, nil
}

func saveConfig(data interface{}) error {
	cfgpath, err := GetConfigPath("config")
	if err != nil {
		return err
	}
	return helpers.SaveDataToJSON(cfgpath, data)
}

func writeDefaultConfig(cfgpath string, cfgtype string) error {
	var data interface{}

	switch cfgtype {
	case "config":
		data = Config{
			StoragePath:            "",
			TemplatePath:           "",
			AutoOpenFileExtensions: []string{"md", "txt"},
		}
	case "affirmations":
		data = []string{
			"You got this",
			"You'll figure it out",
			"You're a smart cookie",
			"I believe in you",
			"Sucking at something is the first step towards being good at something",
			"Struggling is part of learning",
			"Everything has cracks - that's how the light gets in",
			"Mistakes don't make you less capable",
			"We are all works in progress",
			"You are a capable human",
			"You know more than you think",
			"10x engineers are a myth",
			"If everything was easy you'd be bored",
			"I admire you for taking this on",
			"You're resourceful and clever",
			"You'll find a way",
			"I know you'll sort it out",
			"Struggling means you're learning",
			"You're doing a great job",
			"It'll feel magical when it's working",
			"I'm rooting for you",
			"Your mind is full of brilliant ideas",
			"You make a difference in the world by simply existing in it",
			"You are learning valuable lessons from yourself every day",
			"You are worthy and deserving of respect",
			"You know more than you knew yesterday",
			"You're an inspiration",
			"Your life is already a miracle of chance waiting for you to shape its destiny",
			"Your life is about to be incredible",
			"Nothing is impossible. The word itself says 'I’m possible!'",
			"Failure is just another way to learn how to do something right",
			"I give myself permission to do what is right for me",
			"You can do it",
			"It is not a sprint, it is a marathon. One step at a time",
			"Success is the progressive realization of a worthy goal",
			"People with goals succeed because they know where they’re going",
			"All you need is the plan, the roadmap, and the courage to press on to your destination",
			"The opposite of courage in our society is not cowardice... it is conformity",
			"Whenever we’re afraid, it’s because we don’t know enough. If we understood enough, we would never be afraid",
			"The past does not equal the future",
			"The path to success is to take massive, determined action",
			"It’s what you practice in private that you will be rewarded for in public",
			"Small progress is still progress",
			"Don't worry if you find flaws in your past creations, it's because you've evolved",
			"Starting is the most difficult step - but you can do it",
			"Don't forget to enjoy the journey",
			"It's not a mistake, it's a learning opportunity",
		}
	default:
		return nil
	}

	return helpers.SaveDataToJSON(cfgpath, data)
}
