package util

import "fmt"

type HashSetInt map[int]struct{}

func (set HashSetInt) Add(element int) {
	set[element] = struct{}{}
}

func (set HashSetInt) Remove(element int) {
	delete(set, element)
}

func (set HashSetInt) Contains(element int) bool {
	_, found := set[element]
	return found
}

func (set HashSetInt) Print() {
	for element := range set {
		fmt.Println(element)
	}
}
