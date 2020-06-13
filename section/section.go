package section

import (
	"os/exec"
	"strings"
)

// ParseAndRun ... Parse the commands and run them
func ParseAndRun(commands []string, channel chan map[string]Status) {
	statuses := map[string]Status{}

	// Setting initial status of not started
	for _, command := range commands {
		statuses[command] = NotStarted
	}
	channel <- statuses

	for _, command := range commands {
		statuses[command] = Started
		channel <- statuses

		// Parsing and running
		chunks := strings.Split(command, " ")
		var err error
		if len(chunks) == 0 {
			err = exec.Command(chunks[0]).Run()
		} else {
			err = exec.Command(chunks[0], chunks[0:]...).Run()
		}
		if err != nil {
			statuses[command] = Failed
			channel <- statuses
		}
		statuses[command] = Success
		channel <- statuses
	}

	close(channel)
}
