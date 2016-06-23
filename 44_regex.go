package main

import (
	"fmt"
	"regexp"
)

func main() {
	for _, val := range []string{"peach", "patch", "poach", "pitch", "pug"} {
		is_match, _ := regexp.MatchString("p([a-z]+ch)", val)
		fmt.Println(val, is_match)
	}

	fmt.Println("----------------")

	r, _ := regexp.Compile("p([a-z])+ch")
	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.Match([]byte("peach")))

	// Match all
	fmt.Println(r.FindAllString("peach patch pitch pug", -1))

	// Match a maximum of 2
	fmt.Println(r.FindAllString("peach patch pitch pug", 2))

	k, _ := regexp.Compile("{([a-z])+}")
	fmt.Println(k.ReplaceAllString("lol but why use {technology} instead of string substitution tho?", "regex"))
}
