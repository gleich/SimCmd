package runnerconfig

import (
	"fmt"
	"os"
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
