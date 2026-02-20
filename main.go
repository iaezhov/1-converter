package main

import "fmt"

func main() {
	EURCount := getUserInput()
	result := convertCurrency(EURCount, "EUR", "RUB")
	outputResult(result)
}

func getUserInput() float64 {
	var count float64
	fmt.Scan(&count)
	return count
}

func convertCurrency(count float64, currencyFrom string, currencyTo string) float64 {
	fmt.Println("Изначальная валюта:", currencyFrom)
	fmt.Println("Конвертация в валюту:", currencyTo)
	const USDToEUR = 0.92
	const USDToRUB = 76.64
	const EURToRUB = USDToRUB / USDToEUR
	return count * EURToRUB
}

func outputResult(value float64) {
	fmt.Printf("Результат: %.0f", value)
}
