package main

import (
	"log"

	"github.com/spf13/cobra/cobra/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}
}
