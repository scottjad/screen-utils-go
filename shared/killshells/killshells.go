package killshells

import (
	"log"
	"regexp"

	"../repack"
	"../screen"
	"../windows"
)

func Killshells() {
	const titlesToKill = "^\\s*(zsh|bash|pg|htop|nmtui)"

	initialWindow := screen.CurrentWindow()

	windows := windows.GetWindows()
	for index, title := range windows {

		match, err := regexp.MatchString(titlesToKill, title)
		if err != nil {
			log.Fatal(err)
		}

		if index != initialWindow && match {
			screen.KillWindow(index)
		}

	}
	repack.Repack()
}
