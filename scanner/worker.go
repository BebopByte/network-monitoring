package main

import "sync"

func worker(jobs <-chan string, results chan<- Device, wg *sync.WaitGroup) {

	defer wg.Done()

	for ip := range jobs { // receiving work

		device := DiscoverDevice(ip)

		if device != nil {
			results <- *device
		}
	}
}
