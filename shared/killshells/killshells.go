package killshells

import (
	"regexp"

	"../repack"
	"../screen"
	"../windows"
)

func Killshells() error {
	const titlesToKill = "^\\s*(zsh|bash|pg|htop|nmtui)"

	selectedWindow, err := screen.CurrentWindowNumber()
	if err != nil {
		return err
	}
	for i, title := range windows.GetWindows() {

		match, err := regexp.MatchString(titlesToKill, title)
		if err != nil {
			return err
		}

		if i != selectedWindow && match {
			err := screen.KillWindow(i)
			if err != nil {
				return err
			}
		}

	}
	repack.Repack()
	return nil
}
