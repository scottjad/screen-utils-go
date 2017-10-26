package repack

import (
	"sort"

	"../screen"
	"../windows"
)

func Repack() {
	windows := windows.GetWindows()
	newIndex := 1

	var keys []int
	for k := range windows {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, oldIndex := range keys {
		if newIndex != oldIndex {
			screen.Renumber(oldIndex, newIndex)
		}
		newIndex++

	}
}
