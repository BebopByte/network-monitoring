package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"sync"
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

	jobs := make(chan string, 100)
	results := make(chan string, 100)

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

	// print results
	for r := range results {
		fmt.Println(r)
	}
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

	timeout := 300 * time.Millisecond
	ports := []string{"22", "80", "443"}

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

func worker(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	for ip := range jobs { // receiving work

		if isAlive(ip) {
			results <- ip + " is ONLINE"
		}
	}
}
