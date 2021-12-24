package main

import (
	"fmt"
)

func main() {
	var A, B, C, D int

	fmt.Scanf("%d %d %d %d", &A, &B, &C, &D)

	if B > C && D > A && C > 0 && D > 0 && (C+D) > (A+B) && A%2 == 0 {
		fmt.Printf("Valores aceitos\n")
	} else {
		fmt.Printf("Valores nao aceitos\n")
	}
}
