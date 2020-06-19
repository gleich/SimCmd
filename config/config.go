package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// Extract ... Get all the sections and their corresponding commands from the config
func Extract() map[string][]string {
	var configPath string
	if len(os.Args) == 1 {
		path, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		configPath = path + "/updateRunner.yml"
	} else {
		configPath = os.Args[1]
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var yamlContents map[string][]string
	err = yaml.Unmarshal(file, &yamlContents)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return yamlContents
}
