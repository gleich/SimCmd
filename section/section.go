package section

import (
	"os/exec"
	"strings"
	"sync"

	"github.com/buger/goterm"
)

// Initialize ... Initialize a section
func Initialize(sectionName string, commands []string, statuses map[string]map[string]Status, mutex *sync.Mutex, wg *sync.WaitGroup, sectionsRunning *int) {
	*sectionsRunning++
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
	wg.Done()
}

// Run ... Run a section
func Run(sectionName string, commands []string, statuses map[string]map[string]Status, mutex *sync.Mutex, sectionsRunning *int) {
	goterm.Println("Running section:" + sectionName)
	for _, command := range commands {
		chunks := strings.Split(command, " ")
		cmd := exec.Command(chunks[0], chunks[1:]...)
		mutex.Lock()
		statuses[sectionName][command] = Started
		mutex.Unlock()
		err := cmd.Run()
		if err != nil {
			mutex.Lock()
			statuses[sectionName][command] = Failed
			mutex.Unlock()
		} else {
			mutex.Lock()
			statuses[sectionName][command] = Success
			mutex.Unlock()
		}
	}
	*sectionsRunning--
}
