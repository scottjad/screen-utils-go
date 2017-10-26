package main

import (
	"../../shared/killshells"
	"../../shared/screen"
)

func main() {
	screen.Setup()
	killshells.Killshells()
	screen.Cleanup()
}
