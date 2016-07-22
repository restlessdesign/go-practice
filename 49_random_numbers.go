package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// If you re-run this program, these “random” numbers will be the same!
	rand_int_1 := rand.Intn(50)
	rand_int_2 := rand.Intn(50)
	fmt.Println(rand_int_1, rand_int_2)

	rand_float_1 := rand.Float64() * 100
	rand_float_2 := rand.Float64() * 100
	fmt.Println(rand_float_1, rand_float_2)

	// To avoid this, we can create a seed value and construct a new
	// random number generator
	seed := rand.NewSource(time.Now().UnixNano())
	srand := rand.New(seed)

	// These values will change each time the program is run
	srand_int_1 := srand.Intn(50)
	srand_int_2 := srand.Intn(50)
	fmt.Println(srand_int_1, srand_int_2)

	srand_float_1 := srand.Float64() * 100
	srand_float_2 := srand.Float64() * 100
	fmt.Println(srand_float_1, srand_float_2)
}
