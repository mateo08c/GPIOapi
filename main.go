package main

import (
	"GPIOapi/cmd"
	"log"
)

func main() {
	err := cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}
}
