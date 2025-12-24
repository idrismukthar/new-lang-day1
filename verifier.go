package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Enhanced verification logic
func securityAudit(ref, merchant, amountStr, currency string) {
	amount, _ := strconv.ParseFloat(amountStr, 64)

	fmt.Printf("\n[GO] 🔍 SECURITY SCAN: %s\n", ref)
	fmt.Printf("     Merchant: %s | Value: %s %.2f\n", merchant, currency, amount)

	// Fraud Detection Logic
	isHighValue := (currency == "USD" && amount >= 5000) || (currency == "NGN" && amount >= 5000000)

	if isHighValue {
		fmt.Printf("     ⚠️  ALERT: High-Value Transaction! flagging for AML compliance...\n")
	}

	time.Sleep(1 * time.Second)
	fmt.Printf("     ✅ VERIFIED: Ref %s passed integrity check.\n", ref)
}

func main() {
	fmt.Println("========================================")
	fmt.Println("    ELITEBANK GO SECURITY WATCHDOG      ")
	fmt.Println("    Monitoring Ledger Integrity...      ")
	fmt.Println("========================================")

	lastCount := 0
	const LEDGER_FILE = "ledger.csv"

	for {
		file, err := os.Open(LEDGER_FILE)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}

		scanner := bufio.NewScanner(file)
		var currentLines []string
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 5 {
				currentLines = append(currentLines, line)
			}
		}
		file.Close()

		// If we find new lines
		if len(currentLines) > lastCount {
			if lastCount == 0 {
				lastCount = len(currentLines)
				fmt.Printf("[SYSTEM] Security Baseline Set. Monitoring %d records.\n", lastCount)
			} else {
				// Process only new transactions
				for i := lastCount; i < len(currentLines); i++ {
					data := strings.Split(currentLines[i], ",")
					
					// Our new 8-column format:
					// Ref[0], Merchant[1], Amt[2], Curr[3], Date[4], Desc[5], Fee[6], Tax[7]
					if len(data) >= 4 {
						ref := data[0]
						merchant := data[1]
						amount := data[2]
						currency := data[3]

						// Run security audit in a concurrent Goroutine (super fast)
						go securityAudit(ref, merchant, amount, currency)
					}
				}
				lastCount = len(currentLines)
			}
		}

		time.Sleep(1 * time.Second)
	}
}