package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual é o número?")
	fmt.Scan(&x)
	fmt.Print(x-1, x, x+1)

}
