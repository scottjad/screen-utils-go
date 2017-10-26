package screen

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func SetMsgwait(x int) {
	cmd := "screen"
	args := []string{"-X", "msgwait", strconv.Itoa(x)}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		// fmt.Printf("ERROR: Couldn't set msgwait: %d.\n", x)
	}
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

func KillWindow(index int) {
	exec.Command("screen", "-p", strconv.Itoa(index), "-X", "kill").Run()
}

func CurrentWindow() int {
	output, err := exec.Command("screen", "-Q", "number").Output()
	if err != nil {
		log.Fatal(err)
	}
	index, err := strconv.Atoi(strings.Split(string(output), " ")[0])
	return index
}

func SelectWindow(index int) {
	exec.Command("screen", "-Q", "select", strconv.Itoa(index)).Run()
}
