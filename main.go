package main

import (
	"errors"
	"fmt"
)

const (
	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"
	CurrencyRUB = "RUB"
)

var ratesToUSD = map[string]float64{
	CurrencyUSD: 1.0,
	CurrencyEUR: 0.92,
	CurrencyRUB: 76.64,
}

func main() {
	initialCurrency := getInitialCurrency()
	currencyAmount := getCurrencyAmount()
	currentCurrency := getCurrentCurrency(initialCurrency)
	result := convertCurrency(currencyAmount, initialCurrency, currentCurrency, &ratesToUSD)
	outputResult(result)
}

func getCurrencyHint(initialCurrency string) string {
	switch initialCurrency {
	case CurrencyEUR:
		return fmt.Sprintf("(%s/%s)", CurrencyUSD, CurrencyRUB)
	case CurrencyUSD:
		return fmt.Sprintf("(%s/%s)", CurrencyEUR, CurrencyRUB)
	case CurrencyRUB:
		return fmt.Sprintf("(%s/%s)", CurrencyEUR, CurrencyUSD)
	default:
		return fmt.Sprintf("(%s/%s/%s)", CurrencyEUR, CurrencyUSD, CurrencyRUB)
	}
}

func getInitialCurrency() string {
	var currency string

	for {
		fmt.Printf("Изначальная валюта %v: ", getCurrencyHint(""))
		fmt.Scan(&currency)
		isValid, err := validateCurrency(currency, "")
		if isValid {
			break
		}
		fmt.Printf("Ошибка: %v. Попробуйте снова.\n", err)
	}

	return currency
}

func getCurrentCurrency(initialCurrency string) string {
	var currency string

	for {
		fmt.Printf("Желаемая валюта %v: ", getCurrencyHint(initialCurrency))
		fmt.Scan(&currency)
		isValid, err := validateCurrency(currency, initialCurrency)
		if isValid {
			break
		}
		fmt.Printf("Ошибка: %v. Попробуйте снова.\n", err)
	}

	return currency
}

func getCurrencyAmount() float64 {
	var currency float64
	for {
		fmt.Print("Введите сумму для конвертации: ")
		fmt.Scan(&currency)
		if currency > 0 {
			break
		}
		fmt.Println("Ошибка: сумма должна быть больше 0. Попробуйте снова.")
	}
	return currency
}

func validateCurrency(currency, initialCurrency string) (bool, error) {
	_, exists := ratesToUSD[currency]

	if !exists {
		return false, errors.New("неподдерживаемый тип валюты")
	}

	if initialCurrency == currency {
		return false, errors.New("валюта должна отличаться от исходной")
	}

	return true, nil
}

func convertCurrency(count float64, currencyFrom string, currencyTo string, ratesPointer *map[string]float64) float64 {
	// Конвертируем в USD
	inUSD := count / (*ratesPointer)[currencyFrom]

	// Конвертируем из USD в целевую валюту
	return inUSD * (*ratesPointer)[currencyTo]
}

func outputResult(value float64) {
	fmt.Printf("Результат: %.0f\n", value)
}
