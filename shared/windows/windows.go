package windows

import (
	"os/exec"
	"strconv"
)

var WindowLimit int = 50
var EmptyLimit int = 10

func GetWindowTitle(windowIndex int) (string, error) {
	cmd := "screen"
	args := []string{"-p", strconv.Itoa(windowIndex), "-Q", "title"}
	// TODO Output should return a []byte, figure out how that gets converted to string, or if they're the same.
	output, err := exec.Command(cmd, args...).Output()
	title := string(output)
	if err != nil {
		// fmt.Printf("ERROR: Couldn't get window title for window: %d.\n", windowIndex)
	}
	return title, err

}

func GetWindows() map[int]string {
	windows := make(map[int]string) // windowIndex -> title
	empty := 0
	for i := 0; i < WindowLimit; i++ {
		title, err := GetWindowTitle(i)
		if err == nil {
			windows[i] = title
			empty = 0
		} else if empty >= EmptyLimit {
			break
		} else {
			empty++
		}
		// fmt.Printf("Window: %v Title: %v\n", i, title)
	}
	return windows
}
