package main

import "sync"

func worker(jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {

	defer wg.Done()

	for ip := range jobs { // receiving work

		if isAlive(ip) {
			results <- ip + " is ONLINE"
		}
	}
}
