package main

import (
	"fmt"
)

func main() {
	var A, B, C, MEDIA float64

	fmt.Scanf("%f", &A)
	fmt.Scanf("%f", &B)
	fmt.Scanf("%f", &C)

	MEDIA = ((2 * A) + (3 * B) + (5 * C)) / 10

	fmt.Printf("MEDIA = %.1f\n", MEDIA)
}
