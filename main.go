package main

import (
	"bufio"
	"encoding/json"
	"encoding/csv"
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

type Currency struct {
	Symbol string
	Name   string
}

func CurrencyExists(csvFilename, input string) (bool, error) {
	file, err := os.Open("currencies.csv")
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break // End of file reached
			}
			return false, err
		}

		if len(record) > 0 && record[0] == input {
			return true, nil
		}
	}

	return false, nil
}

func main() {

	csvFilename := "currencies.csv"
	var input1, input2 string

	fmt.Print("Enter the 3 letter currency symbol you wish to exchange: ")
	_, err := fmt.Scan(&input1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	found1, err := CurrencyExists(csvFilename, input1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Print("Enter the 3 letter currency symbol you wish to recieve:  ")
	_, err = fmt.Scan(&input2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	found2, err := CurrencyExists(csvFilename, input2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if found1 && found2 {
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
			HaveCurrency: input1,
			WantCurrency: input2,
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

		fmt.Printf("%f %s is equivalent to %f %s\n", amount, input1, response.NewAmount, input2)
	} else {
		fmt.Print("Currencies not found")
	}
}