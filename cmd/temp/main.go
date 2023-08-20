package main

import (
	"log"

	cputemp "github.com/jig/raspi-cputemp"
)

func main() {
	temp, err := cputemp.Get()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CPU temp is %.3f°C", temp)
}
