package main

import "fmt"

func main() {
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
	fmt.Println("Available currency for converting:")

	i := 0
	var currencies []string
	for k, v := range rates {
		i++
		currencies = append(currencies, k)
		fmt.Printf("%d. %s %v\n", i, k, v)
	}

	var USDCount float64
	for {
		fmt.Print("Enter your count in USD: ")
		_, err = fmt.Scan(&USDCount)
		if err != nil {
		}
		if USDCount <= 0 {
			fmt.Println("Count must be greater than zero")
			continue
		}
		fmt.Print("Select your currency you want to convert in upper list:\n")
		selectedCurrency := 1
		_, err = fmt.Scan(&selectedCurrency)
		if err != nil {
		}
		if selectedCurrency < 1 || selectedCurrency > len(currencies) {
			fmt.Println("Wrong choice of currency")
			continue
		}
		convertedCount := USDCount * rates[currencies[selectedCurrency-1]]
		fmt.Printf("Converted result: %.2f USD = %.2f %s \n", USDCount, convertedCount, currencies[selectedCurrency-1])
	}
}
