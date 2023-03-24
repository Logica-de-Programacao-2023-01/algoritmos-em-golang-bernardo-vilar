package main

import "fmt"

func main() {
	var preco float32
	fmt.Print("Qual o preço do produto?")
	fmt.Scan(&preco)
	fmt.Print(preco * 0.9)

}
