package main

import (
	"fmt"
)

func main() {

	var nombre string
	var salarioFijo, ventas, total float64

	fmt.Scanf("%s", &nombre)
	fmt.Scanf("%f", &salarioFijo)
	fmt.Scanf("%f", &ventas)

	total = ((ventas * 15) / 100) + salarioFijo

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
