package main

import (
	"fmt"
)

func main() {
	var NUMBER, HOURS int
	var SALARY float64

	fmt.Scanf("%d", &NUMBER)
	fmt.Scanf("%d", &HOURS)
	fmt.Scanf("%f", &SALARY)

	fmt.Printf("NUMBER = %d\n", NUMBER)
	fmt.Printf("SALARY = U$ %.2f\n", SALARY*float64(HOURS))

}
