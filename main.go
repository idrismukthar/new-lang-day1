package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("======================================")
	fmt.Println("       👻 GHOST-CHAT P2P v1.0        ")
	fmt.Println("======================================")
	fmt.Print("Do you want to (H)ost or (J)oin? ")
	
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(strings.ToUpper(choice))

	if choice == "H" {
		startServer()
	} else {
		fmt.Print("Enter Host IP (e.g., 127.0.0.1): ")
		ip, _ := reader.ReadString('\n')
		startClient(strings.TrimSpace(ip))
	}
}

func startServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("[ERROR] Port busy!")
		return
	}
	fmt.Println("[GHOST] Waiting on port 8080...")
	conn, _ := ln.Accept()
	fmt.Println("[GHOST] Connected!")
	handleChat(conn)
}

func startClient(ip string) {
	// Using time for the timeout
	conn, err := net.DialTimeout("tcp", ip+":8080", 5*time.Second)
	if err != nil {
		fmt.Println("[ERROR] Host not found!")
		return
	}
	handleChat(conn)
}

func handleChat(conn net.Conn) {
	fmt.Println("--- SECURE CHANNEL OPEN ---")
	
	// Background listener for incoming messages
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("\n[STRANGER]: %s\n> ", scanner.Text())
		}
	}()

	// Foreground sender
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Fprintf(conn, msg+"\n")
		fmt.Print("> ")
	}
}