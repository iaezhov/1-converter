package main

import "fmt"

func main() {
	const USDToEUR = 0.92
	const USDToRUB = 76.64
	const EURToRUB = USDToRUB / USDToEUR
	var EURCount float64 = 100
	fmt.Println(EURCount * EURToRUB)
}
