package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ConversionResult struct {
	ConvertedAmount float64 `json:"converted_amount"`
}

func main() {
	haveCurrency := "USD"
	wantCurrency := "EUR"

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid amount.")
		return
	}

	apiURL := fmt.Sprintf("https://api.api-ninjas.com/v1/convertcurrency?have=%s&want=%s&amount=%f", haveCurrency, wantCurrency, amount)

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching conversion data:", err)
		return
	}
	defer resp.Body.Close()

	var result ConversionResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("%f %s is equivalent to %f %s\n", amount, haveCurrency, result.ConvertedAmount, wantCurrency)
}
