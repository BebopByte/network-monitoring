package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Network Monitor Scanner Statring...")

	subnet := os.Getenv("SUBNET")

	for i := 1; i <= 254; i++ {

		ip := generateIP(subnet, i)

		fmt.Println(ip)

	}
}

func generateIP(subnet string, host int) string {
	return fmt.Sprintf("%s%d", subnet, host)
}
