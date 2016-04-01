package main

import "fmt"

func main() {
	string_map := make(map[string]string)
	string_map["foo"] = "bar"
	string_map["cat"] = "dog"
	fmt.Println(string_map)

	age_map := make(map[string]int)
	age_map["Kevin"] = 30
	age_map["Melissa"] = 26
	fmt.Println(age_map)

	pet := "Charlie"
	age_map[pet] = 2
	fmt.Println(age_map)

	delete(age_map, pet)
	fmt.Println(age_map)

	predefined_map := map[string]string{
		"cat": "meow",
		"dog": "woof",
	}

	fmt.Println(predefined_map)
}
