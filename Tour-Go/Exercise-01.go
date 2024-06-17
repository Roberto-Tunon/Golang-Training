package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Println(add(42, 13))
	/*
			  Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
		      Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	*/
	var i, j int = 1, 2
	k := 3

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
