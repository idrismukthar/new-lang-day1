package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	allClients = make(map[net.Conn]string)
	mutex      sync.Mutex
	maxClients int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("======================================")
	fmt.Println("    👻 GHOST-CHAT v5.1: THE FIX      ")
	fmt.Println("======================================")

	fmt.Print("Enter Nickname: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Port (e.g., 9999): ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	fmt.Print("Host (H) or Join (J)? ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(strings.ToUpper(choice))

	if choice == "H" {
		fmt.Print("Set Max People (Limit): ")
		fmt.Scan(&maxClients)
		startServer(port, name)
	} else {
		fmt.Print("Enter Host IP: ")
		ip, _ := reader.ReadString('\n')
		startClient(strings.TrimSpace(ip), port, name)
	}
}

func startServer(port, hostName string) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("❌ BIND ERROR: Port busy! Use a different number.")
		return
	}
	fmt.Printf("[!] Room Live on %s. Limit: %d. Type '/kill' to exit.\n", port, maxClients)

	// Host's own input handler (Server speaks)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			if strings.ToLower(text) == "/kill" {
				broadcast("SYSTEM", "SERVER SHUTDOWN BY HOST", nil)
				os.Exit(0)
			}
			broadcast(hostName, text, nil)
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			break
		}

		mutex.Lock()
		if len(allClients) >= maxClients {
			fmt.Fprintln(conn, "SYSTEM: Room Full")
			conn.Close()
			mutex.Unlock()
			continue
		}
		allClients[conn] = "New Ghost"
		mutex.Unlock()

		// Pass the hostName and handle connection
		go handleConnection(conn, hostName, true) 
	}
}

func startClient(ip, port, name string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("❌ CONNECTION ERROR:", err)
		return
	}
	fmt.Fprintf(conn, "/name %s\n", name)
	handleConnection(conn, name, false)
}

func handleConnection(conn net.Conn, myName string, isServer bool) {
	// 1. RECEIVER
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			msg := scanner.Text()
			
			// If we are the Server, we must broadcast what we receive!
			if isServer {
				if strings.HasPrefix(msg, "/name ") {
					mutex.Lock()
					allClients[conn] = strings.TrimPrefix(msg, "/name ")
					mutex.Unlock()
					continue
				}
				// Server relays client message to everyone else
				broadcast(allClients[conn], msg, conn)
			} else {
				// Client just prints what they get
				fmt.Printf("\r%s\n> ", msg)
			}
		}
	}()

	// 2. SENDER
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		
		// File Sending logic
		if strings.HasPrefix(text, "/send ") {
			fileName := strings.TrimPrefix(text, "/send ")
			content, err := os.ReadFile(fileName)
			if err != nil {
				fmt.Println("❌ File not found!")
				continue
			}
			fmt.Fprintf(conn, "[FILE: %s] %s\n", fileName, string(content))
			continue
		}

		fmt.Fprintf(conn, "%s\n", text)
		if !isServer { fmt.Print("> ") }
	}
}

func broadcast(sender, msg string, exclude net.Conn) {
	mutex.Lock()
	defer mutex.Unlock()
	formatted := fmt.Sprintf("[%s]: %s", sender, msg)
	
	// Print on Server console
	fmt.Printf("\r%s\n> ", formatted)

	for client := range allClients {
		if client != exclude {
			fmt.Fprintln(client, formatted)
		}
	}
}