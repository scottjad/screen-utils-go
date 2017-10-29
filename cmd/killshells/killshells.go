package main

import (
	"log"

	"../../shared/killshells"
	"../../shared/screen"
)

func main() {
	screen.Setup()
	err := killshells.Killshells()
	if err != nil {
		log.Printf("Error killing shells: %v", err)
	}
	screen.Cleanup()
}
