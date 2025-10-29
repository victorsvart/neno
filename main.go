package main

import (
	"log"

	"github.com/victorsvart/neno/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
