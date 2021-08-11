package main

import (
	"log"

	"github.com/zhaoxfan98/Tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}
}
