package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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

		if ping(ip) {
			fmt.Printf("%s is ONLINE\n", ip)
		}
	}
}

func generateIP(subnet string, host int) string {
	return fmt.Sprintf("%s%d", subnet, host)
}

func ping(ip string) bool {

	cmd := exec.Command("ping", "-n", "1", ip)

	err := cmd.Run()

	return err == nil

}
