package main

import (
	"log"

	"github.com/jitesoft/cc-gen/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Panic(err)
	}
}
