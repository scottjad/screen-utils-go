package main

import (
	"fmt"

	"../../shared/screen"
	"../../shared/windows"
)

func main() {
	screen.Setup()
	windows := windows.GetWindows()
	// fmt.Println(windows[1])
	for i, title := range windows {
		fmt.Println(i, title)
	}
	screen.Cleanup()
}
