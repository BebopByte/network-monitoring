package main

import (
	"fmt"
	"net"
	"time"
)

func probe(ip string) (bool, time.Duration) {

	timeout := 300 * time.Millisecond
	ports := []string{"22", "80", "443"}

	for _, port := range ports {

		start := time.Now()

		conn, err := net.DialTimeout("tcp", ip+":"+port, timeout)

		if err == nil {

			elapsed := time.Since(start)

			// for testing connectivity, close immediately after it connects
			conn.Close()

			return true, elapsed
		}
	}

	return false, 0
}

func getHostname(ip string) string {

	names, err := net.LookupAddr(ip)

	if err != nil || len(names) == 0 {
		return "Unknown"
	}

	return names[0]
}

func generateIP(subnet string, host int) string {
	return fmt.Sprintf("%s%d", subnet, host)
}
