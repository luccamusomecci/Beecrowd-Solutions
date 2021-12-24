package main

import (
	"fmt"
)

func main() {
	var edad, dias, mes, ano int

	fmt.Scanf("%d", edad)

	ano = edad / 365

	mes = (edad - (365 * ano)) / 30

	dias = (edad - (365 * ano) - (30 * mes))

	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", ano, mes, dias)
}
