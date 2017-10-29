package screen

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func SetMsgwait(x int) error {
	cmd := "screen"
	args := []string{"-X", "msgwait", strconv.Itoa(x)}
	return exec.Command(cmd, args...).Run()
}

func Setup() {
	SetMsgwait(0)
}

func Cleanup() {
	SetMsgwait(5)
}

func Renumber(oldIndex int, newIndex int) {
	cmd := "screen"
	args := []string{"-p", strconv.Itoa(oldIndex), "-Q", "number", strconv.Itoa(newIndex)}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Printf("ERROR: Couldn't renumber window %v to %v.\n", oldIndex, newIndex)
	}
}

func KillWindow(index int) error {
	return exec.Command("screen", "-p", strconv.Itoa(index), "-X", "kill").Run()
}

func CurrentWindowNumber() (int, error) {
	output, err := exec.Command("screen", "-Q", "number").Output()
	if err == nil {
		index, err := strconv.Atoi(strings.Split(string(output), " ")[0])
		return index, err
	}
	return 0, err

}

func SelectWindow(index int) {
	exec.Command("screen", "-Q", "select", strconv.Itoa(index)).Run()
}
