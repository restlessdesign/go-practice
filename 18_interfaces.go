package main

import (
	"fmt"
	"math"
	"reflect"
)

// Interfaces __________________________________________________________________

type shape interface {
	area() float64
	perimeter() float64
}

// Types _______________________________________________________________________

type circle struct {
	radius float64
}

func (r circle) area() float64 {
	return math.Pi * r.radius * r.radius
}

func (r circle) perimeter() float64 {
	return 2 * math.Pi * r.radius
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// Helper Functions ____________________________________________________________

func measure(s shape) {
	fmt.Println(reflect.TypeOf(s), s)
	fmt.Printf("Area: %.1f\n", s.area())
	fmt.Printf("Perimeter: %.1f\n\n", s.perimeter())
}

// Main ________________________________________________________________________

func main() {
	my_circle := circle{radius: 30}
	my_rectangle := rectangle{width: 10, height: 5}
	my_square := rectangle{width: 9, height: 9}

	my_shapes := []shape{my_circle, my_rectangle, my_square}
	for _, shape := range my_shapes {
		measure(shape)
	}
}
