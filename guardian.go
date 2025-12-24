package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("    ELITEBANK GO VAULT GUARDIAN V1    ")
	fmt.Println("    (Monitoring Treasury Assets)      ")
	fmt.Println("========================================")

	for {
		calculateWealth()
		// Sleep for 10 seconds before checking again
		time.Sleep(10 * time.Second)
	}
}

func calculateWealth() {
	file, err := os.Open("treasury.csv")
	if err != nil {
		fmt.Println("[WAITING] Treasury file not found yet...")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var totalNGN, totalUSD float64

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		// treasury.csv format: Date, Ref, Revenue, Tax, Currency
		if len(record) >= 5 {
			revenue, _ := strconv.ParseFloat(record[2], 64)
			currency := record[4]

			if currency == "NGN" {
				totalNGN += revenue
			} else if currency == "USD" {
				totalUSD += revenue
			}
		}
	}

	fmt.Printf("\n[LIVE ASSETS] %s\n", time.Now().Format("15:04:05"))
	fmt.Printf("🏦 Bank Wealth: ₦%.2f | $%.2f\n", totalNGN, totalUSD)
}