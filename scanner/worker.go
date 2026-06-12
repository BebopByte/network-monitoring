package main

import "sync"

func worker(jobs <-chan string, results chan<- Device, wg *sync.WaitGroup) {

	defer wg.Done()

	for ip := range jobs { // receiving work

		online, responseTime := probe(ip)

		if online {

			device := Device{
				IP:           ip,
				Hostname:     getHostname(ip),
				Online:       true,
				ResponseTime: responseTime,
			}

			results <- device
		}
	}
}
