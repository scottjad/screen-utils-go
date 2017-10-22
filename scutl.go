package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func setMsgwait(x int) {
	cmd := "screen"
	args := []string{"-X", "msgwait", strconv.Itoa(x)}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		// fmt.Printf("ERROR: Couldn't set msgwait: %d.\n", x)
	}
}

func setup() {
	setMsgwait(0)
}

func cleanup() {
	setMsgwait(5)
}

func getWindowTitle(windowIndex int) (string, error) {
	cmd := "screen"
	args := []string{"-p", strconv.Itoa(windowIndex), "-Q", "title"}
	// TODO Output should return a []byte, figure out how that gets converted to string, or if they're the same.
	output, err := exec.Command(cmd, args...).Output()
	title := string(output)
	if err != nil {
		title = ""
		// fmt.Printf("ERROR: Couldn't get window title for window: %d.\n", windowIndex)
	}
	return title, err

}

func GetWindows(windowLimit int, emptyLimit int) map[int]string {
	windows := make(map[int]string) // windowIndex -> title
	empty := 0
	for i := 0; i < windowLimit; i++ {
		title, err := getWindowTitle(i)
		if err == nil {
			windows[i] = title
			empty = 0
		} else if empty >= emptyLimit {
			break
		} else {
			empty++
		}
		// fmt.Printf("Window: %v Title: %v\n", i, title)
	}
	return windows
}

func windows() {

	setup()
	windows := GetWindows(windowLimit, emptyLimit)
	for i, title := range windows {
		fmt.Println(i, title)
	}
	cleanup()
}

func renumber(oldIndex int, newIndex int) {
	cmd := "screen"
	args := []string{"-p", strconv.Itoa(oldIndex), "-Q", "number", strconv.Itoa(newIndex)}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Printf("ERROR: Couldn't renumber window %v to %v.\n", oldIndex, newIndex)
	}
}

func repack() {
	windows := GetWindows(windowLimit, emptyLimit)
	newIndex := 1

	var keys []int
	for k := range windows {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, oldIndex := range keys {
		if newIndex != oldIndex {
			renumber(oldIndex, newIndex)
		}
		newIndex++

	}
}

func killWindow(index int) {
	exec.Command("screen", "-p", strconv.Itoa(index), "-X", "kill").Run()
}

func currentWindow() int {
	output, err := exec.Command("screen", "-Q", "number").Output()
	if err != nil {
		log.Fatal(err)
	}
	index, err := strconv.Atoi(strings.Split(string(output), " ")[0])
	return index
}

func selectWindow(index int) {
	exec.Command("screen", "-Q", "select", strconv.Itoa(index)).Run()
}

func killshells() {
	const titlesToKill = "^(zsh|bash|pg|htop|nmtui)"

	initialWindow := currentWindow()

	windows := GetWindows(windowLimit, emptyLimit)
	for index, title := range windows {

		match, err := regexp.MatchString(titlesToKill, title)
		if err != nil {
			log.Fatal(err)
		}

		if index != initialWindow && match {
			killWindow(index)
		}

	}
	exec.Command("screen-repack-go").Run()

}

const windowLimit = 50
const emptyLimit = 10

func main() {

	killshells()
	repack()
}
