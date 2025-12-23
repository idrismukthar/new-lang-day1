package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// This function simulates the Bank API check
func verifyWithBank(ref string) {
	fmt.Printf("\n [GO] New Transaction Detected: %s\n", ref)
	fmt.Printf(" [GO] Verifying with Central Bank API...\n")
	time.Sleep(2 * time.Second) 
	fmt.Printf(" [GO] VERIFIED: %s is legit. Funds confirmed.\n", ref)
	fmt.Print(" [GO] Waiting for new data... ")
}

func main() {
	fmt.Println("=== GO REAL-TIME MONITOR STARTING ===")
	fmt.Println(" Monitoring ledger.csv for new payments...")
	
	lastCount := 0

	for {
		file, err := os.Open("ledger.csv")
		if err != nil {
			// If file doesn't exist yet, just wait
			time.Sleep(1 * time.Second)
			continue
		}

		scanner := bufio.NewScanner(file)
		var currentLines []string
		for scanner.Scan() {
			currentLines = append(currentLines, scanner.Text())
		}
		file.Close()

		// Logic: If the file has more lines than before, we have new payments!
		if len(currentLines) > lastCount {
			// If it's the first time running, just mark what's already there
			if lastCount == 0 {
				lastCount = len(currentLines)
				fmt.Printf(" [GO] Loaded %d existing records. System Ready.\n", lastCount)
			} else {
				// Process only the brand new lines
				for i := lastCount; i < len(currentLines); i++ {
					data := strings.Split(currentLines[i], ",")
					if len(data) > 0 {
						go verifyWithBank(data[0]) 
					}
				}
				lastCount = len(currentLines)
			}
		}

		time.Sleep(1 * time.Second) // Poll the file every second
	}
}