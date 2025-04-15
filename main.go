package main

import (
	"fmt"
)

func main() {

	const (
		selectCount    = "Enter Count"
		selectCurrency = "Enter Currency Code"
	)

	var err error

	// Курс валют относительно 1 USD
	var rates = map[string]float64{
		"USD": 1.0,   // Доллар США
		"EUR": 0.92,  // Евро
		"RUB": 90.0,  // Российский рубль
		"JPY": 157.0, // Японская иена
		"CNY": 7.25,  // Китайский юань
		"GBP": 0.78,  // Британский фунт
		"KZT": 460.0, // Казахстанский тенге
		"TRY": 32.5,  // Турецкая лира
		"INR": 83.0,  // Индийская рупия
		"BRL": 5.12,  // Бразильский реал
		"AUD": 1.50,  // Австралийский доллар
		"CAD": 1.36,  // Канадский доллар
		"CHF": 0.89,  // Швейцарский франк
		"SEK": 10.8,  // Шведская крона
		"NOK": 10.5,  // Норвежская крона
	}

	fmt.Println("Welcome to the currency converter!")
	fmt.Print("Available currency for converting:\n")
	i := 0
	var currencies []string
	for k, v := range rates {
		i++
		currencies = append(currencies, k)
		fmt.Printf("%d. %s %v\n", i, k, v)
	}

	stage := selectCount

	var USDCount int

	for {

		switch stage {
		case selectCount:

			fmt.Print("Enter your count in USD: ")

			_, err = fmt.Scan(&USDCount)
			if err != nil {
				break
			}
			if USDCount < 1 {
				fmt.Println("Count must be greater than zero")
				break
			}
			stage = selectCurrency
			break

		case selectCurrency:
			fmt.Print("Select your currency you want to convert in upper list:\n")
			SelectedCurrency := 1
			_, err = fmt.Scan(&SelectedCurrency)
			if err != nil {
				stage = selectCount
				break
			}

			if SelectedCurrency < 1 {
				fmt.Println("Selected currency must be greater than zero")
				stage = selectCount
				break
			}
			if SelectedCurrency > len(currencies) {
				fmt.Println("Wrong choice of currency")
				stage = selectCount
				break
			}

			convertedCount := float64(USDCount) * rates[currencies[SelectedCurrency-1]]

			fmt.Printf("Converted result: %.2f USD = %.2f %s \n", float64(USDCount), convertedCount, currencies[SelectedCurrency-1])
			stage = selectCount
			break
		}

	}

}
