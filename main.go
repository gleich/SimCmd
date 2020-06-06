package main

import (
	"fmt"
	"reflect"

	"github.com/Matt-Gleich/Simultaneous-Updates/runnerconfig"
)

func main() {
	path := runnerconfig.Path()
	config := runnerconfig.Extract(path)
	sections := reflect.ValueOf(config).MapKeys()
	for _, section := range sections {
		fmt.Println(section)
	}
}
