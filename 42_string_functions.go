package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

// Taking the example one step further to build a function that can replace
// multiple characters.
func ReplaceMulti(s string, sources []string, replacements []string) string {
	for i, _ := range sources {
		s = strings.Replace(s, sources[i], replacements[i], -1)
	}

	return s
}

func L33tSp34k(s string) string {
	return ReplaceMulti(s, []string{"to", "o", "l", "r", "e", "a", "t", "y"}, []string{"2", "0", "1", "2", "3", "4", "7", "j"})
}

func main() {
	p(strings.Contains("test", "es"))
	p(strings.HasPrefix("inception", "in"))
	p(strings.HasPrefix("inception", "con"))
	p(strings.HasSuffix("inception", "ion"))
	p(strings.Index("abcdefg", "b"))
	p(strings.Join([]string{"a", "b", "c"}, " | "))

	fmt.Printf("g%sal!\n", strings.Repeat("o", 10))
	p(strings.Replace("leet", "e", "3", -1))
	p(ReplaceMulti("leet", []string{"l", "e", "t"}, []string{"1", "3", "7"}))
	p(L33tSp34k("all your base are belong to us"))
}
