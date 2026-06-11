package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Network Monitor Scanner Statring...")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	subnet := os.Getenv("SUBNET")

	jobs := make(chan string, 100)
	results := make(chan Device, 100)

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := 1; i <= 254; i++ {
		jobs <- generateIP(subnet, i) // sending work
	}

	close(jobs)

	for device := range results {

		fmt.Printf(
			"IP: %-15s Hostname: %s\n",
			device.IP,
			device.Hostname,
		)
	}
}

func generateIP(subnet string, host int) string {
	return fmt.Sprintf("%s%d", subnet, host)
}
