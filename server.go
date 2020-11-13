package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	p := person{name: "Bob", age: 14}
	fmt.Println(&p)
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

type person struct {
	name string
	age  int
}
