
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

type ConversionRequest struct {
	HaveCurrency string  `json:"have"`
	WantCurrency string  `json:"want"`
	Amount       float64 `json:"amount"`
}

type ConversionResponse struct {
	NewAmount     float64 `json:"new_amount"`
	NewCurrency   string  `json:"new_currency"`
	OldCurrency   string  `json:"old_currency"`
	OldAmount     float64 `json:"old_amount"`
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

	request := ConversionRequest{
		HaveCurrency: haveCurrency,
		WantCurrency: wantCurrency,
		Amount:       amount,
	}

	apiURL := fmt.Sprintf("https://api.api-ninjas.com/v1/convertcurrency?have=%s&want=%s&amount=%f", request.HaveCurrency, request.WantCurrency, request.Amount)

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error fetching conversion data:", err)
		return
	}
	defer resp.Body.Close()

	var response ConversionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Printf("%f %s is equivalent to %f %s\n", amount, haveCurrency, response.NewAmount, wantCurrency)
}
