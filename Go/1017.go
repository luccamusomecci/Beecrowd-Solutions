package main

import (
	"fmt"
)

func main() {
	var tiempo, velocidad, kilometros, media float64

	fmt.Scanf("%f", &tiempo)
	fmt.Scanf("%f", &velocidad)
	kilometros = tiempo * velocidad
	media = kilometros / 12
	fmt.Printf("%.3f\n", media)
}
