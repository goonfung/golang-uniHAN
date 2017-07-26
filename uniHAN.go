package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/// Unicode block
var unicock = "^[" +
	"\u0080-\u00ff" + //Latin-1 Supplement
	"\u2000-\u206f" + // General Punctuation
	"\u2600-\u27BF" + // other
	"\u2E80-\u2EFF" + // CJK Radicals Supplement
	"\u3000-\u303f" + // CJK Symbols and Punctuation
	"\uFE10-\ufe1f" + // Vertical Forms
	"\ufe30-\ufe4f" + // CJK Compatibility Forms
	"\uff00-\uffef" + // Halfwidth and Fullwidth Forms
	"]$"

func main() {
	var bigMap = make(map[rune]rune)
	var kidMap = make(map[rune]rune)

	han, err := ioutil.ReadFile("y:\\a.log")
	check(err)

	/// buiding big map
	for _, runie := range string(han) {
		if _, ok := bigMap[runie]; !ok {
			bigMap[runie] = runie
		}
	}
	var kid = regexp.MustCompile(unicock)
	for _, v := range bigMap {
		if kid.MatchString(string(v)) {
			kidMap[v] = v
			fmt.Printf("%s %U\n", string(v), v)
		}
	}

	println(len(kidMap))
}
