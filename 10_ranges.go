package main

import "fmt"

func main() {
	iterable_a := []byte{1, 2, 3, 4, 5}
	iterable_b := []string{"foo", "bar", "baz"}
	animal_sounds := map[string]string{"dog": "woof!", "cat": "meow", "fox": "agijogew!!oijdklsd!!!"}
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, val := range iterable_a {
		fmt.Println(val)
	}

	for i, val := range iterable_b {
		fmt.Println(i, val)
	}

	for key, val := range animal_sounds {
		fmt.Printf("The %s goes “%s!”\n", key, val)
	}

	for unicode_index, code_point := range alphabet {
		fmt.Println(unicode_index, code_point)

		if unicode_index == len(alphabet)-1 {
			fmt.Printf("Now I know my %s %s %s’s! Next time, won’t you fmt.Println() with me?", "0x97", "0x98", "0x99")
		}
	}
}
