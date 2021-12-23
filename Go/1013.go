package main

import (
	"fmt"
)

func main() {
	var A, B, C int

	fmt.Scanf("%d", &A)
	fmt.Scanf("%d", &B)
	fmt.Scanf("%d", &C)

	if A > B && A > C {
		fmt.Printf("%d eh o maior\n", A)
	} else if B > C {
		fmt.Printf("%d eh o maior\n", B)
	} else {
		fmt.Printf("%d eh o maior\n", C)
	}

}
