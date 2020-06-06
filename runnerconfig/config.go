package runnerconfig

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// Path ... Get the path of the runner config
func Path() string {
	var configPath string
	if len(os.Args) == 1 {
		path, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		configPath = path + "/"
	} else {
		configPath = os.Args[1]
	}
	return configPath
}

// Extract ... Get all the commands from the config
func Extract(path string) map[string][]string {
	file, err := ioutil.ReadFile(path)
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
