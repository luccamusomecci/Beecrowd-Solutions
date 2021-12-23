package main

import (
	"fmt"
	"math"
)

const pi = 3.14159

func main() {
	var R, VOLUME float64

	fmt.Scanf("%f", &R)
	VOLUME = (4.0 / 3) * pi * math.Pow(R, 3.0)
	fmt.Printf("VOLUME = %.3f\n", VOLUME)
}
