package main

import (
	"errors"
	"fmt"
)

const (
	USDToEUR = 0.92
	USDToRUB = 76.64

	CurrencyUSD = "USD"
	CurrencyEUR = "EUR"
	CurrencyRUB = "RUB"
)

func main() {
	initialCurrency := getInitialCurrency()
	currencyAmount := getCurrencyAmount()
	currentCurrency := getCurrentCurrency(initialCurrency)
	result := convertCurrency(currencyAmount, initialCurrency, currentCurrency)
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
	isValid := currency == CurrencyEUR || currency == CurrencyUSD || currency == CurrencyRUB

	if !isValid {
		return false, errors.New("неподдерживаемый тип валюты")
	}

	if isValid && initialCurrency == currency {
		return false, errors.New("валюта должна отличаться от исходной")
	}

	return true, nil
}

func convertCurrency(count float64, currencyFrom string, currencyTo string) float64 {
	var inUSD float64

	switch currencyFrom {
	case CurrencyUSD:
		inUSD = count
	case CurrencyEUR:
		inUSD = count / USDToEUR
	case CurrencyRUB:
		inUSD = count / USDToRUB
	}

	switch currencyTo {
	case CurrencyUSD:
		return inUSD
	case CurrencyEUR:
		return inUSD * USDToEUR
	case CurrencyRUB:
		return inUSD * USDToRUB
	}

	return 0
}

func outputResult(value float64) {
	fmt.Printf("Результат: %.0f\n", value)
}
