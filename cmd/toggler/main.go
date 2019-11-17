package main

import (
	"log"

	togglersvc "github.com/davlis/toggler/internal/toggler"
)

func main() {
	log.Println("cmd/toggler/main: Toggler started")

	togglersvc.Create()

	log.Println("cmd/toggler/main: Toggler finished")
}
