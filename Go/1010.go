package main

import (
	"fmt"
)

func main() {
	var cod1, cant1, cod2, cant2 int
	var prec1, prec2, total float64

	fmt.Scanf("%d %d %f", &cod1, &cant1, &prec1)
	fmt.Scanf("%d %d %f", &cod2, &cant2, &prec2)

	total = ((prec1 * float64(cant1)) + (prec2 * float64(cant2)))

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
