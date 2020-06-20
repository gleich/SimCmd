package main

import (
	"fmt"
	"sync"

	"github.com/Matt-Gleich/Simultaneous-Updates/config"
	"github.com/Matt-Gleich/Simultaneous-Updates/section"
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
	for sectionName, commands := range config {
		wg.Add(1)
		sectionsRunning++
		go section.Initialize(sectionName, commands, sectionStatuses, &mutex, &sectionsRunning, &wg)
	}

	wg.Wait()

	goterm.Clear()
	goterm.MoveCursor(1, 1)
	for sectionName, statuses := range sectionStatuses {
		goterm.Println(statuses)
		goterm.Println(goterm.Color(fmt.Sprintf("%v ⋯ %v commands", sectionName, len(statuses)), goterm.GREEN))
	}
	goterm.Flush()
}
