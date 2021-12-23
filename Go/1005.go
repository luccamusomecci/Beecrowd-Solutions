package main

import (
	"fmt"
)

func main() {
	var A, B, MEDIA float64

	fmt.Scanf("%f", &A)
	fmt.Scanf("%f", &B)

	MEDIA = ((3.5 * A) + (7.5 * B)) / 11

	fmt.Printf("MEDIA = %.5f", MEDIA)
}
