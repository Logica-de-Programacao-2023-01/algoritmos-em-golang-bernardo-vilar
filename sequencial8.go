package main

import "fmt"

func main() {
	var diaria float32
	var dias float32
	fmt.Print("Quantos dias você trabalha?")
	fmt.Scan(&dias)
	fmt.Print("Qual o valor da sua diária?")
	fmt.Scan(&diaria)
	fmt.Println("Seu salário é:", diaria*dias)
}
