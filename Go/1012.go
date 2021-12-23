package main

import (
	"fmt"
)

func main() {
	var A, B, C, TRIAN, CICR, TRAP, CUADR, RECT float64

	fmt.Scanf("%f %f %f", &A, &B, &C)

	TRIAN = A * C / 2
	CICR = 3.14159 * C * C
	TRAP = ((A + B) * C) / 2.0
	CUADR = B * B
	RECT = A * B

	fmt.Printf("TRIANGULO: %.3f\n", TRIAN)
	fmt.Printf("CIRCULO: %.3f\n", CICR)
	fmt.Printf("TRAPEZIO: %.3f\n", TRAP)
	fmt.Printf("QUADRADO: %.3f\n", CUADR)
	fmt.Printf("RETANGULO: %.3f\n", RECT)
}
