package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	me := person{"Kevin", 30}
	fmt.Println(me)

	myself := person{
		name: "Kevin",
		age:  30,
	}

	fmt.Printf("My name is %s. I am %d years old.", myself.name, myself.age)
}
