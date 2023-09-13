package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	// zero value of maps is nil; a map with no k,v ; k cannot be added either??
	// make is always better to init nil map instead of empty literal
	//wc := map[string]int {} // empty map literal
	wc := make(map[string]int)

	sub_strs := strings.Fields(s)

	for _, str := range sub_strs {
		elem, ok := wc[str] // check for existing value for key str

		if !ok { // elem does not not exist
			wc[str] = 1
		} else { // elem exists
			wc[str] = elem + 1
		}
	}

	return wc
}
