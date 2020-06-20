package section

import (
	"sync"
)

// Initialize ... Initialize a section
func Initialize(sectionName string, commands []string, statuses map[string]map[string]Status, mutex *sync.Mutex, sectionsRunning *int, wg *sync.WaitGroup) {
	// Setting initial as NotStarted
	mutex.Lock()
	commandsAdded := 0
	for _, command := range commands {
		if commandsAdded == 0 {
			statuses[sectionName] = make(map[string]Status)
		}
		statuses[sectionName][command] = NotStarted
		commandsAdded++
	}
	mutex.Unlock()

	*sectionsRunning--
	wg.Done()
}
