package main

import (
	"github.com/Matt-Gleich/Simultaneous-Updates/config"
	"github.com/Matt-Gleich/Simultaneous-Updates/section"
	"github.com/buger/goterm"
)

func main() {
	configContents := config.Extract(config.Path())
	var sectionStatuses map[string]chan map[string]section.Status

	for sectionName, commands := range configContents {
		go section.ParseAndRun(commands, sectionStatuses[sectionName])
	}

	for {
		// Checking to see if they all finished
		var allDone bool
		for _, section := range sectionStatuses {
			_, open := <-section
			allDone = !open
		}
		if allDone {
			break
		}

		goterm.Clear()

	}
}
