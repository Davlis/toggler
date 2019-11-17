package toggler

import (
	"log"
	"os/exec"
	"sync"
)

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		log.Println("error occured")
		log.Printf("%s", err)
	}
	log.Printf("%s", out)

	wg.Done()
}

func Create() {
	Load()

	var command = "cd " + Cfg.ProjectDir + " && git standup -m 3"

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go exe_cmd(command, wg)

	wg.Wait()
}
