package main

import (
	"fmt"
)

func main() {
	var radio, A float64

	fmt.Scanf("%f", &radio)

	A = 3.14159 * (radio * radio)

	fmt.Printf("A=%.4f\n", A)

}
