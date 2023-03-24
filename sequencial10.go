package main

import "fmt"

func main() {
	var peso float32
	fmt.Print("Qual o seu peso em kg?")
	fmt.Scan(&peso)
	fmt.Println("Seu peso em libras é:", peso*2.205)
}
