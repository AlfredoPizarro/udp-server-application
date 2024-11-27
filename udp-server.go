package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func main() {
	port := ":1814"
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v\n", err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP port %s: %v\n", port, err)
	}
	defer conn.Close()

	log.Printf("UDP server listening on port %s\n", port)

	var connectionCount int
	var mu sync.Mutex

	buffer := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP connection: %v\n", err)
			continue
		}

		// Increment connection counter
		mu.Lock()
		connectionCount++
		log.Printf("Connection number %d\n", connectionCount)
		mu.Unlock()

		// Respond to the client
		response := fmt.Sprintf("Hello client %s", remoteAddr.IP.String())
		_, err = conn.WriteToUDP([]byte(response), remoteAddr)
		if err != nil {
			log.Printf("Error sending response to client: %v\n", err)
			continue
		}

		log.Printf("Received message: %s from %s\n", string(buffer[:n]), remoteAddr.String())
	}
}
