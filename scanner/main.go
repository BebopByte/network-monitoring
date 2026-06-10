package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Network Monitor Scanner Statring...")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	subnet := os.Getenv("SUBNET")

	for i := 1; i <= 20; i++ {

		ip := generateIP(subnet, i)

		go scan(ip)
	}

	time.Sleep(3 * time.Second) // giving goroutines some time to finish
}

func generateIP(subnet string, host int) string {
	return fmt.Sprintf("%s%d", subnet, host)
}

func scan(ip string) {

	if isAlive(ip) {
		fmt.Printf("%s is ONLINE\n", ip)
	} else {
		fmt.Printf("%s is UNREACHABLE\n", ip)
	}
}

func isAlive(ip string) bool {

	ports := []string{"22", "80", "443"}

	timeout := 300 * time.Millisecond

	for _, port := range ports {

		conn, err := net.DialTimeout("tcp", ip+":"+port, timeout)

		if err == nil {
			// for testing connectivity, close immediately after it connects
			conn.Close()
			return true
		}

	}

	return false

}
