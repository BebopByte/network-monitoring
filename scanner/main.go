package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Network Monitor Scanner Statring...")

	cfg := LoadConfig()

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
		jobs <- generateIP(cfg.Subnet, i) // sending work
	}

	close(jobs)

	var devices []DeviceRequest

	for device := range results {

		fmt.Printf(
			"IP: %-15s Hostname: %-30s Response: %v\n",
			device.IP,
			device.Hostname,
			device.ResponseTime,
		)

		devices = append(devices, DeviceRequest{
			IP:           device.IP,
			Hostname:     device.Hostname,
			Online:       device.Online,
			ResponseTime: device.ResponseTime.Milliseconds(),
			ScannedAt:    time.Now().UTC().Format(time.RFC3339),
		})
	}

	err := sendDevices(devices, cfg.ApiUrl)

	if err != nil {
		fmt.Println("Failed to send devices: ", err)
	}
}
