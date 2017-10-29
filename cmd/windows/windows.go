package main

import (
	"fmt"

	"../../shared/screen"
	"../../shared/windows"
)

func main() {
	screen.Setup()
	windows := windows.GetWindows()
	for i, title := range windows {
		fmt.Println(i, title)
	}
	screen.Cleanup()
}
