package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	b, err := os.ReadFile("cmds.txt")
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	cmds := strings.Split(string(b), "\n")
	for _, cmd := range cmds {
		if cmd == "" {
			continue
		}
		wg.Add(1)
		go func(cmd string, wg *sync.WaitGroup) {
			command := exec.Command("bash", "-c", cmd)
			err = command.Run()
			if err != nil {
				log.Println("Failed to run ", cmd)
			}
			log.Println(cmd, " ... done")
			wg.Done()
		}(cmd, &wg)
	}
	wg.Wait()
}
