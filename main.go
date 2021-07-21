package main

import (
	"fmt"
	"sync"

	"github.com/gleich/Simultaneous-Updates/config"
	"github.com/gleich/Simultaneous-Updates/section"
	"github.com/buger/goterm"
)

func main() {
	var (
		mutex           sync.Mutex
		wg              sync.WaitGroup
		sectionsRunning int
	)

	config := config.Extract()
	sectionStatuses := map[string]map[string]section.Status{}

	// fmt.Println("Extracted config")

	// Initalizing the section statuses
	for sectionName, commands := range config {
		wg.Add(1)
		go section.Initialize(sectionName, commands, sectionStatuses, &mutex, &wg, &sectionsRunning)
	}
	wg.Wait()

	// fmt.Println("Initalized all sections")

	// Running the sections
	for sectionName, commands := range config {
		go section.Run(sectionName, commands, sectionStatuses, &mutex, &sectionsRunning)
	}

	// fmt.Println("Kicked off all sections")

	for sectionsRunning != 0 {
		goterm.Clear()
		goterm.MoveCursor(1, 1)
		goterm.Println("Starting cycle")
		for sectionName, statuses := range sectionStatuses {
			goterm.Println(statuses)
			goterm.Println(goterm.Color(fmt.Sprintf("%v ⋯ %v commands", sectionName, len(statuses)), goterm.GREEN))
			commandsOutput := 0
			for command, status := range statuses {
				prefix := " ├──"
				if commandsOutput != len(statuses) {
					prefix = " └──"
				}
				var indicator string
				var color int
				switch status {
				case section.NotStarted:
					indicator = " "
					color = goterm.WHITE
				case section.Started:
					indicator = "◌"
					color = goterm.YELLOW
				case section.Success:
					indicator = "✔"
					color = goterm.GREEN
				case section.Failed:
					indicator = "✗"
					color = goterm.RED
				}
				goterm.Println(prefix + goterm.Color(indicator, color) + " " + command)
			}
		}
	}
	goterm.Flush()
}
